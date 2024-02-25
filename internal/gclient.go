// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/g41797/pingopong/pingpong"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func NewClient(trl hclog.Level) (pingpong.PingPong, func()) {
	c := Gclient{Level: trl}
	return &c, c.Clean
}

var _ pingpong.PingPong = (*Gclient)(nil)

type Gclient struct {
	Level   hclog.Level
	cleanup func()
	impl    pingpong.PingPong
}

func (s *Gclient) Play(ctx context.Context, b *pingpong.Ball) (*pingpong.Ball, error) {

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

	s.impl = raw.(pingpong.PingPong)
	s.cleanup = clean

	return nil
}
