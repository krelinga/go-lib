package chans

// spell-checker:ignore chans stretchr chanstest

import (
	"sync"
	"testing"

	"github.com/krelinga/go-lib/chans/chanstest"
	"github.com/stretchr/testify/assert"
)

func TestParallelErr(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		in := make(chan int)
		out, err := ParallelErr(10, in, func(int) (int, error) {
			return 0, nil
		})
		close(in)
		chanstest.AssertEventuallyClosed(t, out)
		chanstest.AssertEventuallyClosed(t, err)
	})
	t.Run("EveryInputSeen", func(t *testing.T) {
		t.Parallel()
		in := make(chan int, 3)
		for i := 0; i < 3; i++ {
			in <- i
		}
		close(in)
		out, err := ParallelErr(10, in, func(x int) (int, error) {
			return x * 2, nil
		})
		found := []int{}
		for x := range out {
			found = append(found, x)
		}
		assert.ElementsMatch(t, []int{0, 2, 4}, found)
		chanstest.AssertEventuallyClosed(t, err)
	})
	t.Run("Errors", func(t *testing.T) {
		t.Parallel()
		in := make(chan int, 3)
		for i := 0; i < 3; i++ {
			in <- i
		}
		close(in)
		out, err := ParallelErr(10, in, func(x int) (int, error) {
			if x == 1 {
				return 0, assert.AnError
			}
			return x * 2, nil
		})
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			found := []int{}
			for x := range out {
				found = append(found, x)
			}
			assert.ElementsMatch(t, []int{0, 4}, found)
		}()
		go func() {
			defer wg.Done()
			assert.ErrorIs(t, <-err, assert.AnError)
			chanstest.AssertEventuallyClosed(t, err)
		}()
		wg.Wait()
	})
}

func TestParallel(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		in := make(chan int)
		out := Parallel(10, in, func(int) int {
			return 0
		})
		close(in)
		chanstest.AssertEventuallyClosed(t, out)
	})
	t.Run("EveryInputSeen", func(t *testing.T) {
		t.Parallel()
		in := make(chan int, 3)
		for i := 0; i < 3; i++ {
			in <- i
		}
		close(in)
		out := Parallel(10, in, func(x int) int {
			return x * 2
		})
		found := []int{}
		for x := range out {
			found = append(found, x)
		}
		assert.ElementsMatch(t, []int{0, 2, 4}, found)
	})
}
