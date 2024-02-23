// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingopong

import (
	"os"

	"github.com/mitchellh/go-ps"
)

// Returns true if called within plugin process
func IsPluginProcess() bool {
	return RunningExeName() == parentExeName()
}

// Executable name running this process. This is not a path to the
// executable.
func RunningExeName() string {
	proc, _ := ps.FindProcess(os.Getpid())
	return proc.Executable()
}

func parentExeName() string {
	proc, _ := ps.FindProcess(os.Getppid())
	return proc.Executable()
}
