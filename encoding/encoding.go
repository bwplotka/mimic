package encoding

import (
	"bytes"
	"encoding/json"
	"io"

	ghodssyaml "github.com/ghodss/yaml"
	jsonggpb "github.com/gogo/protobuf/jsonpb"
	protogg "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	yaml2 "gopkg.in/yaml.v2"
)

type errReader struct{ err error }

func (r errReader) Read(_ []byte) (int, error) { return 0, r.err }

// JSON returns reader that encodes anything to JSON.
func JSON(in interface{}) io.Reader {
	b, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal to JSON: %v", in)}
	}
	return bytes.NewBuffer(b)
}

// JSONPB returns reader that encodes protobuf messages to JSON.
// NOTE: The jsonpb marshaler behaves slightly differently to go's built in marshaler.
func JSONPB(in proto.Message) io.Reader {
	str, err := (&jsonpb.Marshaler{Indent: "  "}).MarshalToString(in.(proto.Message))
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal protobuf to JSONPB: %v", in)}
	}
	return bytes.NewBufferString(str)
}

// GhodssYAML returns reader that encodes anything to YAML using github.com/ghodss/yaml.
// Desired for e.g:
// * Kubernetes
func GhodssYAML(in ...interface{}) io.Reader {
	return yaml(ghodssyaml.Marshal, in...)
}

// YAML returns reader that encodes anything to YAML using gopkg.in/yaml.v2.
// Desired for e.g:
// * Prometheus, Alertmanager configuration
func YAML(in ...interface{}) io.Reader {
	return yaml(yaml2.Marshal, in...)
}

type MarshalFunc func(o interface{}) ([]byte, error)

func yaml(marshalFn MarshalFunc, in ...interface{}) io.Reader {
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
			b, err := marshalFn(entry)
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

// GogoJSONPB returns reader that encodes protobuf messages built with gogo/protobuf implementation.
func GogoJSONPB(in protogg.Message) io.Reader {
	str, err := (&jsonggpb.Marshaler{Indent: " "}).MarshalToString(in.(protogg.Message))
	if err != nil {
		return errReader{err: errors.Wrapf(err, "unable to marshal protobuf to GogoJSONPB: %v", in)}
	}
	return bytes.NewBufferString(str)
}
