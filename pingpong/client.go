// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingpong

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/g41797/grpc-pingpong/shared"
	"github.com/hashicorp/go-plugin"
	"github.com/mitchellh/go-ps"
)

func IsPluginProcess() bool {

	proc, _ := ps.FindProcess(os.Getpid())
	pproc, _ := ps.FindProcess(os.Getppid())

	return proc.Executable() == pproc.Executable()
}

func RunClient() {

	pcl := Client{}
	defer pcl.Clean()

	b := shared.Ball{Player: "noname"}

	res, err := pcl.Play(context.Background(), &b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if res.Player != b.Player {
		fmt.Printf("expected %s actual %s", b.Player, res.Player)
		return
	}

	return
}

var _ shared.PingPong = (*Client)(nil)

type Client struct {
	cleanup func()
	impl    shared.PingPong
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
