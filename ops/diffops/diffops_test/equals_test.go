package diffops_test

import (
	"testing"

	"github.com/krelinga/go-lib/ops/diffops"
	"github.com/stretchr/testify/assert"
)

func TestEqualsPlan(t *testing.T) {
	tests := []struct {
		name string
		lhs, rhs int
		want bool
	}{
		{
			name: "compatible types, different values",
			lhs:  1,
			rhs:  2,
		},
		{
			name: "compatible types, same values",
			lhs:  1,
			rhs:  1,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffops.EqualsPlan(tt.lhs, tt.rhs, diffops.ByEqual[int]())
			assert.Equal(t, tt.want, got, "EqualsPlan(%v, %v) = %v, want %v", tt.lhs, tt.rhs, got, tt.want)
		})
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		name string
		lhs diffops.DiffOper
		rhs any
		want bool
	} {
		{
			name: "compatible types, different values",
			lhs:  &Foo{IsInt: 1},
			rhs:  &Foo{IsInt: 2},
		},
		{
			name: "compatible types, same values",
			lhs:  &Foo{IsInt: 1},
			rhs:  &Foo{IsInt: 1},
			want: true,
		},
		{
			name: "incompatible types",
			lhs:  &Foo{IsInt: 1},
			rhs:  &Bar{IsDouble: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffops.Equals(tt.lhs, tt.rhs)
			assert.Equal(t, tt.want, got, "Equals(%v, %v) = %v, want %v", tt.lhs, tt.rhs, got, tt.want)
		})
	}
}
