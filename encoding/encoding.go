package encoding

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/ghodss/yaml"
	jsonggpb "github.com/gogo/protobuf/jsonpb"
	protogg "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type errReader struct{ err error }

func (r errReader) Read(_ []byte) (int, error) { return 0, r.err }

func JSON(in interface{}) io.Reader {
	b, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal to JSON: %v", in)}
	}
	return bytes.NewBuffer(b)
}

// JSONPB should be used when serialising protobuf messages.
// The jsonpb marshaler behaves slightly differently to go's built in marshaler.
func JSONPB(in proto.Message) io.Reader {
	str, err := (&jsonpb.Marshaler{Indent: "  "}).MarshalToString(in.(proto.Message))
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal protobuf to JSONPB: %v", in)}
	}
	return bytes.NewBufferString(str)
}

func YAML(in ...interface{}) io.Reader {
	var concatDelim = []byte("---\n")

	if len(in) == 0 {
		return errReader{err: errors.New("Nothing to output")}
	}
	var res [][]byte
	for _, entry := range in {
		var entryBytes []byte

		// Do not marshal strings - they should be appended directly
		if extraString, ok := entry.(string); ok {
			entryBytes = []byte(extraString)
		} else {
			b, err := yaml.Marshal(in)
			if err != nil {
				return errReader{err: errors.Wrapf(err, "unable to marshal to YAML: %v", in)}
			}
			entryBytes = b
		}
		res = append(res, entryBytes)
	}

	if len(res) == 1 {
		return bytes.NewBuffer(res[0])
	}

	return bytes.NewBuffer(bytes.Join(res, concatDelim))
}

// GogoJSONPB should be used when serialising protobuf messages that were generated
// using gogo/protobuf library.
func GogoJSONPB(in protogg.Message) io.Reader {
	str, err := (&jsonggpb.Marshaler{Indent: " "}).MarshalToString(in.(protogg.Message))
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal protobuf to GogoJSONPB: %v", in)}
	}
	return bytes.NewBufferString(str)
}
