// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	ghodssyaml "github.com/ghodss/yaml"
	yaml2 "gopkg.in/yaml.v3"
	yaml3 "gopkg.in/yaml.v3"
)

// GhodssYAML returns reader that encodes anything to YAML using github.com/ghodss/yaml.
// It works by first marshalling to JSON, so no `yaml` directive will work (it accepts `json` though).
//
// Recommended for:
// * Kubernetes
func GhodssYAML(in ...interface{}) io.Reader {
	return yaml(ghodssyaml.Marshal, in...)
}

// YAML returns reader that encodes anything to YAML using gopkg.in/yaml.v3.
// NOTE: Indentations are currently "weird": https://github.com/go-yaml/yaml/issues/661
func YAML(in ...interface{}) io.Reader {
	return yaml(yaml3.Marshal, in...)
}

// YAML2 returns reader that encodes anything to YAML using gopkg.in/yaml.v2.
// NOTE: Indentations are currently "weird": https://github.com/go-yaml/yaml/issues/661
// Recommended for:
// * Prometheus, Alertmanager configuration
func YAML2(in ...interface{}) io.Reader {
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
				return errReader{err: fmt.Errorf("unable to marshal to YAML: %v: %w", in, err)}
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
