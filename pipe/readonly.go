package pipe

// spell-checker:ignore chans

// ReadOnly returns a read-only channel.
//
// This is useful to work around some limitations in go generics.
func ReadOnly[t any](in chan t) <-chan t {
	return in
}
