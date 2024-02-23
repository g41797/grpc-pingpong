// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingopong

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func NewGame(trl hclog.Level) (PingPong, func()) {
	result := &gclient{level: trl}
	return result, result.Clean
}

var _ PingPong = (*gclient)(nil)

type gclient struct {
	level   hclog.Level
	cleanup func()
	impl    PingPong
}

func (s *gclient) Play(ctx context.Context, b *Ball) (*Ball, error) {

	if err := s.run(); err != nil {
		return nil, err
	}

	return s.impl.Play(ctx, b)
}

func (s *gclient) Clean() {
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

func (s *gclient) run() error {
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
			Level:  s.level,
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

	s.impl = raw.(PingPong)
	s.cleanup = clean

	return nil
}
