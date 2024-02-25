// Copyright (c) 2024 g41797
// SPDX-License-Identifier: MIT

package pingpong

// Analog of http header or meta data
type Property struct {
	Key   string `json:"Key,omitempty"`
	Value string `json:"Value,omitempty"`
}

// Main unit of information
type Ball struct {
	// Name of the player (handler) - must
	Player string `json:"Player"`
	// Optional headers/meta data
	Properties []Property `json:"Properties,omitempty"`
	// Optional content, usually JSON byte array
	Raw []byte `json:"Raw,omitempty"`
}
