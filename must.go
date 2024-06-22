package xerr

import (
	"errors"
	"fmt"
	"reflect"
)

// Because func(...) (T, error) is the most one
// error return type in go project, So we named
// it `Must` as the most easily to use.
//
// if you have a function:
// func Fn(...) (T, error)
// you can just use this function like:
// result := xerr.Must(Fn(...)) to ignore the error handling.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// A function to handle func Fn(...) error.
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

// A function to handle func Fn(...) (T0, T1, error).
func Must2[T0, T1 any](v0 T0, v1 T1, err error) (T0, T1) {
	if err != nil {
		panic(err)
	}
	return v0, v1
}

// MustOk
// If the ok is true, return v.(T)
// otherwise panic.
// you can use this function in get a map value.
func MustOk[T any](v any, ok bool) T {
	if ok {
		res, castOk := v.(T)
		if !castOk {
			panic(fmt.Errorf("can't cast v[%s]", reflect.TypeOf(v).String()))
		}
		return res
	}
	panic(errors.New("not ok"))
}
