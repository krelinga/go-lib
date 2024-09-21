package routines

// spell-checker:ignore chans stretchr chanstest

import (
	"testing"

	"github.com/krelinga/go-lib/chans/chanstest"
	"github.com/stretchr/testify/assert"
)

func TestRunAndWait(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		RunAndWait()
	})
	t.Run("Multiple", func(t *testing.T) {
		t.Parallel()
		chanIsClosed := func(c <-chan struct{}) bool {
			select {
			case _, ok := <-c:
				return !ok
			default:
				return false
			}
		}
		aWait := make(chan struct{})
		bWait := make(chan struct{})
		cWait := make(chan struct{})
		fnReturned := func() <-chan struct{} {
			c := make(chan struct{})
			go func() {
				RunAndWait(func() { <-aWait }, func() { <-bWait }, func() { <-cWait })
				close(c)
			}()
			return c
		}()
		assert.False(t, chanIsClosed(fnReturned))
		close(aWait)
		assert.False(t, chanIsClosed(fnReturned))
		close(bWait)
		assert.False(t, chanIsClosed(fnReturned))
		close(cWait)
		chanstest.AssertEventuallyEmpty(t, fnReturned)
	})
}
