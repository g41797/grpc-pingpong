// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package plugin

import (
	"context"

	"github.com/g41797/grpc-pingpong/shared"
	"github.com/hashicorp/go-plugin"
)

var _ shared.PingPong = (*pingPongServer)(nil)

type pingPongServer struct{}

func (s *pingPongServer) Play(ctx context.Context, b *shared.Ball) (*shared.Ball, error) {
	// TODO: Add real implementation
	return b, nil
}

func (s *pingPongServer) Run() {

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			shared.PingPongPluginName: &shared.PingPongGRPCPlugin{Impl: s},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})

}
