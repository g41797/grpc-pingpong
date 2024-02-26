// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package internal

import (
	"context"
	"strings"
	"sync"

	"github.com/g41797/pingopong"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func NewPingPing(trl hclog.Level) (pingopong.PingPong, func()) {
	srv := gserver{level: trl}
	return &srv, srv.Clean
}

func NewServer(trl hclog.Level) func() {
	srv := gserver{level: trl}
	return srv.Run
}

var _ pingopong.PingPong = (*gserver)(nil)

type gserver struct {
	lock  sync.Mutex
	level hclog.Level
	g     *guard
}

func (s *gserver) Run() {

	defer s.Clean()

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshake,
		Plugins: map[string]plugin.Plugin{
			pingPongPluginName: &pingPongGRPCPlugin{Impl: s},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
		Logger: hclog.New(&hclog.LoggerOptions{
			Output: hclog.DefaultOutput,
			Level:  s.level,
			Name:   RunningExeName() + "_plugin",
		}),
	})

}

func (s *gserver) Play(ctx context.Context, b *pingopong.Ball) (*pingopong.Ball, error) {

	s.lock.Lock()
	defer s.lock.Unlock()

	if s.g == nil {
		s.g = &guard{}
	}

	name := strings.ToLower(b.Player)
	if s.g.name == name {
		return s.g.Play(ctx, b)
	}

	s.g.FinishOnce()
	s.g = nil

	var g guard
	if err := g.tryCreate(name); err != nil {
		return nil, err
	}

	s.g = &g
	return s.g.Play(ctx, b)
}

func (s *gserver) Clean() {
	if s == nil {
		return
	}
	if s.g != nil {
		s.g.FinishOnce()
	}
}
