// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingpong

import (
	"context"

	"github.com/g41797/grpc-pingpong/internal"
	"github.com/g41797/grpc-pingpong/shared"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

var _ shared.PingPong = (*server)(nil)

type server struct {
	level hclog.Level
}

func NewServer(trl hclog.Level) *server {
	return &server{level: trl}
}

func (s *server) Play(ctx context.Context, b *shared.Ball) (*shared.Ball, error) {
	// TODO: Add real implementation
	return b, nil
}

func (s *server) Run() {

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: internal.Handshake,
		Plugins: map[string]plugin.Plugin{
			internal.PingPongPluginName: &internal.PingPongGRPCPlugin{Impl: s},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
		Logger: hclog.New(&hclog.LoggerOptions{
			Output: hclog.DefaultOutput,
			Level:  s.level,
			Name:   ownExeName() + "_plugin",
		}),
	})

}
