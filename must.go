package xerr

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

func Must0(err error) {
	if err != nil {
		panic(err)
	}
}

func Must2[T0, T1 any](v0 T0, v1 T1, err error) (T0, T1) {
	if err != nil {
		panic(err)
	}
	return v0, v1
}
