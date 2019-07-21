package encoding

import (
	"bytes"
	"io"

	jsonggpb "github.com/gogo/protobuf/jsonpb"
	protogg "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// JSONPB returns reader that encodes protobuf messages to JSON.
// NOTE: The jsonpb marshaler behaves slightly differently to go's built in marshaler.
func JSONPB(in proto.Message) io.Reader {
	str, err := (&jsonpb.Marshaler{Indent: "  "}).MarshalToString(in.(proto.Message))
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal protobuf to JSONPB: %v", in)}
	}
	return bytes.NewBufferString(str)
}

// GogoJSONPB returns reader that encodes protobuf messages built with gogo/protobuf implementation.
func GogoJSONPB(in protogg.Message) io.Reader {
	str, err := (&jsonggpb.Marshaler{Indent: " "}).MarshalToString(in.(protogg.Message))
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal protobuf to GogoJSONPB: %v", in)}
	}
	return bytes.NewBufferString(str)
}
