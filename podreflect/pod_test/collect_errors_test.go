package pod_test

import (
	"errors"
	"testing"

	pod "github.com/krelinga/go-lib/podreflect"
	"github.com/krelinga/go-lib/podreflect/podmock"
	"github.com/stretchr/testify/assert"
)

var (
	errA = errors.New("a must be greater than 0")
	errB = errors.New("b must be nonempty")
)

type testStruct struct {
	pod.Struct

	A      int
	Nested *testStruct2
}

func (ts testStruct) PODLocalValidate(reporter pod.ErrorReporter) {
	if ts.A <= 0 {
		reporter.Err(errA)
	}
}

type testStruct2 struct {
	pod.Struct

	B string
}

func (ts *testStruct2) PODLocalValidate(reporter pod.ErrorReporter) {
	if ts.B == "" {
		reporter.Err(errB)
	}
}

func TestCollectErrors(t *testing.T) {
	reporter := podmock.ErrorReporter{}
	pod.CollectErrors(testStruct{
		A: -1,
		Nested: &testStruct2{
			B: "",
		},
	}, &reporter)
	assert.ElementsMatch(t, reporter.Errs, []error{errA, errB})
}

func TestValidationConfig(t *testing.T) {
	config := pod.StructConfig[testStruct](
		pod.FieldConfig[int]("A",
			pod.WithFieldValidator(func(a int, reporter pod.ErrorReporter) {
				if a <= 0 {
					reporter.Err(errA)
				}
			}),
		),
	)
	assert.Equal(t, 1, len(config.FieldConfigs))
}
