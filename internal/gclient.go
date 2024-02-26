// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/g41797/pingopong"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func NewClient(trl hclog.Level) (pingopong.PingPong, func()) {
	c := Gclient{Level: trl}
	return &c, c.Clean
}

var _ pingopong.PingPong = (*Gclient)(nil)

type Gclient struct {
	lock    sync.Mutex
	Level   hclog.Level
	cleanup func()
	impl    pingopong.PingPong
}

func (s *Gclient) Play(ctx context.Context, b *pingopong.Ball) (*pingopong.Ball, error) {

	s.lock.Lock()
	defer s.lock.Unlock()

	if err := s.run(); err != nil {
		return nil, err
	}

	return s.impl.Play(ctx, b)
}

func (s *Gclient) Clean() {
	if s == nil {
		return
	}
	if s.cleanup == nil {
		return
	}
	s.cleanup()
	s.cleanup = nil
	s.impl = nil
}

func (s *Gclient) run() error {
	if s == nil {
		return fmt.Errorf("nil client")
	}

	if s.cleanup != nil || s.impl != nil {
		return nil
	}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshake,
		Plugins:         pluginMap,
		Cmd:             exec.Command(os.Args[0], os.Args[1:]...),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC},
		Logger: hclog.New(&hclog.LoggerOptions{
			Output: hclog.DefaultOutput,
			Level:  s.Level,
			Name:   RunningExeName() + "_client",
		}),
	})

	clean := client.Kill

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(pingPongPluginName)
	if err != nil {
		return err
	}

	s.impl = raw.(pingopong.PingPong)
	s.cleanup = clean

	return nil
}
