package kslice

func Flatten[T any](in... []T) []T {
	var length int
	for _, s := range in {
		length += len(s)
	}
	out := make([]T, 0, length)
	for _, s := range in {
		out = append(out, s...)
	}
	return out
}