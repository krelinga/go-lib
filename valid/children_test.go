package valid_test

import (
	"errors"
	"testing"

	"github.com/krelinga/go-lib/valid"
)

var (
	errVInt            = errors.New("vInt is invalid")
	errVFloat64        = errors.New("vFloat64 is invalid")
	errHidden          = errors.New("hidden is invalid")
	errVStruct         = errors.New("vStruct is invalid")
	errVSlice          = errors.New("vSlice is invalid")
	errVMap            = errors.New("vMap is invalid")
	errVStructNested   = errors.New("vStructNested is invalid")
	errExportedVStruct = errors.New("ExportedVStruct is invalid")
	errDoubleNested    = errors.New("vDoubleNested is invalid")
)

type vInt int

func (v vInt) Validate() error {
	if int(v) == 0 {
		return errVInt
	}
	return nil
}

type vFloat64 float64

func (v vFloat64) Validate() error {
	if float64(v) == 0 {
		return errVFloat64
	}
	return nil
}

type vHidden struct{}

func (v vHidden) Validate() error {
	return errHidden
}

type vSlice []vInt

func (v vSlice) Validate() error {
	return errors.Join(errVSlice, valid.Children(v))
}

type vMap map[vInt]vFloat64

func (v vMap) Validate() error {
	return errors.Join(errVMap, valid.Children(v))
}

type noValidate struct {
	VInt     vInt
	Int      int
	VFloat64 vFloat64
	Float64  float64
}

type vStruct struct {
	VInt     vInt
	VFloat64 vFloat64
	VSlice   vSlice
	VMap     vMap
	Nested   noValidate
	//lint:ignore U1000 This field should not be visible to valid.Children().
	hidden   vHidden
}

func (s vStruct) Validate() error {
	return errors.Join(errVStruct, valid.Children(s))
}

type vStructNested struct {
	noValidate
}

func (s vStructNested) Validate() error {
	return errors.Join(errVStructNested, valid.Children(s))
}

type noValidateNested struct {
	vStruct
}

type noValidateNestedPtr struct {
	*vStruct
}

type ExportedVStruct struct {
	VInt     vInt
	VFloat64 vFloat64
	VSlice   vSlice
	VMap     vMap
}

func (s ExportedVStruct) Validate() error {
	return errors.Join(errExportedVStruct, valid.Children(s))
}

type vDoubleNested struct {
	ExportedVStruct
}

func (s vDoubleNested) Validate() error {
	return errors.Join(errDoubleNested, valid.Children(s))
}

