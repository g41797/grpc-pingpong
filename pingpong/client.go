// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingpong

import (
	"context"
	"fmt"
	"os"
	"os/exec"

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

var _ shared.PingPong = (*Client)(nil)

type Client struct {
	level   hclog.Level
	cleanup func()
	impl    shared.PingPong
}

func NewClient(trl hclog.Level) *Client {
	return &Client{level: trl}
}

func (s *Client) Play(ctx context.Context, b *shared.Ball) (*shared.Ball, error) {

	if err := s.run(); err != nil {
		return nil, err
	}

	return s.impl.Play(ctx, b)
}

func (s *Client) Clean() {
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

func (s *Client) run() error {
	if s == nil {
		return fmt.Errorf("nil client")
	}

	if s.cleanup != nil || s.impl != nil {
		return nil
	}

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		Cmd:             exec.Command(os.Args[0]),
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
	raw, err := rpcClient.Dispense(shared.PingPongPluginName)
	if err != nil {
		return err
	}

	s.impl = raw.(shared.PingPong)
	s.cleanup = clean

	return nil
}
