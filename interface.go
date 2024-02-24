// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingopong

import (
	"context"
)

// PingPong is the interface that we're exposing as a visible part of plugin.
type PingPong interface {
	// ctx used for external cancel of the 'game'.
	// how to use it see:
	// https://www.sohamkamani.com/golang/context/
	// https://www.willem.dev/articles/context-cancellation-explained/
	// https://www.willem.dev/articles/context-todo-or-background/
	// If ctx is not initialized, context.Background() will be used
	// cancel is valid flow, don't return error for this case
	// Content of returned Ball should reflect cancellation
	//
	// error should be returned for the cases when continue of the game is impossible.
	//
	Play(ctx context.Context, b *Ball) (*Ball, error)
}

// PingPongPlayer - interface of server part
type PingPongPlayer interface {
	// Initiation  - once for lifecycle.
	// Called by factory function provided by developed during creation
	// config (may be empty) - usually JSON byte array with configuration.
	InitOnce(config []byte) error

	PingPong

	// Stop processing, clean resources - once for lifecycle.
	// Called by infrastructure
	FinishOnce() error
}

// Creates PingPongPlayer implementation
// Calls InitOnce
// err is InitOnce error
// name converted to lower case befor call
type PingPongPlayerFactory func(name string) (player PingPongPlayer, err error)

// In order to allow creation of PingPongPlayer implementation:
// 1 - 	PingPongPlayerFactory should be provided for every player in the process
// 2 - 	PingPongPlayerFactory should be registered before creation via init()
//
//	 func init() {
//			pingopong.RegisterFactory("CopyFiles", cpfFactory)
//		}

// Store factory for further usage
// name of player stored in lower case
func RegisterFactory(name string, fact PingPongPlayerFactory) {
	storeFactory(name, fact)
}
