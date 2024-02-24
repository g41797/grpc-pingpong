// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package example

import (
	"context"

	"github.com/g41797/pingopong"
)

func init() {
	pingopong.RegisterFactory("Echo", echoPlayerFactory)
}

func echoPlayerFactory(_ string) (pingopong.PingPongPlayer, error) {
	pppl := new(echoPlayer)
	if err := pppl.InitOnce(nil); err != nil {
		return nil, err
	}
	return pppl, nil
}

var _ pingopong.PingPongPlayer = &echoPlayer{}

type echoPlayer struct {
}

func (p *echoPlayer) InitOnce(config []byte) error {
	return nil
}

func (p *echoPlayer) FinishOnce() error {
	return nil
}

func (p *echoPlayer) Play(ctx context.Context, b *pingopong.Ball) (*pingopong.Ball, error) {
	return b, nil
}
