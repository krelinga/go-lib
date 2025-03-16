package pipe

// spell-checker:ignore chans stretchr pipetest

import (
	"context"
	"testing"

	"github.com/krelinga/go-lib/pipe/pipetest"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		c := Merge[int, chan int](ctx)
		pipetest.AssertEventuallyEmpty(t, c)
	})
	t.Run("Single", func(t *testing.T) {
		t.Parallel()
		c1 := make(chan int, 1)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		c := Merge(ctx, c1)
		c1 <- 1024
		close(c1)
		pipetest.AssertElementsEventuallyMatch(t, c, []int{1024})
	})
	t.Run("Multiple", func(t *testing.T) {
		t.Parallel()
		c1 := make(chan int, 1)
		c2 := make(chan int, 1)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		c := Merge(ctx, c1, c2)
		c1 <- 1024
		c2 <- 2048
		close(c1)
		close(c2)
		pipetest.AssertElementsEventuallyMatch(t, c, []int{1024, 2048})
	})
}
