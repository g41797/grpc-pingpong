// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package shared

import (
	"fmt"

	"github.com/g41797/grpc-pingpong/pb"
)

type Meta struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

type Ball struct {
	Player string `json:"Player,omitempty"`
	Metas  []Meta `json:"Metas,omitempty"`
	Raw    []byte `json:"Raw,omitempty"`
}

func ToProto(b *Ball) (out *pb.Ball, err error) {
	if b == nil {
		return nil, fmt.Errorf("nil ball")
	}

	out = new(pb.Ball)
	out.Player = b.Player

	lm := len(b.Metas)
	if lm > 0 {
		out.Metas = make([]*pb.Ball_Meta, lm)
		for i, mt := range b.Metas {
			out.Metas[i] = &pb.Ball_Meta{Key: mt.Key, Value: mt.Value}
		}
	}

	lr := len(b.Raw)
	if lr > 0 {
		out.Raw = make([]byte, lr)
		copy(out.Raw, b.Raw)
	}

	return
}

func FromProto(b *pb.Ball) (out *Ball, err error) {
	if b == nil {
		return nil, fmt.Errorf("nil ball")
	}

	out = new(Ball)
	out.Player = b.GetPlayer()

	lm := len(b.Metas)
	if lm > 0 {
		out.Metas = make([]Meta, lm)
		for i, mt := range b.Metas {
			out.Metas[i] = Meta{Key: mt.Key, Value: mt.Value}
		}
	}

	lr := len(b.Raw)
	if lr > 0 {
		out.Raw = make([]byte, lr)
		copy(out.Raw, b.Raw)
	}

	return
}
