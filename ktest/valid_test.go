package ktest

import (
	"fmt"
	"testing"
)

type alwaysValid struct {
	foo string
}

func (alwaysValid) Validate() error {
	return nil
}

type alwaysInvalid struct{}

func (alwaysInvalid) Validate() error {
	return fmt.Errorf("always invalid")
}

func TestValid(t *testing.T) {
	t.Run("AssertValid() should not fail for a valid input", func(t *testing.T) {
		av := AssertValid(t, alwaysValid{foo: "bar"})
		if av.foo != "bar" {
			t.Errorf("Expected foo to be bar, got %s", av.foo)
		}
	})

	t.Run("AssertValid() should fail for an invalid input", func(t *testing.T) {
		t.Skip("Skipping this test because it should fail")
		AssertValid(t, alwaysInvalid{})
		t.Error("Execution should never get here...")
	})
}
