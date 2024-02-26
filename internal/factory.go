// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal

import (
	"log"
	"strings"
	"sync"

	"github.com/g41797/pingopong"
)

func StoreFactory(name string, fact pingopong.PingPongPlayerFactory) {
	if len(name) == 0 {
		log.Panic("empty player name")
	}
	if fact == nil {
		log.Panicf("nil player factory for %s", name)
	}

	if _, exists := Factories.LoadOrStore(strings.ToLower(name), fact); exists {
		log.Panicf("player factory for %s already exists", name)
	}
	return
}

var Factories sync.Map
