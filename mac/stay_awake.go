package mac

import (
	"os/exec"
)

type StayAwakeOpts struct {
	// Prevents the display from sleeping if true.
	Display bool

	// Prevents the system from sleeping if true.
	System bool

	// Prevents the disk from sleeping if true.
	Disk bool
}

// StayAwake prevents the system from sleeping.
//
// Returns two values: a function and an error.
// If the error is not nil, then the function must be called to resume normal sleep behavior.
// If the error is nil then it isn't necessary to call the function.
func StayAwake(opts StayAwakeOpts) (func() error, error) {
	args := []string{}
	doNothing := func() error { return nil }
	if opts.Display {
		args = append(args, "-d")
	}
	if opts.System {
		args = append(args, "-i")
	}
	if opts.Disk {
		args = append(args, "-m")
	}
	cmd := exec.Command("caffeinate", args...)

	if err := cmd.Start(); err != nil {
		return doNothing, err
	}
	stopChan := make(chan struct{})
	errChan := make(chan error)
	go func() {
		<-stopChan
		if err := cmd.Process.Kill(); err != nil {
			errChan <- err
			return
		}
		// No need to check the error returned by this command, we expect it to fail because we killed the process.
		cmd.Wait()
		close(errChan)
	}()

	stopFn := func() error {
		close(stopChan)
		return <-errChan
	}

	return stopFn, nil
}
