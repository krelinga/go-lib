package pipe

// spell-checker:ignore chans stretchr pipetest

import (
	"context"
	"testing"

	"github.com/krelinga/go-lib/pipe/pipetest"
	"github.com/stretchr/testify/assert"
)

func TestParDoErr(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		in := make(chan int)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		out, err := ParDoErr(ctx, 10, in, func(int) (int, error) {
			return 0, nil
		})
		close(in)
		pipetest.AssertEventuallyEmpty(t, out)
		pipetest.AssertEventuallyEmpty(t, err)
	})
	t.Run("EveryInputSeen", func(t *testing.T) {
		t.Parallel()
		in := make(chan int, 3)
		for i := 0; i < 3; i++ {
			in <- i
		}
		close(in)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		out, err := ParDoErr(ctx, 10, in, func(x int) (int, error) {
			return x * 2, nil
		})
		pipetest.AssertElementsEventuallyMatch(t, out, []int{0, 2, 4})
		pipetest.AssertEventuallyEmpty(t, err)
	})
	t.Run("Errors", func(t *testing.T) {
		t.Parallel()
		in := make(chan int, 3)
		for i := 0; i < 3; i++ {
			in <- i
		}
		close(in)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		out, err := ParDoErr(ctx, 10, in, func(x int) (int, error) {
			if x == 1 {
				return 0, assert.AnError
			}
			return x * 2, nil
		})
		Wait(
			func() {
				pipetest.AssertElementsEventuallyMatch(t, out, []int{0, 4})
			},
			func() {
				assert.ErrorIs(t, <-err, assert.AnError)
				pipetest.AssertEventuallyEmpty(t, err)
			},
		)
	})
}

func TestParDo(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		in := make(chan int)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		out := ParDo(ctx, 10, in, func(int) int {
			return 0
		})
		close(in)
		pipetest.AssertEventuallyEmpty(t, out)
	})
	t.Run("EveryInputSeen", func(t *testing.T) {
		t.Parallel()
		in := make(chan int, 3)
		for i := 0; i < 3; i++ {
			in <- i
		}
		close(in)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		out := ParDo(ctx, 10, in, func(x int) int {
			return x * 2
		})
		found := []int{}
		for x := range out {
			found = append(found, x)
		}
		assert.ElementsMatch(t, []int{0, 2, 4}, found)
	})
}
