// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/g41797/pingopong/pingpong"
)

var _ pingpong.PingPongPlayer = &guard{}

type guardState int

const (
	processfinishallowed guardState = iota
	nothingallowed       guardState = iota + 1
)

func (gst guardState) String() string {
	return []string{"initallowed", "processfinishallowed", "nothingallowed"}[gst]
}

type guard struct {
	lock  sync.Mutex
	state guardState
	name  string
	pl    pingpong.PingPongPlayer
}

func (grd *guard) tryCreate(name string) error {

	fact, exists := factories.Load(strings.ToLower(name))
	if !exists {
		return fmt.Errorf("factory for %s does not exist", name)
	}

	if pl, err := fact.(pingpong.PingPongPlayerFactory)(name); err == nil {
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

func (grd *guard) Play(ctx context.Context, b *pingpong.Ball) (*pingpong.Ball, error) {
	if grd == nil {
		return nil, fmt.Errorf("Process nil guard")
	}

	grd.lock.Lock()
	defer grd.lock.Unlock()

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

	grd.lock.Lock()
	defer grd.lock.Unlock()

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

func StoreFactory(name string, fact pingpong.PingPongPlayerFactory) {
	if len(name) == 0 {
		log.Panic("empty player name")
	}
	if fact == nil {
		log.Panicf("nil player factory for %s", name)
	}

	if _, exists := factories.LoadOrStore(strings.ToLower(name), fact); exists {
		log.Panicf("player factory for %s already exists", name)
	}
	return
}

var factories sync.Map
