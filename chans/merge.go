package chans

// spell-checker:ignore chans

import "sync"

// Merge() merges multiple channels into a single channel.
func Merge[t any](channels ...<-chan t) <-chan t {
	out := make(chan t, len(channels))
	var wg sync.WaitGroup
	wg.Add(len(channels))
	for _, c := range channels {
		c := c
		go func() {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
