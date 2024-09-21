package routines

import "sync"

// RunAndWait runs the given functions in parallel and waits for all of them to finish.
func RunAndWait(funcs ...func()) {
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
