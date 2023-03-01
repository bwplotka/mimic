package encoding

import "io"

// Encoder implements the needed functions to encode a Go struct to a particular config language.
type Encoder interface {
	io.Reader
	Commenter(b []byte, comments []string) []byte
}
