// Copyright (c) bwplotka/mimic Authors
// Licensed under the Apache License 2.0.

package encoding

import (
	"io/ioutil"
	"testing"

	"github.com/efficientgo/tools/core/pkg/testutil"
)

func TestYaml_EncodingToStructs(t *testing.T) {
	type A struct {
		Field1 string
		Field2 int
		Inner  *A
	}

	actual, err := ioutil.ReadAll(GhodssYAML(
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
		},
	))
	testutil.Ok(t, err)
	testutil.Equals(t, `Field1: "1"
Field2: 1
Inner:
  Field1: inner1
  Field2: 11
  Inner: null
---
Field1: "2"
Field2: 2
Inner: null
`, string(actual))
}

func TestHCL_EncodingToStructs(t *testing.T) {
	type Inner struct {
		Key    string `hcl:",key"`
		Field1 string `hcl:"field1"`
		Field2 int    `hcl:"field2"`
	}

	actual, err := ioutil.ReadAll(HCL(
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
