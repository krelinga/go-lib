package pod_test

import (
	"testing"

	"github.com/krelinga/go-lib/pod"
)

func TestCopy(t *testing.T) {
	ma := map[string]int{"a": 1, "b": 2}
	mb := map[string]int{}

	pod.DeepCopyTo(pod.WrapMapComp(ma), pod.WrapMapComp(mb))
}