package pipe

// The ToArray function returns a function which, when called, reads from the input channel
// and appends the values to the output slice.  The returned function terminates
// when the input channel is closed.
func ToArrayFunc[V any](in <-chan V, out *[]V) func() {
	return func() {
		for v := range in {
			*out = append(*out, v)
		}
	}
}

func dropAll[V any](in <-chan V) {
	for range in {
	}
}

// The First function returns a function which, when called, reads the first value
// from the input channel, and passes that value to the given function.  If the channel
// never receives a value, then the given function is never called.  The returned function
// consumes all values from the input channel, and only terminates when the input channel
// is closed.
func FirstFunc[V any](in <-chan V, fn func(V)) func() {
	return func() {
		first, ok := <-in
		if ok {
			fn(first)
			dropAll(in)
		}
	}
}

// The Last function returns a function which, when called, reads all values from the input
// channel, and passes the last value to the given function.  If the channel never receives
// a value, then the given function is never called.  The returned function only terminates
// when the input channel is closed.
func LastFunc[V any](in <-chan V, fn func(V)) func() {
	return func() {
		var last V
		var ok bool
		for v := range in {
			last = v
			ok = true
		}
		if ok {
			fn(last)
		}
	}
}

// The Discard function returns a function which, when called, consumes all values from
// the input channel.  The returned function only terminates when the input channel is closed.
func DiscardFunc[V any](in <-chan V) func() {
	return func() {
		dropAll(in)
	}
}

// The ToMap function returns a function which, when called, consumes all key-value pairs
// from the input channel and stores them in the given map.  The returned function only
// terminates when the input channel is closed.  If the same key is encountered multiple
// times, the value in the map will be the last value seen.
func ToMapFunc[K comparable, V any](in <-chan KV[K, V], out *map[K]V) func() {
	return func() {
		for kv := range in {
			(*out)[kv.Key] = kv.Val
		}
	}
}
