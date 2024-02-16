// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package shared

import (
	"context"

	"github.com/g41797/grpc-pingpong/pb"
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
