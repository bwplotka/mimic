// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import "io"

// Encoder implements the needed functions to encode a Go struct to a particular config language.
type Encoder interface {
	io.Reader

	// EncodeComment returns a slice of bytes that represents `lines` as a comment string
	// in a particular config language. `lines` can be a single or multiple line comment
	EncodeComment(lines string) []byte
}
