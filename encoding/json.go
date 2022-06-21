// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// JSON returns reader that encodes anything to JSON.
func JSON(in interface{}) io.Reader {
	b, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return errReader{err: fmt.Errorf("unable to marshal to JSON: %v: %w", in, err)}
	}
	return bytes.NewBuffer(b)
}
