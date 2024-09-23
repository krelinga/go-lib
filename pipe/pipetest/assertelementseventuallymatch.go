package pipetest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func AssertElementsEventuallyMatch[payload any](t *testing.T, c <-chan payload, expected []payload) {
	found := []payload{}
	consumeAllAvailable := func() bool {
		for {
			select {
			case v, ok := <-c:
				if !ok {
					return true
				}
				found = append(found, v)
			default:
				return false
			}
		}
	}
	assert.Eventually(t, consumeAllAvailable, time.Second, 10*time.Millisecond)
}
