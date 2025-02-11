package pod_test

import (
	"testing"

	"github.com/krelinga/go-lib/pod"
	"github.com/stretchr/testify/assert"
)

func TestSame(t *testing.T) {
	tests := []struct {
		name string
		run  func() bool
		want bool
	}{
		{
			name: "Same map[int]int",
			run: func() bool {
				a := map[int]int{1: 2, 3: 4}
				b := map[int]int{1: 2, 3: 4}
				return pod.Same(pod.WrapMapComp(a), pod.WrapMapComp(b))
			},
			want: true,
		},
		{
			name: "Different value map[int]int",
			run: func() bool {
				a := map[int]int{1: 2, 3: 4}
				b := map[int]int{1: 2, 3: 5}
				return pod.Same(pod.WrapMapComp(a), pod.WrapMapComp(b))
			},
			want: false,
		},
		{
			name: "Different key map[int]int",
			run: func() bool {
				a := map[int]int{1: 2, 3: 4}
				b := map[int]int{1: 2, 4: 4}
				return pod.Same(pod.WrapMapComp(a), pod.WrapMapComp(b))
			},
			want: false,
		},
		{
			name: "Extra key map[int]int",
			run: func() bool {
				a := map[int]int{1: 2}
				b := map[int]int{1: 2, 3: 4}
				return pod.Same(pod.WrapMapComp(a), pod.WrapMapComp(b))
			},
			want: false,
		},
		{
			name: "Missing key map[int]int",
			run: func() bool {
				a := map[int]int{1: 2, 3: 4}
				b := map[int]int{1: 2}
				return pod.Same(pod.WrapMapComp(a), pod.WrapMapComp(b))
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.run())
		})
	}
}
