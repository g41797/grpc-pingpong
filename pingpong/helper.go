// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingpong

import (
	_ "github.com/g41797/grpc-pingpong/internal"
	"github.com/g41797/grpc-pingpong/shared"
	"github.com/hashicorp/go-hclog"
)

func NewGame(trl hclog.Level) (shared.PingPong, func()) {
	result := &internal.gclient{level: trl}
	return result, result.Clean
}

func NewServer(trl hclog.Level) func() {
	srv := internal.gserver{level: trl}
	return srv.Run
}
