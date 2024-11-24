package gutil

func ToPtr[T any](t T) *T {
	return &t
}
