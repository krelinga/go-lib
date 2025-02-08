package filesystem

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/krelinga/go-lib/pipe"
	"github.com/stretchr/testify/assert"
)

func TestWalkAll(t *testing.T) {
	t.Parallel()

	// Setup a temporary directory structure for testing
	root := t.TempDir()
	subDir := filepath.Join(root, "subdir")
	err := os.Mkdir(subDir, 0755)
	assert.NoError(t, err)

	file1 := filepath.Join(root, "file1.txt")
	err = os.WriteFile(file1, []byte("content1"), 0644)
	assert.NoError(t, err)

	file2 := filepath.Join(subDir, "file2.txt")
	err = os.WriteFile(file2, []byte("content2"), 0644)
	assert.NoError(t, err)

	t.Run("ToCompletion", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		files, errs := WalkAll(ctx, root)

		var entries []DirEntry
		var errors []error
		pipe.Wait(
			pipe.ToArrayFunc(files, &entries),
			pipe.ToArrayFunc(errs, &errors),
		)

		assert.Empty(t, errors)
		expectedPaths := []string{root, file1, subDir, file2}
		assert.Len(t, entries, len(expectedPaths))

		for _, entry := range entries {
			_ = entry.Path()    // Ensure the method is implemented.
			_ = entry.IsDir()   // Ensure the method is implemented.
			_ = entry.Name()    // Ensure the method is implemented.
			_ = entry.Type()    // Ensure the method is implemented.
			_, _ = entry.Info() // Ensure the method is implemented.
			assert.Contains(t, expectedPaths, entry.Path())
		}
	})
	t.Run("ContextCancelled", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel right away.
		files, errs := WalkAll(ctx, root)

		var entries []DirEntry
		var errors []error
		pipe.Wait(
			pipe.ToArrayFunc(files, &entries),
			pipe.ToArrayFunc(errs, &errors),
		)

		assert.Empty(t, entries)
		assert.Empty(t, errors)
	})
}
