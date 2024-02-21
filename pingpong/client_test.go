// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingpong_test

import (
	"context"
	"testing"

	"github.com/g41797/grpc-pingpong/pingpong"
	"github.com/g41797/grpc-pingpong/shared"
	"github.com/hashicorp/go-hclog"
)

func TestPingPongClient_Play(t *testing.T) {
	if pingpong.IsPluginProcess() {
		RunServer(t)
		return
	}

	RunClient(t)

	return
}

func RunServer(*testing.T) {
	pingpong.NewServer(hclog.Debug).Run()
	return
}

func RunClient(t *testing.T) {

	pcl, clean := pingpong.NewGame(hclog.Debug)

	if pcl == nil {
		t.Fatal("cannot create new game")
		return
	}

	t.Cleanup(clean)

	b := shared.Ball{Player: "noname"}

	res, err := pcl.Play(context.Background(), &b)
	if err != nil {
		t.Fatal(err)
		return
	}

	if res.Player != b.Player {
		t.Errorf("expected %s actual %s", b.Player, res.Player)
		return
	}

	return
}