func TestChildren(t *testing.T) {
	tests := []struct {
		name       string
		fn         func() any
		in         any
		wantErrs   []error
		unwantErrs []error
	}{
		{
			name:       "vInt Valid Produces No Error",
			in:         vInt(1),
			wantErrs:   nil,
			unwantErrs: []error{errVInt},
		},
		{
			name:       "vInt Invalid Produces No Error",
			in:         vInt(0),
			wantErrs:   nil,
			unwantErrs: []error{errVInt},
		},
		{
			name: "*vInt Valid Produces No Error",
			fn: func() any {
				v := vInt(1)
				return &v
			},
			wantErrs:   nil,
			unwantErrs: []error{errVInt},
		},
		{
			name: "*vInt Invalid Produces Error",
			fn: func() any {
				v := vInt(0)
				return &v
			},
			wantErrs: []error{errVInt},
		},
		{
			name:     "int Produces No Error",
			in:       int(1),
			wantErrs: nil,
		},
		{
			name: "*int Produces No Error",
			fn: func() any {
				v := int(1)
				return &v
			},
			wantErrs: nil,
		},
		{
			name:       "vFloat64 Valid Produces No Error",
			in:         vFloat64(1),
			wantErrs:   nil,
			unwantErrs: []error{errVFloat64},
		},
		{
			name:       "vFloat64 Invalid Produces No Error",
			in:         vFloat64(0),
			wantErrs:   nil,
			unwantErrs: []error{errVFloat64},
		},
		{
			name: "*vFloat64 Valid Produces No Error",
			fn: func() any {
				v := vFloat64(1)
				return &v
			},
			wantErrs:   nil,
			unwantErrs: []error{errVFloat64},
		},
		{
			name: "*vFloat64 Invalid Produces Error",
			fn: func() any {
				v := vFloat64(0)
				return &v
			},
			wantErrs: []error{errVFloat64},
		},
		{
			name:     "float64 Produces No Error",
			in:       float64(1),
			wantErrs: nil,
		},
		{
			name: "*float64 Produces No Error",
			fn: func() any {
				v := float64(1)
				return &v
			},
			wantErrs: nil,
		},
		{
			name:       "vSlice Valid Children Produces No Error",
			in:         vSlice{vInt(1), vInt(1)},
			wantErrs:   nil,
			unwantErrs: []error{errVSlice},
		},
		{
			name:       "vSlice Invalid Children Produces Error",
			in:         vSlice{vInt(0), vInt(1)},
			wantErrs:   []error{errVInt},
			unwantErrs: []error{errVSlice},
		},
		{
			name:     "*vSlice Valid Children Produces Error",
			in:       &vSlice{vInt(1), vInt(1)},
			wantErrs: []error{errVSlice},
		},
		{
			name:       "vMap Valid Children Produces No Error",
			in:         vMap{vInt(1): vFloat64(1), vInt(2): vFloat64(2)},
			wantErrs:   nil,
			unwantErrs: []error{errVMap},
		},
		{
			name:       "vMap Invalid Key Produces Error",
			in:         vMap{vInt(0): vFloat64(1), vInt(2): vFloat64(2)},
			wantErrs:   []error{errVInt},
			unwantErrs: []error{errVMap},
		},
		{
			name:       "vMap Invalid Value Produces Error",
			in:         vMap{vInt(1): vFloat64(0), vInt(2): vFloat64(2)},
			wantErrs:   []error{errVFloat64},
			unwantErrs: []error{errVMap},
		},
		{
			name:     "*vMap Valid Children Produces Error",
			in:       &vMap{vInt(1): vFloat64(1), vInt(2): vFloat64(2)},
			wantErrs: []error{errVMap},
		},
		{
			name:       "vStruct Default Has All Child Errors",
			in:         vStruct{},
			wantErrs:   []error{errVInt, errVFloat64, errVSlice, errVMap},
			unwantErrs: []error{errVStruct},
		},
		{
			name:     "*vStruct Default Has All Child Errors And Its Own Error",
			in:       &vStruct{},
			wantErrs: []error{errVInt, errVFloat64, errVSlice, errVMap, errVStruct},
		},
		{
			name: "vStruct Valid vInt",
			in: vStruct{
				VInt: vInt(1),
				Nested: noValidate{
					VInt: vInt(2),
				},
			},
			wantErrs:   []error{errVFloat64, errVSlice, errVMap},
			unwantErrs: []error{errVInt},
		},
		{
			name: "vStruct Valid vFloat64",
			in: vStruct{
				VFloat64: vFloat64(1),
				Nested: noValidate{
					VFloat64: vFloat64(2),
				},
			},
			wantErrs:   []error{errVInt, errVSlice, errVMap},
			unwantErrs: []error{errVFloat64},
		},
		{
			name:       "vStructNested Default Has All Child Errors",
			in:         vStructNested{},
			wantErrs:   []error{errVInt, errVFloat64},
			unwantErrs: []error{errVStructNested},
		},
		{
			name:     "*vStructNested Default Has All Child Errors And Its Own Error",
			in:       &vStructNested{},
			wantErrs: []error{errVInt, errVFloat64, errVStructNested},
		},
		{
			name: "noValidateNested Default Has All Child Errors",
			in: noValidateNested{
				vStruct: vStruct{},
			},
			wantErrs:   []error{errVInt, errVFloat64, errVSlice, errVMap},
			unwantErrs: []error{errVStruct},
		},
		{
			name: "noValidateNestedPtr Default Has All Child Errors",
			in: noValidateNestedPtr{
				vStruct: &vStruct{},
			},
			wantErrs:   []error{errVInt, errVFloat64, errVSlice, errVMap},
			unwantErrs: []error{errVStruct},
		},
		{
			name: "vDoubleNested Default Has All Child Errors",
			in: vDoubleNested{
				ExportedVStruct: ExportedVStruct{},
			},
			wantErrs:   []error{errVInt, errVFloat64, errVSlice, errVMap, errExportedVStruct},
			unwantErrs: []error{errDoubleNested},
		},
		{
			name: "nil",
			in:   nil,
			wantErrs: nil,
		},
		{
			name: "interface created from nil pointer to type",
			fn: func() any {
				var v *vInt
				return v
			},
			wantErrs: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var in any
			if tt.fn != nil {
				in = tt.fn()
			} else {
				in = tt.in
			}
			err := valid.Children(in)
			if err == nil {
				for _, wantErr := range tt.wantErrs {
					if wantErr != nil {
						t.Errorf("got nil, want %v", wantErr)
					}
				}
			} else {
				if len(tt.wantErrs) == 0 {
					t.Errorf("got %v, want nil", err)
				}
				for _, wantErr := range tt.wantErrs {
					if !errors.Is(err, wantErr) {
						t.Errorf("got %v, want %v", err, wantErr)
					}
				}
			}
			tt.unwantErrs = append(tt.unwantErrs, errHidden)
			for _, unwantErr := range tt.unwantErrs {
				if errors.Is(err, unwantErr) {
					t.Errorf("got %v, unwanted %v", err, unwantErr)
				}
			}
		})
	}
}
