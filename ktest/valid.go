package ktest

import (
	"testing"

	"github.com/krelinga/go-lib/valid"
)

// AssertValid() checks if the given input is valid.
//
// If it is not valid then AssertValid() calls t.Fatalf() with the error message, which
// should cause the test to stop immediately.
// AssertValid() returns the input, so it can be used in the same line as the input is created.
func AssertValid[T valid.Validator](t *testing.T, v T) T{
	if err := v.Validate(); err != nil {
		t.Fatalf("Validation failed: %v", err)
	}
	return v
}