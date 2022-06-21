// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"bytes"
	"fmt"
	"io"

	"github.com/rodaine/hclencoder"
)

func HCL(in interface{}) io.Reader {
	b, err := hclencoder.Encode(in)
	if err != nil {
		return errReader{err: fmt.Errorf("unable to marshal to HCL: %v: %w", in, err)}
	}
	return bytes.NewBuffer(b)
}
