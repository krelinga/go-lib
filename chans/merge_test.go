package chans

// spell-checker:ignore chans stretchr chanstest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/krelinga/go-lib/chans/chanstest"
)

func TestMerge(t *testing.T) {
	t.Parallel()

	t.Run("Empty", func(t *testing.T) {
		t.Parallel()
		c := Merge[int]()
		chanstest.AssertEventuallyClosed(t, c)
	})
	t.Run("Single", func(t *testing.T) {
		t.Parallel()
		c1 := make(chan int, 1)
		c := Merge(c1)
		c1 <- 1024
		close(c1)
		assert.Equal(t, 1024, <-c)
		chanstest.AssertEventuallyClosed(t, c)
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
