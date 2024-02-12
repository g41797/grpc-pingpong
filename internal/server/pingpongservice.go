// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package server

import (
	"github.com/g41797/grpc-pingpong/pb"
	"google.golang.org/grpc"
)

var _ pb.PingPongServiceServer = (*PingPongService)(nil)

type PingPongService struct {
	pb.UnimplementedPingPongServiceServer
}

func CreatePingPongServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterPingPongServiceServer(server, &PingPongService{})
	return server
}

func (srv *PingPongService) Play(p pb.PingPongService_PlayServer) error {
	return nil
}
