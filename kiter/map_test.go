package kiter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	in := FromSlice([]int{1, 2, 3, 4})
	fn := func(x int) int {
		return x * 2
	}

	out := Map(in, fn)
	var result []int
	for v := range out {
		result = append(result, v)
	}

	expected := []int{2, 4, 6, 8}
	assert.Equal(t, expected, result)
}

func TestMapEmpty(t *testing.T) {
	in := FromSlice([]int{})
	fn := func(x int) int {
		return x * 2
	}

	out := Map(in, fn)
	result := []int{}
	for v := range out {
		result = append(result, v)
	}

	expected := []int{}
	assert.Equal(t, expected, result)
}

func TestMapStopEarly(t *testing.T) {
	in := FromSlice([]int{1, 2, 3, 4})
	fn := func(x int) int {
		return x * 2
	}

	out := Map(in, fn)
	var result []int
	for v := range out {
		result = append(result, v)
		if v == 4 {
			break
		}
	}

	expected := []int{2, 4}
	assert.Equal(t, expected, result)
}
