// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// jsonEncoder implements the Encoder interface.
type jsonEncoder struct {
	io.Reader
}

// EncodeComment is a no-op for JSON encoder, as JSON doesn't support comments.
func (jsonEncoder) EncodeComment(lines string) []byte {
	return []byte{}
}

// JSON returns reader that encodes anything to JSON.
func JSON(in interface{}) jsonEncoder {
	b, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return jsonEncoder{Reader: errReader{err: fmt.Errorf("unable to marshal to JSON: %v: %w", in, err)}}
	}

	return jsonEncoder{Reader: bytes.NewBuffer(b)}
}
