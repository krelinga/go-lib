package mac

import "runtime"

func Is() bool {
	return runtime.GOOS == "darwin"
}