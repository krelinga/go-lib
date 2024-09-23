package chans

// spell-checker:ignore chans stretchr chanstest

import (
	"context"
	"testing"

	"github.com/krelinga/go-lib/chans/chanstest"
	"github.com/krelinga/go-lib/routines"
	"github.com/stretchr/testify/assert"
)

func TestParallelErr(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		in := make(chan int)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		out, err := ParallelErr(ctx, 10, in, func(int) (int, error) {
			return 0, nil
		})
		close(in)
		chanstest.AssertEventuallyEmpty(t, out)
		chanstest.AssertEventuallyEmpty(t, err)
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
		out, err := ParallelErr(ctx, 10, in, func(x int) (int, error) {
			return x * 2, nil
		})
		chanstest.AssertElementsEventuallyMatch(t, out, []int{0, 2, 4})
		chanstest.AssertEventuallyEmpty(t, err)
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
		out, err := ParallelErr(ctx, 10, in, func(x int) (int, error) {
			if x == 1 {
				return 0, assert.AnError
			}
			return x * 2, nil
		})
		routines.RunAndWait(
			func() {
				chanstest.AssertElementsEventuallyMatch(t, out, []int{0, 4})
			},
			func() {
				assert.ErrorIs(t, <-err, assert.AnError)
				chanstest.AssertEventuallyEmpty(t, err)
			},
		)
	})
}

func TestParallel(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		in := make(chan int)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		out := Parallel(ctx, 10, in, func(int) int {
			return 0
		})
		close(in)
		chanstest.AssertEventuallyEmpty(t, out)
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
		out := Parallel(ctx, 10, in, func(x int) int {
			return x * 2
		})
		found := []int{}
		for x := range out {
			found = append(found, x)
		}
		assert.ElementsMatch(t, []int{0, 2, 4}, found)
	})
}
