package pipe

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	t.Parallel()

	t.Run("RunsToCompletion", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		in := make(chan *KV[string, int])

		go func() {
			defer close(in)
			in <- &KV[string, int]{Key: "a", Val: 1}
			in <- &KV[string, int]{Key: "b", Val: 2}
			in <- &KV[string, int]{Key: "a", Val: 3}
			in <- &KV[string, int]{Key: "b", Val: 4}
		}()

		out := GroupBy(ctx, in)
		result := make(map[string][]int)

		Wait(
			ToMapFunc(out, &result),
		)

		expected := map[string][]int{
			"a": {1, 3},
			"b": {2, 4},
		}

		assert.Equal(t, expected, result)
	})

	t.Run("ContextCancelled", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel right away.

		in := make(chan *KV[string, int])
		go func() {
			defer close(in)
			in <- &KV[string, int]{Key: "a", Val: 1}
		}()

		out := GroupBy(ctx, in)
		result := make(map[string][]int)

		Wait(
			ToMapFunc(out, &result),
		)

		assert.Equal(t, 0, len(result))
	})
}
