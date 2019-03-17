package gocodeit

import "fmt"

// Panic is a simple combination of Sprintf+panic.
func Panicf(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a...))
}
