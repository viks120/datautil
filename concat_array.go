package datautil

func ConcatArray[T any](a, b []T) []T {
	result := make([]T, len(a)+len(b))
	copy(result, a)
	copy(result[len(a):], b)
	return result
}
