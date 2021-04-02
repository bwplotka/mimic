// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// JSON returns reader that encodes anything to JSON.
func JSON(in interface{}) io.Reader {
	b, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal to JSON: %v", in)}
	}
	return bytes.NewBuffer(b)
}
