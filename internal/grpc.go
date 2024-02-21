// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal

import (
	"context"
	"fmt"

	"github.com/g41797/grpc-pingpong/internal/pb"
	"github.com/g41797/grpc-pingpong/shared"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// GRPCClient is an implementation of PingPong that talks over GRPC.
type GRPCClient struct{ client pb.PingPongClient }

func (gc *GRPCClient) Play(ctx context.Context, in *shared.Ball) (*shared.Ball, error) {
	pb, err := toProto(in)
	if err != nil {
		return nil, err
	}

	pb, err = gc.client.Play(ctx, pb)
	if err != nil {
		return nil, err
	}

	return fromProto(pb)
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl shared.PingPong
}

func (gs *GRPCServer) Play(ctx context.Context, pb *pb.Ball) (*pb.Ball, error) {
	b, _ := fromProto(pb)
	b, err := gs.Impl.Play(ctx, b)
	if err != nil {
		return nil, err
	}
	return toProto(b)
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type PingPongGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go.
	Impl shared.PingPong
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

func toProto(b *shared.Ball) (out *pb.Ball, err error) {
	if b == nil {
		return nil, fmt.Errorf("nil ball")
	}

	out = new(pb.Ball)
	out.Player = b.Player

	lm := len(b.Metas)
	if lm > 0 {
		out.Metas = make([]*pb.Ball_Meta, lm)
		for i, mt := range b.Metas {
			out.Metas[i] = &pb.Ball_Meta{Key: mt.Key, Value: mt.Value}
		}
	}

	lr := len(b.Raw)
	if lr > 0 {
		out.Raw = make([]byte, lr)
		copy(out.Raw, b.Raw)
	}

	return
}

func fromProto(b *pb.Ball) (out *shared.Ball, err error) {
	if b == nil {
		return nil, fmt.Errorf("nil ball")
	}

	out = new(shared.Ball)
	out.Player = b.GetPlayer()

	lm := len(b.Metas)
	if lm > 0 {
		out.Metas = make([]shared.Meta, lm)
		for i, mt := range b.Metas {
			out.Metas[i] = shared.Meta{Key: mt.Key, Value: mt.Value}
		}
	}

	lr := len(b.Raw)
	if lr > 0 {
		out.Raw = make([]byte, lr)
		copy(out.Raw, b.Raw)
	}

	return
}
