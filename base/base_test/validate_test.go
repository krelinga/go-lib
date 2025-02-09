package base_test

import (
	"errors"
	"testing"

	"github.com/krelinga/go-lib/base"
	"github.com/stretchr/testify/assert"
)

var errNotValid = errors.New("not valid")

type testValidator struct {
	err error
}

func (tv testValidator) Validate() error {
	return tv.err
}

type testValidatorPtr struct {
	err error
}

func (tv *testValidatorPtr) Validate() error {
	return tv.err
}

func TestValidOrPanic(t *testing.T) {
	tests := []struct {
		name    string
		v       base.Validator
		wantErr error
	}{
		{
			name: "Valid By Value",
			v:    testValidator{},
		},
		{
			name:    "Invalid By Value",
			v:       testValidator{err: errNotValid},
			wantErr: errNotValid,
		},
		{
			name: "Valid By Pointer",
			v:    &testValidatorPtr{},
		},
		{
			name:    "Invalid By Pointer",
			v:       &testValidatorPtr{err: errNotValid},
			wantErr: errNotValid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Validate()
			assert.ErrorIs(t, got, tt.wantErr)
			if tt.wantErr != nil {
				assert.PanicsWithError(t, tt.wantErr.Error(), func() {
					base.ValidOrPanic(tt.v)
				})
			} else {
				if !assert.NotPanics(t, func() {
					base.ValidOrPanic(tt.v)
				}) {
					return
				}
				got := base.ValidOrPanic(tt.v)
				assert.Equal(t, tt.v, got)
			}
		})
	}

	// Ensure that the test can be used in a fluid style.
	assert.PanicsWithError(t, errNotValid.Error(), func() {
		base.ValidOrPanic(testValidator{err: errNotValid})
	})
	assert.PanicsWithError(t, errNotValid.Error(), func() {
		base.ValidOrPanic(&testValidatorPtr{err: errNotValid})
	})
	assert.Nil(t, base.ValidOrPanic(testValidator{}).err)
	assert.Nil(t, base.ValidOrPanic(&testValidatorPtr{}).err)
}
