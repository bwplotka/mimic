// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	ghodssyaml "github.com/ghodss/yaml"
	yaml2 "gopkg.in/yaml.v3"
	yaml3 "gopkg.in/yaml.v3"
)

// yamlEncoder implements the Encoder interface.
type yamlEncoder struct {
	io.Reader
}

// EncodeComment returns byte slice that represents a YAML comment.
// We split `lines` by '\n' and encode as a single/multi line comment.
func (yamlEncoder) EncodeComment(lines string) []byte {
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

// GhodssYAML returns reader that encodes anything to YAML using github.com/ghodss/yaml.
// It works by first marshalling to JSON, so no `yaml` directive will work (it accepts `json` though).
//
// Recommended for:
// * Kubernetes
func GhodssYAML(in ...interface{}) yamlEncoder {
	return yaml(ghodssyaml.Marshal, in...)
}

// YAML returns reader that encodes anything to YAML using gopkg.in/yaml.v3.
// NOTE: Indentations are currently "weird": https://github.com/go-yaml/yaml/issues/661
func YAML(in ...interface{}) yamlEncoder {
	return yaml(yaml3.Marshal, in...)
}

// YAML2 returns reader that encodes anything to YAML using gopkg.in/yaml.v2.
// NOTE: Indentations are currently "weird": https://github.com/go-yaml/yaml/issues/661
// Recommended for:
// * Prometheus, Alertmanager configuration
func YAML2(in ...interface{}) yamlEncoder {
	return yaml(yaml2.Marshal, in...)
}

type MarshalFunc func(o interface{}) ([]byte, error)

func yaml(marshalFn MarshalFunc, in ...interface{}) yamlEncoder {
	var concatDelim = []byte("---\n")

	if len(in) == 0 {
		return yamlEncoder{Reader: errReader{err: errors.New("nothing to output")}}
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
				return yamlEncoder{Reader: errReader{err: fmt.Errorf("unable to marshal to YAML: %v: %w", in, err)}}
			}
			entryBytes = b
		}
		res = append(res, entryBytes)
	}

	if len(res) == 1 {
		return yamlEncoder{
			Reader: bytes.NewBuffer(res[0]),
		}
	}

	return yamlEncoder{
		Reader: bytes.NewBuffer(bytes.Join(res, concatDelim)),
	}
}
