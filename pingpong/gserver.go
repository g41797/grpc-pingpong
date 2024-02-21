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

var _ shared.PingPong = (*gserver)(nil)

type gserver struct {
	level hclog.Level
}

func NewServer(trl hclog.Level) *gserver {
	return &gserver{level: trl}
}

func (s *gserver) Play(ctx context.Context, b *shared.Ball) (*shared.Ball, error) {
	// TODO: Add real implementation
	return b, nil
}

func (s *gserver) Run() {

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
