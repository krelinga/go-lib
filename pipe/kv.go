package pipe

// KV is a simple key-value pair.
type KV[K comparable, V any] struct {
	Key K
	Val V
}