package ktest_test

import (
	"testing"

	"github.com/krelinga/go-lib/ktest"
)

type testStruct struct {
	A int
	B string
	C map[string]int
	D []int
}

func TestAssertEqual(t *testing.T) {
	t.Skip("This test will always fail, it is used to demonstrate the AssertEqual function.")
	ktest.AssertEqual(t, int(1), int(2))
	t1 := testStruct{A: 1, B: "2", C: map[string]int{"3": 4, "7": 8}, D: []int{5, 6}}
	t2 := testStruct{A: 1, B: "2", C: map[string]int{"3": 4, "9": 10}, D: []int{5}}
	ktest.AssertEqual(t, t1, t2)
}
