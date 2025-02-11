package pod_test

import (
	"errors"
	"testing"

	"github.com/krelinga/go-lib/podreflect"
	"github.com/krelinga/go-lib/podreflect/podmock"
	"github.com/stretchr/testify/assert"
)

var (
	errA = errors.New("a must be greater than 0")
	errB = errors.New("b must be nonempty")
)

type testStruct struct {
	podreflect.Struct

	A      int
	Nested *testStruct2
}

func (ts testStruct) PODLocalValidate(reporter podreflect.ErrorReporter) {
	if ts.A <= 0 {
		reporter.Err(errA)
	}
}

type testStruct2 struct {
	podreflect.Struct

	B string
}

func (ts *testStruct2) PODLocalValidate(reporter podreflect.ErrorReporter) {
	if ts.B == "" {
		reporter.Err(errB)
	}
}

func TestCollectErrors(t *testing.T) {
	reporter := podmock.ErrorReporter{}
	podreflect.CollectErrors(testStruct{
		A: -1,
		Nested: &testStruct2{
			B: "",
		},
	}, &reporter)
	assert.ElementsMatch(t, reporter.Errs, []error{errA, errB})
}
