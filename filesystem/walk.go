package filesystem

import (
	"context"
	"io/fs"
	"path/filepath"

	"github.com/krelinga/go-lib/pipe"
)

// The type used to represent individual files or directories on the filesystem.
// WalkAll() returns a channel of DirEntry objects.
type DirEntry interface {
	fs.DirEntry
	Path() string
}

type dirEntry struct {
	fs.DirEntry
	path string
}

func (de *dirEntry) Path() string {
	return de.path
}

// WalkAll() spawns a new goroutine to recursively walk the filesystem starting at the given root directory.
//
// All files that are discovered in this way are sent to the DirEntry output channel.
// Any errors that are encountered are sent to the error output channel.
// If any error is encountered, the walk is continued (callers can choose to stop the walk by cancelling the context).
// WalkAll() does not follow symbolic links.
func WalkAll(ctx context.Context, root string) (<-chan DirEntry, <-chan error) {
	out := make(chan DirEntry)
	errs := make(chan error)

	go func() {
		defer close(out)
		defer close(errs)

		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if ctx.Err() != nil {
				return filepath.SkipAll
			}
			if err != nil {
				if !pipe.TryWrite(ctx, errs, err) {
					return filepath.SkipAll
				}
				return nil
			}

			var result DirEntry = &dirEntry{
				DirEntry: d,
				path: path,
			}
			if (!pipe.TryWrite(ctx, out, result)) {
				return filepath.SkipAll
			}
			return nil
		})
		if err != nil {
			// We can ignore the return value here because the function will return immediately either way.
			pipe.TryWrite(ctx, errs, err)
		}
	}()
	return out, errs
}
