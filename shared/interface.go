// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package shared

import (
	"context"
)

// PingPong is the interface that we're exposing as a visible part of plugin.
type PingPong interface {
	Play(ctx context.Context, b *Ball) (*Ball, error)
}

// PingPongHandler - interface of server part
type PingPongHandler interface {
	InitOnce(config []byte) error
	PingPong
	FinishOnce() error
}

// Creator of PingPongHandler implementation "object"
type PingPongHandlerFactory func() (PingPongHandler, error)
