// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package example

import (
	"context"
	"fmt"
	"os"

	"github.com/g41797/pingopong"
	"github.com/g41797/pingopong/api"
)

// Register factory during initialization
func init() {
	api.RegisterFactory("Echo", echoPlayerFactory)
}

func echoPlayerFactory(_ string) (pingopong.PingPongPlayer, error) {

	// Creation
	pppl := new(echoPlayer)

	// Initialization of plugin
	// without configuration
	if err := pppl.InitOnce(nil); err != nil {
		return nil, err
	}

	return pppl, nil
}

type echoPlayer struct {
}

func (p *echoPlayer) InitOnce(_ []byte) error {
	return nil
}

func (p *echoPlayer) FinishOnce() error {
	return nil
}

// Returns Player and Raw without changes
// Properties[0] - set to current process id
// Properties[1] - set to parent process id
func (p *echoPlayer) Play(_ context.Context, b *pingopong.Ball) (*pingopong.Ball, error) {

	b.Properties = make([]pingopong.Property, 2)

	// For DirectCall - Value contains the pid of client process
	b.Properties[0].Key = "PID"
	b.Properties[0].Value = fmt.Sprint(os.Getpid())

	// For in-directCall - Value contains the pid of client(parent) process
	b.Properties[1].Key = "PPID"
	b.Properties[1].Value = fmt.Sprint(os.Getppid())

	return b, nil
}
