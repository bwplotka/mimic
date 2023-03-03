// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"io"
	"testing"

	"github.com/efficientgo/tools/core/pkg/testutil"
)

type A struct {
	Field1     string `yaml:"FieldYolo1"`
	Field2     int    `yaml:",omitempty"`
	Inner      *A
	InnerSlice []A `yaml:",omitempty"`
}

var testA = []interface{}{
	A{
		Field1: "1",
		Field2: 1,
		Inner: &A{
			Field1: "inner1",
			Field2: 11,
		},
	},
	A{
		Field1: "2",
		Field2: 2,
		InnerSlice: []A{
			{Field2: 3},
			{Field2: 3},
		},
	},
}

func TestYaml_EncodingToStructs(t *testing.T) {
	for _, tcase := range []struct {
		encoder  io.Reader
		expected string
	}{
		{
			encoder: GhodssYAML(testA...),
			expected: `Field1: "1"
Field2: 1
Inner:
  Field1: inner1
  Field2: 11
  Inner: null
  InnerSlice: null
InnerSlice: null
---
Field1: "2"
Field2: 2
Inner: null
InnerSlice:
- Field1: ""
  Field2: 3
  Inner: null
  InnerSlice: null
- Field1: ""
  Field2: 3
  Inner: null
  InnerSlice: null
`,
		},
		{
			encoder: YAML(testA...),
			expected: `FieldYolo1: "1"
field2: 1
inner:
    FieldYolo1: inner1
    field2: 11
    inner: null
---
FieldYolo1: "2"
field2: 2
inner: null
innerslice:
    - FieldYolo1: ""
      field2: 3
      inner: null
    - FieldYolo1: ""
      field2: 3
      inner: null
`,
		},
		{
			encoder: YAML(testA...),
			expected: `FieldYolo1: "1"
field2: 1
inner:
    FieldYolo1: inner1
    field2: 11
    inner: null
---
FieldYolo1: "2"
field2: 2
inner: null
innerslice:
    - FieldYolo1: ""
      field2: 3
      inner: null
    - FieldYolo1: ""
      field2: 3
      inner: null
`,
		},
	} {
		if ok := t.Run("", func(t *testing.T) {
			actual, err := io.ReadAll(tcase.encoder)
			testutil.Ok(t, err)
			testutil.Equals(t, tcase.expected, string(actual))
		}); !ok {
			return
		}
	}
}

func TestHCL_EncodingToStructs(t *testing.T) {
	type Inner struct {
		Key    string `hcl:",key"`
		Field1 string `hcl:"field1"`
		Field2 int    `hcl:"field2"`
	}

	actual, err := io.ReadAll(HCL(
		struct {
			Inner `hcl:"inner"`
		}{Inner{
			Key:    "test",
			Field1: "first",
			Field2: 12,
		},
		},
	))
	testutil.Ok(t, err)
	testutil.Equals(t, `inner "test" {
  field1 = "first"
  field2 = 12
}
`, string(actual))
}

func TestEncodingComments(t *testing.T) {
	for _, tcase := range []struct {
		name     string
		encoder  Encoder
		comment  string
		expected string
	}{
		{
			name:     "single line ghodss",
			encoder:  GhodssYAML(),
			comment:  "This is a single line comment.",
			expected: "# This is a single line comment.\n",
		},
		{
			name:     "multi line ghodss",
			encoder:  GhodssYAML(),
			comment:  "This is a multi\n line comment.",
			expected: "# This is a multi\n# line comment.\n",
		},
		{
			name:     "single line yaml2",
			encoder:  YAML2(),
			comment:  "This is a single line comment.",
			expected: "# This is a single line comment.\n",
		},
		{
			name:     "multi line yaml2",
			encoder:  YAML2(),
			comment:  "This is a multi\n line comment.",
			expected: "# This is a multi\n# line comment.\n",
		},
		{
			name:     "single line hcl",
			encoder:  HCL(testA),
			comment:  "This is a single line comment.",
			expected: "# This is a single line comment.\n",
		},
		{
			name:     "multi line hcl",
			encoder:  HCL(testA),
			comment:  "This is a multi\n line comment.",
			expected: "# This is a multi\n# line comment.\n",
		},
	} {
		if ok := t.Run(tcase.name, func(t *testing.T) {
			actual := tcase.encoder.EncodeComment(tcase.comment)
			testutil.Equals(t, tcase.expected, string(actual))
		}); !ok {
			return
		}
	}
}
