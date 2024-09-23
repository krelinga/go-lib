package pipe

import "sync"

// Wait runs the given functions in new goroutines and waits for all of them to finish.
func Wait(funcs ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(funcs))
	for _, f := range funcs {
		go func() {
			defer wg.Done()
			f()
		}()
	}

	wg.Wait()
}
