// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
	"github.com/rodaine/hclencoder"
)

func HCL(in interface{}) io.Reader {
	b, err := hclencoder.Encode(in)
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal to HCL: %v", in)}
	}
	return bytes.NewBuffer(b)
}
