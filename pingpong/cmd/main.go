// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"fmt"

	"github.com/g41797/grpc-pingpong/pingpong"
	"github.com/g41797/grpc-pingpong/shared"
	"github.com/hashicorp/go-hclog"
)

func main() {
	if pingpong.IsPluginProcess() {
		RunServer()
		return
	}
	RunClient()
}

func RunClient() {

	pcl := pingpong.NewClient(hclog.Debug)
	defer pcl.Clean()

	b := shared.Ball{Player: "noname"}

	res, err := pcl.Play(context.Background(), &b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if res.Player != b.Player {
		fmt.Printf("expected %s actual %s", b.Player, res.Player)
		return
	}

	return
}

func RunServer() {
	srv := pingpong.NewServer(hclog.Debug)
	srv.Run()
}
