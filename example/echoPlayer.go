// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package example

import (
	"context"

	"github.com/g41797/pingopong/api"
	"github.com/g41797/pingopong/pingpong"
)

func init() {
	api.RegisterFactory("Echo", echoPlayerFactory)
}

func echoPlayerFactory(_ string) (pingpong.PingPongPlayer, error) {
	pppl := new(echoPlayer)
	if err := pppl.InitOnce(nil); err != nil {
		return nil, err
	}
	return pppl, nil
}

var _ pingpong.PingPongPlayer = &echoPlayer{}

type echoPlayer struct {
}

func (p *echoPlayer) InitOnce(config []byte) error {
	return nil
}

func (p *echoPlayer) FinishOnce() error {
	return nil
}

func (p *echoPlayer) Play(ctx context.Context, b *pingpong.Ball) (*pingpong.Ball, error) {
	return b, nil
}
