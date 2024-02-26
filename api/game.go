// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package api

import (
	"os"

	"github.com/g41797/pingopong"
	"github.com/g41797/pingopong/internal"
	"github.com/hashicorp/go-hclog"
)

// Prepares new game with log level 'trl'
// Returns:
// - PingPong interface for the play
// - Cleanup function - should be called after finish of the game.
func NewGame(trl hclog.Level) (pingopong.PingPong, func()) {

	if IsDirectCall() {
		return internal.NewPingPing(trl)
	}

	return internal.NewClient(trl)
}

// There are two modes of "game":
// - DirectCall   - plugin is created within the same process
// - IndirectCall - plugin is created withing child process
// Default mode - 'IndirectCall'
// IsDirectCall functions returns current mode
func IsDirectCall() bool {

	val, exsists := os.LookupEnv(directCALL)

	if !exsists || val != directCALLON {
		return false
	}

	return true
}

const (
	directCALL    = "DIRECTCALL"
	directCALLON  = "ON"
	directCALLOFF = "OFF"
)

// Set DirectCall mode - as a rule for debug purposes
func SetDirectCall() {
	os.Setenv(directCALL, directCALLON)
}

// Reset DirectCall mode
func ResetDirectCall() {
	os.Setenv(directCALL, directCALLOFF)
}
