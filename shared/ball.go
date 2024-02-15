// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package shared

import "github.com/g41797/grpc-pingpong/pb"

type Meta struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

type Ball struct {
	Player string `json:"Player,omitempty"`
	Metas  []Meta `json:"Metas,omitempty"`
	Raw    []byte `json:"Raw,omitempty"`
}

func ToProto(b *Ball) (*pb.Ball, error) {
	return nil, nil
}

func FromProto(b *pb.Ball) (*Ball, error) {
	return nil, nil
}
