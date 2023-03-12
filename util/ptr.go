package util

func Ptr[T any](v T) *T {
	a := new(T)
	*a = v

	return a
}
