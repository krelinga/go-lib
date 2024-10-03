package video

import (
	"encoding/json"
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
func TestDirKindJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		dirKind     DirKind
		expected    string
		expectError bool
	}{
		{
			name:        "Marshal Movie",
			dirKind:     DirKindMovie,
			expected:    `"movie"`,
			expectError: false,
		},
		{
			name:        "Marshal Show",
			dirKind:     DirKindShow,
			expected:    `"show"`,
			expectError: false,
		},
		{
			name:        "Unmarshal Movie",
			dirKind:     DirKindMovie,
			expected:    `"movie"`,
			expectError: false,
		},
		{
			name:        "Unmarshal Show",
			dirKind:     DirKindShow,
			expected:    `"show"`,
			expectError: false,
		},
		{
			name:        "Unmarshal Invalid",
			dirKind:     0,
			expected:    `"invalid"`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.name[:7] == "Marshal" {
				// Test JSON marshalling
				result, err := json.Marshal(tt.dirKind)
				if tt.expectError {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.JSONEq(t, tt.expected, string(result))
				}
			} else {
				// Test JSON unmarshalling
				var dk DirKind
				err := json.Unmarshal([]byte(tt.expected), &dk)
				if tt.expectError {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.dirKind, dk)
				}
			}
		})
	}
}
