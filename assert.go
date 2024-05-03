package xerr

import (
	"fmt"
)

// assert the cond is true, or panic out an error
func Assert(cond bool, err error) {
	if !cond {
		panic(err)
	}
}

// assert the cond is true, or panic out an error
// you can use it like fmt.Errorf
func Assertf(cond bool, format string, args ...any) {
	Assert(cond, fmt.Errorf(format, args...))
}
