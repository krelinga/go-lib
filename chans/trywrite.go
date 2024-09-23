package chans

import "context"

// spell-checker:ignore chans

// TryWrite() writes a value to a channel if the context is not done.
//
// Returns true if the value was written, false if the context was done.
func TryWrite[chanType any](ctx context.Context, out chan<- chanType, val chanType) bool {
	select {
	case <-ctx.Done():
		return false
	case out <- val:
		return true
	}
}
