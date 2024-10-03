package video

import (
	"context"
	"strings"
	"testing"

	"github.com/krelinga/go-lib/filesystem"
	"github.com/krelinga/go-lib/filesystem/filesystemtest"
	"github.com/stretchr/testify/assert"
)

func TestPaths(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *Paths
	}{
		{
			name:  "NewPathsFromMkv valid input",
			input: "example.mkv",
			expected: &Paths{
				mkv:       "example.mkv",
				nfo:       "example.nfo",
				tcProfile: "example.tcprofile",
			},
		},
		{
			name:  "NewPathsFromNfo valid input",
			input: "example.nfo",
			expected: &Paths{
				mkv:       "example.mkv",
				nfo:       "example.nfo",
				tcProfile: "example.tcprofile",
			},
		},
		{
			name:  "NewPathsFromTcProfile valid input",
			input: "example.tcprofile",
			expected: &Paths{
				mkv:       "example.mkv",
				nfo:       "example.nfo",
				tcProfile: "example.tcprofile",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result *Paths
			switch {
			case strings.HasSuffix(tt.input, ".mkv"):
				result = NewPathsFromMkv(tt.input)
			case strings.HasSuffix(tt.input, ".nfo"):
				result = NewPathsFromNfo(tt.input)
			case strings.HasSuffix(tt.input, ".tcprofile"):
				result = NewPathsFromTcProfile(tt.input)
			default:
				assert.Fail(t, "invalid input")
			}

			assert.Equal(t, tt.expected, result)
		})
	}

	t.Run("NewPathsFromMkv invalid input", func(t *testing.T) {
		assert.Panics(t, func() {
			NewPathsFromMkv("example.txt")
		})
	})

	t.Run("NewPathsFromNfo invalid input", func(t *testing.T) {
		assert.Panics(t, func() {
			NewPathsFromNfo("example.txt")
		})
	})

	t.Run("NewPathsFromTcProfile invalid input", func(t *testing.T) {
		assert.Panics(t, func() {
			NewPathsFromTcProfile("example.txt")
		})
	})
}

func TestBuildFileInfo(t *testing.T) {
	tests := []struct {
		name     string
		entries  []filesystem.DirEntry
		expected []*FileInfo
	}{
		{
			name: "Single MKV file",
			entries: []filesystem.DirEntry{
				filesystemtest.NewMockDirEntry("example.mkv", false),
			},
			expected: []*FileInfo{
				{
					MkvPath:      "example.mkv",
					HasNfo:       false,
					HasTcProfile: false,
				},
			},
		},
		{
			name: "MKV with NFO and TcProfile",
			entries: []filesystem.DirEntry{
				filesystemtest.NewMockDirEntry("example.mkv", false),
				filesystemtest.NewMockDirEntry("example.nfo", false),
				filesystemtest.NewMockDirEntry("example.tcprofile", false),
			},
			expected: []*FileInfo{
				{
					MkvPath:      "example.mkv",
					HasNfo:       true,
					HasTcProfile: true,
				},
			},
		},
		{
			name: "Multiple MKV files with mixed related files",
			entries: []filesystem.DirEntry{
				filesystemtest.NewMockDirEntry("example1.mkv", false),
				filesystemtest.NewMockDirEntry("example1.nfo", false),
				filesystemtest.NewMockDirEntry("example2.mkv", false),
				filesystemtest.NewMockDirEntry("example2.tcprofile", false),
			},
			expected: []*FileInfo{
				{
					MkvPath:      "example1.mkv",
					HasNfo:       true,
					HasTcProfile: false,
				},
				{
					MkvPath:      "example2.mkv",
					HasNfo:       false,
					HasTcProfile: true,
				},
			},
		},
		{
			name: "No MKV files",
			entries: []filesystem.DirEntry{
				filesystemtest.NewMockDirEntry("example.nfo", false),
				filesystemtest.NewMockDirEntry("example.tcprofile", false),
			},
			expected: []*FileInfo{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			files := make(chan filesystem.DirEntry, len(tt.entries))
			for _, entry := range tt.entries {
				files <- entry
			}
			close(files)

			resultChan := BuildFileInfo(ctx, files)
			result := []*FileInfo{}
			for info := range resultChan {
				result = append(result, info)
			}

			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}
