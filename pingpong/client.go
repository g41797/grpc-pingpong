// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingpong

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/g41797/grpc-pingpong/internal"
	_ "github.com/g41797/grpc-pingpong/internal"
	"github.com/g41797/grpc-pingpong/shared"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/mitchellh/go-ps"
)

func ownExeName() string {
	proc, _ := ps.FindProcess(os.Getpid())
	return proc.Executable()
}

func parentExeName() string {
	proc, _ := ps.FindProcess(os.Getppid())
	return proc.Executable()
}

func IsPluginProcess() bool {
	return ownExeName() == parentExeName()
}

var _ shared.PingPong = (*client)(nil)

type client struct {
	level   hclog.Level
	cleanup func()
	impl    shared.PingPong
}

func NewGame(trl hclog.Level) (shared.PingPong, func()) {
	result := &client{level: trl}
	return result, result.Clean
}

func (s *client) Play(ctx context.Context, b *shared.Ball) (*shared.Ball, error) {

	if err := s.run(); err != nil {
		return nil, err
	}

	return s.impl.Play(ctx, b)
}

func (s *client) Clean() {
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

func (s *client) run() error {
	if s == nil {
		return fmt.Errorf("nil client")
	}

	if s.cleanup != nil || s.impl != nil {
		return nil
	}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: internal.Handshake,
		Plugins:         internal.PluginMap,
		Cmd:             exec.Command(os.Args[0], os.Args[1:]...),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolGRPC},
		Logger: hclog.New(&hclog.LoggerOptions{
			Output: hclog.DefaultOutput,
			Level:  s.level,
			Name:   ownExeName() + "_client",
		}),
	})

	clean := client.Kill

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(internal.PingPongPluginName)
	if err != nil {
		return err
	}

	s.impl = raw.(shared.PingPong)
	s.cleanup = clean

	return nil
}
