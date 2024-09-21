package chans

// spell-checker:ignore chans stretchr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	assertEventuallyClosed := func(t *testing.T, c <-chan int) {
		assert.Eventually(t, func() bool {
			select {
			case _, ok := <-c:
				return !ok
			default:
				return false
			}
		}, time.Second, 10*time.Millisecond)
	}

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		c := Merge[int]()
		assertEventuallyClosed(t, c)
	})
	t.Run("Single", func(t *testing.T) {
		t.Parallel()
		c1 := make(chan int, 1)
		c := Merge(c1)
		c1 <- 1024
		close(c1)
		assert.Equal(t, 1024, <-c)
		assertEventuallyClosed(t, c)
	})
	t.Run("Multiple", func(t *testing.T) {
		t.Parallel()
		c1 := make(chan int, 1)
		c2 := make(chan int, 1)
		c := Merge(c1, c2)
		c1 <- 1024
		c2 <- 2048
		close(c1)
		close(c2)
		found := []int{}
		for v := range c {
			found = append(found, v)
		}
		assert.ElementsMatch(t, []int{1024, 2048}, found)
	})
}
