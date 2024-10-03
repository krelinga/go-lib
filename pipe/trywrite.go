package pipe

import "context"

// spell-checker:ignore chans

// TryWrite() blocks until either the value is written to the channel or the context becomes done.
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
