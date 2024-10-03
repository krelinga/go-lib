package pipe

import "context"

// The GroupBy function reads from the input channel and groups the values by key.
// This requires storing all of the values in memory, so it is not suitable for large datasets.
func GroupBy[K comparable, V any](ctx context.Context, in <-chan KV[K, V]) <-chan KV[K, []V] {
	out := make(chan KV[K, []V])

	go func() {
		defer close(out)

		seen := make(map[K][]V)

		for record := range in {
			if ctx.Err() != nil {
				return
			}
			seen[record.Key] = append(seen[record.Key], record.Val)
		}

		for key, values := range seen {
			if !TryWrite(ctx, out, KV[K, []V]{Key: key, Val: values}) {
				return
			}
		}
	}()

	return out
}
