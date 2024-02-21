// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package shared

type Meta struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

type Ball struct {
	Player string `json:"Player,omitempty"`
	Metas  []Meta `json:"Metas,omitempty"`
	Raw    []byte `json:"Raw,omitempty"`
}
