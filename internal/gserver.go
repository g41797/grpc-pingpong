// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal

import (
	"context"

	"github.com/g41797/grpc-pingpong/pingpong"
	"github.com/g41797/grpc-pingpong/shared"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

var _ shared.PingPong = (*gserver)(nil)

type gserver struct {
	level hclog.Level
}

func (s *gserver) Play(ctx context.Context, b *shared.Ball) (*shared.Ball, error) {
	// TODO: Add real implementation
	return b, nil
}

func (s *gserver) Run() {

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: Handshake,
		Plugins: map[string]plugin.Plugin{
			PingPongPluginName: &PingPongGRPCPlugin{Impl: s},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
		Logger: hclog.New(&hclog.LoggerOptions{
			Output: hclog.DefaultOutput,
			Level:  s.level,
			Name:   pingpong.RunningExeName() + "_plugin",
		}),
	})

}
