package video

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirKind(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		dir         string
		setup       func(parentDir string) error
		expected    DirKind
		expectError bool
	}{
		{
			name: "Movie directory",
			dir:  "movie",
			setup: func(parentDir string) error {
				return os.MkdirAll(parentDir+"/movie", 0755)
			},
			expected:    DirKindMovie,
			expectError: false,
		},
		{
			name: "Show directory",
			dir:  "show",
			setup: func(parentDir string) error {
				err := os.MkdirAll(parentDir+"/show", 0755)
				if err != nil {
					return err
				}
				_, err = os.Create(parentDir + "/show/tvshow.nfo")
				return err
			},
			expected:    DirKindShow,
			expectError: false,
		},
		{
			name: "Non-directory path",
			dir:  "file",
			setup: func(parentDir string) error {
				_, err := os.Create(parentDir + "/file")
				return err
			},
			expected:    0,
			expectError: true,
		},
		{
			name: "Non-existent path",
			dir:  "nonexistent",
			setup: func(parentDir string) error {
				return nil
			},
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Setup test environment
			tempDir, err := os.MkdirTemp("", "test")
			if err != nil {
				t.Fatalf("failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tempDir)

			tt.dir = tempDir + "/" + tt.dir
			if err := tt.setup(tempDir); err != nil {
				t.Fatalf("setup failed: %v", err)
			}

			// Run the function under test
			result, err := GetDirKind(tt.dir)

			// Verify results
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
