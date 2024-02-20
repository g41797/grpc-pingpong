// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package shared

import (
	"context"

	"github.com/g41797/grpc-pingpong/pb"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// GRPCClient is an implementation of PingPong that talks over GRPC.
type GRPCClient struct{ client pb.PingPongClient }

func (gc *GRPCClient) Play(ctx context.Context, in *Ball) (*Ball, error) {
	pb, err := ToProto(in)
	if err != nil {
		return nil, err
	}

	pb, err = gc.client.Play(ctx, pb)
	if err != nil {
		return nil, err
	}

	return FromProto(pb)
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl PingPong
}

func (gs *GRPCServer) Play(ctx context.Context, pb *pb.Ball) (*pb.Ball, error) {
	b, _ := FromProto(pb)
	b, err := gs.Impl.Play(ctx, b)
	if err != nil {
		return nil, err
	}
	return ToProto(b)
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type PingPongGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go.
	Impl PingPong
}

func (p *PingPongGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterPingPongServer(s, &GRPCServer{Impl: p.Impl})
	return nil
}

func (p *PingPongGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{client: pb.NewPingPongClient(c)}, nil
}

const PingPongPluginName = "pingpong_grpc"

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "UnaryGRPC",
	MagicCookieValue: "pingpong",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	PingPongPluginName: &PingPongGRPCPlugin{},
}
