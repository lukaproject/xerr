package xerr

func Recover[T any](v *T) {
	if e := recover(); e != nil {
		*v = e.(T)
	}
}
