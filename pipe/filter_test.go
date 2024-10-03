package pipe

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParDoFilterErr(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		filterFunc  func(int) (bool, error)
		expectedOut []int
		expectedErr []error
	}{
		{
			name:        "No errors, all pass",
			input:       []int{1, 2, 3, 4, 5},
			filterFunc:  func(v int) (bool, error) { return v%2 == 0, nil },
			expectedOut: []int{2, 4},
			expectedErr: nil,
		},
		{
			name:        "No errors, none pass",
			input:       []int{1, 3, 5},
			filterFunc:  func(v int) (bool, error) { return v%2 == 0, nil },
			expectedOut: []int{},
			expectedErr: nil,
		},
		{
			name:  "Some errors",
			input: []int{1, 2, 3, 4, 5},
			filterFunc: func(v int) (bool, error) {
				if v == 3 {
					return false, errors.New("error on 3")
				}
				return v%2 == 0, nil
			},
			expectedOut: []int{2, 4},
			expectedErr: []error{errors.New("error on 3")},
		},
		{
			name:  "All errors",
			input: []int{1, 2, 3, 4, 5},
			filterFunc: func(v int) (bool, error) {
				return false, errors.New("error on all")
			},
			expectedOut: []int{},
			expectedErr: []error{
				errors.New("error on all"),
				errors.New("error on all"),
				errors.New("error on all"),
				errors.New("error on all"),
				errors.New("error on all"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan int, len(tt.input))
			for _, v := range tt.input {
				in <- v
			}
			close(in)

			out, errs := ParDoFilterErr(context.Background(), 1, in, tt.filterFunc)

			var actualOut []int
			var actualErr []error
			Wait(
				ToArrayFunc(out, &actualOut),
				ToArrayFunc(errs, &actualErr),
			)

			assert.ElementsMatch(t, tt.expectedOut, actualOut)
			assert.ElementsMatch(t, tt.expectedErr, actualErr)
		})
	}
}
func TestParDoFilter(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		filterFunc  func(int) bool
		expectedOut []int
	}{
		{
			name:        "All pass",
			input:       []int{1, 2, 3, 4, 5},
			filterFunc:  func(v int) bool { return v%2 == 0 || v%2 != 0 },
			expectedOut: []int{1, 2, 3, 4, 5},
		},
		{
			name:        "None pass",
			input:       []int{1, 3, 5},
			filterFunc:  func(v int) bool { return v%2 == 0 },
			expectedOut: []int{},
		},
		{
			name:        "Some pass",
			input:       []int{1, 2, 3, 4, 5},
			filterFunc:  func(v int) bool { return v%2 == 0 },
			expectedOut: []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan int, len(tt.input))
			for _, v := range tt.input {
				in <- v
			}
			close(in)

			out := ParDoFilter(context.Background(), 1, in, tt.filterFunc)

			var actualOut []int
			Wait(ToArrayFunc(out, &actualOut))

			assert.ElementsMatch(t, tt.expectedOut, actualOut)
		})
	}
}
