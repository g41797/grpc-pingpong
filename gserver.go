// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingopong

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func NewServer(trl hclog.Level) func() {
	srv := gserver{level: trl}
	return srv.Run
}

var _ PingPong = (*gserver)(nil)

type gserver struct {
	level hclog.Level
}

func (s *gserver) Play(ctx context.Context, b *Ball) (*Ball, error) {
	// TODO: Add real implementation
	return b, nil
}

func (s *gserver) Run() {

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshake,
		Plugins: map[string]plugin.Plugin{
			pingPongPluginName: &pingPongGRPCPlugin{Impl: s},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
		Logger: hclog.New(&hclog.LoggerOptions{
			Output: hclog.DefaultOutput,
			Level:  s.level,
			Name:   RunningExeName() + "_plugin",
		}),
	})

}
