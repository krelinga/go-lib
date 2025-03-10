package difftestutil

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/krelinga/go-lib/diff"
	"github.com/krelinga/go-lib/kslice"
	"github.com/krelinga/go-lib/ktest"
	"github.com/stretchr/testify/assert"
)

var updateFlag = flag.Bool("update", false, "update golden files")

func (c testCase[T]) Name() string {
	return c.name
}

func (c testCase[T]) RunDiffTest(t *testing.T) {
	t.Run(c.name, func(t *testing.T) {
		got := diff.Diff(c.lhs, c.rhs)
		assert.Equalf(t, c.want, got, "Diff() = %v, want %v", got, c.want)
	})
}

type fakeTestingT []string

func (t *fakeTestingT) Helper() {}

func (t *fakeTestingT) Errorf(format string, args ...interface{}) {
	*t = append(*t, fmt.Sprintf(format, args...))
}

func (c testCase[T]) RunAssertEqualTest(t *testing.T) {
	t.Run(c.name, func(t *testing.T) {
		goldenFilePath := filepath.Join("testdata", t.Name()+".golden")
		rawGot := &fakeTestingT{}
		ktest.AssertEqual(rawGot, c.lhs, c.rhs)
		got := strings.Join(*rawGot, "")
		if *updateFlag {
			if len(got) == 0 {
				// Delete any file that exists.
				err := os.Remove(goldenFilePath)
				if err != nil && !os.IsNotExist(err) {
					t.Fatalf("failed to remove golden file: %v", err)
				}
			} else {
				// Write the output to a golden file.
				file, err := os.Create(goldenFilePath)
				if err != nil {
					t.Fatalf("failed to write golden file: %v", err)
				}
				defer file.Close()
				fmt.Fprint(file, got)
			}
		} else {
			rawWant, err := os.ReadFile(goldenFilePath)
			if os.IsNotExist(err) {
				rawWant = []byte{}
			} else if err != nil {
				t.Fatalf("failed to read golden file: %v", err)
			}
			want := string(rawWant)
			assert.Equal(t, want, got)
		}
	})
}

func nilPtr[T any]() *T {
	return nil
}

var TestCases = kslice.Flatten(floatCases, intCases, interfaceCases, mapCases, pointerCases, sliceCases, stringCases, structCases)

func isComparable[T any]() bool {
	return reflect.TypeFor[T]().Comparable()
}

func init() {
	if !isComparable[compStruct]() {
		panic("compStruct is not comparable")
	}
	if isComparable[nonCompStruct]() {
		panic("nonCompStruct is comparable")
	}

	names := make(map[string]struct{})
	for _, tc := range TestCases {
		if _, ok := names[tc.Name()]; ok {
			panic(fmt.Sprintf("duplicate test case name: %s", tc.Name()))
		}
		names[tc.Name()] = struct{}{}
	}

	diff.Register[getter](diff.WithMethods("Get"))
	diff.Register[ptrGetter](diff.WithMethods("GetPtr"))
	diff.Register[myInt](diff.WithMethods("Get"))
	// TODO: this is failing with empty package path right now ... I might need some special handling for pointers
	// diff.Register[*myInt](diff.WithMethods("GetPtr"))
}
