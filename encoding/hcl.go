// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"bytes"
	"fmt"
	"io"

	"github.com/rodaine/hclencoder"
)

// hclEncoder implements the Encoder interface.
type hclEncoder struct {
	io.Reader
}

// EncodeComment returns byte slice that represents a HCL comment (same as YAML).
// We split `lines` by '\n' and encode as a single/multi line comment.
func (hclEncoder) EncodeComment(lines string) []byte {
	return YAML().EncodeComment(lines)
}

func HCL(in interface{}) hclEncoder {
	b, err := hclencoder.Encode(in)
	if err != nil {
		return hclEncoder{Reader: errReader{err: fmt.Errorf("unable to marshal to HCL: %v: %w", in, err)}}
	}
	return hclEncoder{Reader: bytes.NewBuffer(b)}
}
