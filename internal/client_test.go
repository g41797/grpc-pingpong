// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal_test

import (
	"context"
	"testing"

	// Attach package with players to the process using
	// so called "blank import" :
	// for this kind of import only init() functions will be called
	"github.com/g41797/pingopong/api"
	_ "github.com/g41797/pingopong/example"
	"github.com/g41797/pingopong/internal"
	"github.com/g41797/pingopong/pingpong"

	"github.com/hashicorp/go-hclog"
)

func TestPingPongClient_Play(t *testing.T) {
	if internal.IsPluginProcess() {
		RunServer(t)
		return
	}

	RunClient(t)

	return
}

func RunServer(*testing.T) {
	internal.NewServer(hclog.Debug)()
	return
}

func RunClient(t *testing.T) {

	pcl, clean := api.NewGame(hclog.Debug)

	if pcl == nil {
		t.Fatal("cannot create new game")
		return
	}

	t.Cleanup(clean)

	b := pingpong.Ball{Player: "echo"}

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
