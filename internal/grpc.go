// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal

import (
	"context"
	"fmt"

	"github.com/g41797/pingopong/internal/pb"
	"github.com/g41797/pingopong/pingpong"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

// gRPCClient is an implementation of PingPong that talks over GRPC.
type gRPCClient struct{ client pb.PingPongClient }

func (gc *gRPCClient) Play(ctx context.Context, in *pingpong.Ball) (*pingpong.Ball, error) {
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

// Here is the gRPC server that gRPCClient talks to.
type gRPCServer struct {
	// This is the real implementation
	Impl pingpong.PingPong
}

func (gs *gRPCServer) Play(ctx context.Context, pb *pb.Ball) (*pb.Ball, error) {
	b, _ := fromProto(pb)
	b, err := gs.Impl.Play(ctx, b)
	if err != nil {
		return nil, err
	}
	return toProto(b)
}

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type pingPongGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.Plugin
	// Concrete implementation, written in Go.
	Impl pingpong.PingPong
}

func (p *pingPongGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterPingPongServer(s, &gRPCServer{Impl: p.Impl})
	return nil
}

func (p *pingPongGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &gRPCClient{client: pb.NewPingPongClient(c)}, nil
}

const pingPongPluginName = "pingpong_grpc"

// handshake is a common handshake that is pingopong by plugin and host.
var handshake = plugin.HandshakeConfig{
	// This isn't required when using VersionedPlugins
	ProtocolVersion:  1,
	MagicCookieKey:   "UnaryGRPC",
	MagicCookieValue: "pingpong",
}

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	pingPongPluginName: &pingPongGRPCPlugin{},
}

func toProto(b *pingpong.Ball) (out *pb.Ball, err error) {
	if b == nil {
		return nil, fmt.Errorf("nil ball")
	}

	out = new(pb.Ball)
	out.Player = b.Player

	lm := len(b.Properties)
	if lm > 0 {
		out.Properties = make([]*pb.Ball_Property, lm)
		for i, mt := range b.Properties {
			out.Properties[i] = &pb.Ball_Property{Key: mt.Key, Value: mt.Value}
		}
	}

	lr := len(b.Raw)
	if lr > 0 {
		out.Raw = make([]byte, lr)
		copy(out.Raw, b.Raw)
	}

	return
}

func fromProto(b *pb.Ball) (out *pingpong.Ball, err error) {
	if b == nil {
		return nil, fmt.Errorf("nil ball")
	}

	out = new(pingpong.Ball)
	out.Player = b.GetPlayer()

	lm := len(b.Properties)
	if lm > 0 {
		out.Properties = make([]pingpong.Property, lm)
		for i, mt := range b.Properties {
			out.Properties[i] = pingpong.Property{Key: mt.Key, Value: mt.Value}
		}
	}

	lr := len(b.Raw)
	if lr > 0 {
		out.Raw = make([]byte, lr)
		copy(out.Raw, b.Raw)
	}

	return
}
