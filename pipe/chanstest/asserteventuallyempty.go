package chanstest

// spell-checker:ignore chanstest stretchr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// AssertEventuallyEmpty asserts that the channel is eventually closed without any elements being written to it
func AssertEventuallyEmpty[payload any](t *testing.T, c <-chan payload) {
	assert.Eventually(t, func() bool {
		select {
		case _, ok := <-c:
			return !ok
		default:
			return false
		}
	}, time.Second, 10*time.Millisecond)
}
