// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package shared

import (
	"context"

	"github.com/g41797/grpc-pingpong/pb"
	"google.golang.org/grpc"
)

// GRPCClient is an implementation of PingPong that talks over GRPC.
type GRPCClient struct{ client pb.PingPongClient }

func (gc *GRPCClient) Play(ctx context.Context, in *pb.Ball, opts ...grpc.CallOption) (*pb.Ball, error) {
	return nil, nil
}

// Here is the gRPC server that GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl PingPong
}

func (gs *GRPCServer) Play(ctx context.Context, b *pb.Ball) (*pb.Ball, error) {
	return nil, nil
}
