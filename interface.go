// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingopong

import (
	"context"
)

// PingPong is the interface that we're exposing as a visible part of plugin.
type PingPong interface {
	Play(ctx context.Context, b *Ball) (*Ball, error)
}

// PingPongPlayer - interface of server part
type PingPongPlayer interface {
	InitOnce(config []byte) error
	PingPong
	FinishOnce() error
}

// Creator of PingPongPlayer implementation "object"
type PingPongPlayerFactory func() (PingPongPlayer, error)
