// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingopong

type Property struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

type Ball struct {
	Player     string     `json:"Player,omitempty"`
	Properties []Property `json:"Properties,omitempty"`
	Raw        []byte     `json:"Raw,omitempty"`
}
