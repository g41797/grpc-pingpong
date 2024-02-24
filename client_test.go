// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingopong_test

import (
	"context"
	"testing"

	"github.com/g41797/pingopong"

	// Attach package with players to the process using
	// so called "blank import" :
	// for this kind of import only init() functions will be called
	_ "github.com/g41797/pingopong/example"

	"github.com/hashicorp/go-hclog"
)

func TestPingPongClient_Play(t *testing.T) {
	if pingopong.IsPluginProcess() {
		RunServer(t)
		return
	}

	RunClient(t)

	return
}

func RunServer(*testing.T) {
	pingopong.NewServer(hclog.Debug)()
	return
}

func RunClient(t *testing.T) {

	pcl, clean := pingopong.NewGame(hclog.Debug)

	if pcl == nil {
		t.Fatal("cannot create new game")
		return
	}

	t.Cleanup(clean)

	b := pingopong.Ball{Player: "echo"}

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
