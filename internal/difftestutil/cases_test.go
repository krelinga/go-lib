package difftestutil

import (
	"reflect"
	"testing"
)


func isComparable[T any]() bool {
	return reflect.TypeFor[T]().Comparable()
}

func TestClaims(t *testing.T) {
	if !isComparable[compStruct]() {
		t.Fatal("compStruct is not comparable")
	}
	if isComparable[nonCompStruct]() {
		t.Fatal("nonCompStruct is comparable")
	}
}