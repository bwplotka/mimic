// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/rodaine/hclencoder"
)

// hclEncoder implements the Encoder interface.
type hclEncoder struct {
	io.Reader
}

// EncodeComment returns byte slice that represents a HCL comment.
// We split `lines` by U+000A and encode as a single/multi line comment.
func (hclEncoder) EncodeComment(lines string) []byte {
	commentLines := strings.Split(lines, "\n")

	finalString := ""
	for _, comment := range commentLines {
		if comment == "" {
			continue
		}

		if finalString == "" {
			finalString = "# " + strings.TrimLeft(comment, " ")
		} else {
			finalString = finalString + "\n" + "# " + strings.TrimLeft(comment, " ")
		}
	}

	if finalString == "" {
		return []byte{}
	}

	finalString = finalString + "\n"
	return []byte(finalString)
}

func HCL(in interface{}) hclEncoder {
	b, err := hclencoder.Encode(in)
	if err != nil {
		return hclEncoder{Reader: errReader{err: fmt.Errorf("unable to marshal to HCL: %v: %w", in, err)}}
	}
	return hclEncoder{Reader: bytes.NewBuffer(b)}
}
