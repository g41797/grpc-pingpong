// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package api

import (
	"github.com/g41797/pingopong"
	"github.com/g41797/pingopong/internal"
)

// Store factory for further usage
// name of player stored in lower case
func RegisterFactory(name string, fact pingopong.PingPongPlayerFactory) {
	internal.StoreFactory(name, fact)
}
