package encoding

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestYAML_Encoding2Structs(t *testing.T) {
	type A struct {
		Field1 string
		Field2 int
		Inner  *A
	}

	b, err := ioutil.ReadAll(YAML(
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
	require.NoError(t, err)
	require.Equal(t, `Field1: "1"
Field2: 1
Inner:
  Field1: inner1
  Field2: 11
  Inner: null
---
Field1: "2"
Field2: 2
Inner: null
`, string(b))
}
