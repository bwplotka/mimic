// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// JSONPB returns reader that encodes protobuf messages to JSON.
// NOTE: The jsonpb marshaler behaves slightly differently to go's built in marshaler.
func JSONPB(in proto.Message) io.Reader {
	b, err := (&protojson.MarshalOptions{Indent: "  "}).Marshal(in.(proto.Message))
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal protobuf to JSONPB: %v", in)}
	}
	return bytes.NewBuffer(b)
}
