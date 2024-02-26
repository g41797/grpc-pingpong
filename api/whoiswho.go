// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package api

import "github.com/g41797/pingopong/internal"

// Returns true if called within plugin process
func IsPluginProcess() bool {
	return internal.RunningExeName() == internal.ParentExeName()
}
