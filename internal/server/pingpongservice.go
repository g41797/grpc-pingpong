// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package server

import (
	"context"

	"github.com/g41797/grpc-pingpong/pb"
	"google.golang.org/grpc"
)

var _ pb.PingPongServiceServer = (*PingPongService)(nil)

type PingPongService struct {
}

func CreateGrpcPingPongServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterPingPongServiceServer(server, &PingPongService{})
	return server
}

func (srv *PingPongService) Play(ctx context.Context, b *pb.Ball) (*pb.Ball, error) {
	return nil, nil
}
