//go:build !jsoniter
// +build !jsoniter

package json

import "encoding/json"

// RawMessage is exported by component-base/pkg/json package.
type RawMessage = json.RawMessage

var (
	// Marshal is exported by component-base/pkg/json package.
	Marshal = json.Marshal
	// UnMarshal is exported by component-base/pkg/json package.
	UnMarshal = json.Unmarshal
	// MarshalIndent is exported by component-base/pkg/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by component-base/pkg/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by component-base/pkg/json package.
	NewEncoder = json.NewEncoder
)
