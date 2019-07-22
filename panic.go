package mimic

import "fmt"

// Panicf allows panic error propagation using sprintf-like formatting.
func Panicf(format string, a ...interface{}) {
	panic(fmt.Sprintf("mimic: "+format, a...))
}

// PanicErr allows to panic because of certain error.
func PanicErr(err error) {
	Panicf("failed to execute; err:", err)
}
