// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal

import (
	"context"
	"fmt"
	"strings"

	"github.com/g41797/pingopong"
)

var _ pingopong.PingPongPlayer = &guard{}

type guardState int

const (
	processfinishallowed guardState = iota
	nothingallowed       guardState = iota + 1
)

func (gst guardState) String() string {
	return []string{"initallowed", "processfinishallowed", "nothingallowed"}[gst]
}

type guard struct {
	state guardState
	name  string
	pl    pingopong.PingPongPlayer
}

func (grd *guard) tryCreate(name string) error {

	fact, exists := Factories.Load(strings.ToLower(name))
	if !exists {
		return fmt.Errorf("factory for %s does not exist", name)
	}

	if pl, err := fact.(pingopong.PingPongPlayerFactory)(name); err == nil {
		grd.pl = pl
		grd.state = processfinishallowed
		grd.name = name
		return nil
	} else {
		return err
	}
}

func (grd *guard) InitOnce(config []byte) error {
	return fmt.Errorf("init disabled")
}

func (grd *guard) Play(ctx context.Context, b *pingopong.Ball) (*pingopong.Ball, error) {
	if grd == nil {
		return nil, fmt.Errorf("Process nil guard")
	}

	if grd.pl == nil {
		return nil, fmt.Errorf("player was not created")
	}

	if grd.state != processfinishallowed {
		return nil, fmt.Errorf("Process disabled for %s", grd.state.String())
	}

	if ctx == nil {
		ctx = context.Background()
	}

	return grd.pl.Play(ctx, b)
}

func (grd *guard) FinishOnce() error {
	if grd == nil {
		return nil
	}

	if grd.pl == nil {
		return nil
	}

	if grd.state != processfinishallowed {
		return fmt.Errorf("Finish disabled for %s", grd.state.String())
	}

	err := grd.pl.FinishOnce()

	grd.state = nothingallowed

	return err
}
