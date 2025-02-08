package nfo

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompose(t *testing.T) {
	type myType interface {
		WithTitle
		WithDimensions
	}
	tests := []struct {
		name       string
		nfoPath    string
		wantTitle  string
		wantWidth  int
		wantHeight int
	}{
		{
			name:       "Movie",
			nfoPath:    "testdata/movie.nfo",
			wantTitle:  "Ghostbusters",
			wantWidth:  3840,
			wantHeight: 2160,
		}, {
			name:       "Episode",
			nfoPath:    "testdata/episode.nfo",
			wantTitle:  "Asteroid Blues",
			wantWidth:  720,
			wantHeight: 480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, err := os.Open(tt.nfoPath)
			require.NoError(t, err, "Failed to open file", tt.nfoPath)
			out, err := ReadFrom(in)
			require.NoError(t, err, "Failed to read NFO")
			my, ok := out.(myType)
			require.True(t, ok, "Failed to cast NFO")
			require.Equal(t, tt.wantTitle, my.Title(), "Title() mismatch")
			require.Equal(t, tt.wantWidth, my.Width(), "Width() mismatch")
			require.Equal(t, tt.wantHeight, my.Height(), "Height() mismatch")

			my.SetTitle("New Title")
			builder := &strings.Builder{}
			err = WriteTo(my, builder)
			assert.NoError(t, err, "Failed to write NFO")
		})
	}
}
