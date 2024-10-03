package pipe

// spell-checker:ignore stretchr

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTryWrite(t *testing.T) {
	t.Run("ContextDone", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Immediately cancel the context

		out := make(chan int)
		val := 42

		written := TryWrite(ctx, out, val)
		assert.False(t, written, "Expected write to fail, but it succeeded")
	})

	t.Run("Success", func(t *testing.T) {
		ctx := context.Background()

		out := make(chan int, 1)
		val := 42

		written := TryWrite(ctx, out, val)
		assert.True(t, written, "Expected write to succeed, but it failed")

		select {
		case v := <-out:
			assert.Equal(t, val, v, "Expected %d, got %d", val, v)
		default:
			assert.Fail(t, "Expected value in channel, but it was empty")
		}
	})

	t.Run("ContextTimeout", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		out := make(chan int)
		val := 42

		written := TryWrite(ctx, out, val)
		assert.False(t, written, "Expected write to fail, but it succeeded")
	})
}
