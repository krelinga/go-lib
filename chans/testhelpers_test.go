package chans

// spell-checker:ignore chans stretchr

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func assertEventuallyClosed[payload any](t *testing.T, c <-chan payload) {
	assert.Eventually(t, func() bool {
		select {
		case _, ok := <-c:
			return !ok
		default:
			return false
		}
	}, time.Second, 10*time.Millisecond)
}
