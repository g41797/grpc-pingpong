// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	// Attach package with players to the process using
	// so called "blank import" :
	// for this kind of import only init() functions will be called
	_ "github.com/g41797/pingopong/example"

	"github.com/g41797/pingopong/api"
	"github.com/g41797/pingopong/internal"
	"github.com/g41797/pingopong/pingpong"

	"github.com/hashicorp/go-hclog"
)

func TestPingPongClient_Play(t *testing.T) {

	api.ResetDirectCall()

	if api.IsPluginProcess() {
		RunServer(t)
		return
	}

	RunClient(t)

	return
}

func TestPingPongDirectCall_Play(t *testing.T) {
	api.SetDirectCall()
	defer api.ResetDirectCall()
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

	if len(res.Properties) <= 2 {
		return
	}

	propPID := res.Properties[0]
	propPPID := res.Properties[1]

	spid := fmt.Sprint(os.Getpid())
	sppid := fmt.Sprint(os.Getppid())

	if api.IsDirectCall() {
		if propPID.Value != spid {
			t.Errorf("direct call within the same process: expected %s actual %s", spid, propPID.Value)
			return
		}
	}

	if propPPID.Value != sppid {
		t.Errorf("in-direct call within child process expected %s actual %s", spid, propPPID.Value)
		return
	}

	return
}
