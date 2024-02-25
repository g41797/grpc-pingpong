// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package api

import (
	"github.com/g41797/pingopong/internal"
	"github.com/g41797/pingopong/pingpong"
	"github.com/hashicorp/go-hclog"
)

// Prepares new game with log level 'trl'
// Returns PingPong interface for the play and cleanup function
// for calling after finish of the game.
func NewGame(trl hclog.Level) (pingpong.PingPong, func()) {
	result := &internal.Gclient{Level: trl}
	return result, result.Clean
}
