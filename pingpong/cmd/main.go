// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package main

import (
	"github.com/g41797/grpc-pingpong/pingpong"
)

func main() {
	if pingpong.IsPluginProcess() {
		pingpong.RunServer()
		return
	}
	pingpong.RunClient()
}
