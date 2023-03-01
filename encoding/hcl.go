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

// Commenter adds comment strings at the top of a HCL file.
// Each string in comments slice is treated as a new comment.
func (hclEncoder) Commenter(b []byte, comments []string) []byte {
	finalString := ""
	for _, comment := range comments {
		if comment == "" {
			continue
		}

		if finalString == "" {
			finalString = "# " + comment
		} else {
			finalString = finalString + "\n" + "# " + comment
		}
	}

	if finalString == "" {
		return b
	}

	finalString = finalString + "\n"
	b = append([]byte(finalString), b...)
	return b
}

func HCL(in interface{}) hclEncoder {
	b, err := hclencoder.Encode(in)
	if err != nil {
		return hclEncoder{Reader: errReader{err: fmt.Errorf("unable to marshal to HCL: %v: %w", in, err)}}
	}
	return hclEncoder{Reader: bytes.NewBuffer(b)}
}
