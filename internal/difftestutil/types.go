package difftestutil

import (
	"testing"

	"github.com/krelinga/go-lib/diff"
)

type TestCase interface {
	Name() string
	RunDiffTest(t *testing.T)
	RunAssertEqualTest(t *testing.T)
}

type testCase[T any] struct {
	name string
	lhs  T
	rhs  T
	want []diff.Result
}

func ptr[T any](v T) *T {
	return &v
}

type getter interface {
	Get() int
}

type ptrGetter interface {
	GetPtr() int
}

type myInt int

func (m myInt) Get() int {
	return int(m)
}

func (m *myInt) GetPtr() int {
	return int(*m)
}

type myString string

func (m myString) Get() int {
	return len(m)
}

type compStruct struct {
	Str string
	Int int
}

type nonCompStruct struct {
	Slice []int
	pInt  int
}

type ChildStruct struct {
	Str string
}

type ParentStruct struct {
	ChildStruct
	Int int
}

type mySlice []int

type myMap map[int]string
