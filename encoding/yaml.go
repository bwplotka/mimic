package encoding

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	yaml2 "gopkg.in/yaml.v2"
	ghodssyaml "github.com/ghodss/yaml"
)

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
				return errReader{err: errors.New(fmt.Sprintf("unable to marshal to YAML: %v", in))}
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

