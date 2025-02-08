package nfo

import (
	"io"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadTvShow(t *testing.T) {
	tests := []struct {
		name        string
		nfo         string
		nfoPath     string
		wantErr     error
		wantTitle   string
		wantGeneres []string
		wantTags    []string
	}{
		{
			name:      "Happy Path",
			nfoPath:   "testdata/tvshow.nfo",
			wantTitle: "Cowboy Bebop",
			wantGeneres: []string{
				"Science Fiction",
				"Drama",
				"Comedy",
				"Animation",
				"Adventure",
				"Action",
				"Western",
				"Anime",
			},
			wantTags: []string{"planet mars", "spacecraft"},
		},
		{
			name:    "No title",
			nfo:     `<tvshow></tvshow>`,
			wantErr: ErrNoTitle,
		},
		{
			name:    "Multiple titles",
			nfo:     `<tvshow><title>Title 1</title><title>Title 2</title></tvshow>`,
			wantErr: ErrMultipleTitles,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r io.Reader
			if tt.nfoPath != "" {
				var err error
				r, err = os.Open(tt.nfoPath)
				require.NoError(t, err, "Failed to open file", tt.nfoPath)
			} else {
				r = strings.NewReader(tt.nfo)
			}
			nfo, err := ReadFrom(r)
			require.ErrorIs(t, err, tt.wantErr, "Failed to read NFO")
			if err != nil {
				return
			}

			tvshow, ok := nfo.(*TvShow)
			require.True(t, ok, "Failed to cast NFO")
			assert.Equal(t, tt.wantTitle, tvshow.Title(), "Title() mismatch")
			assert.Equal(t, tt.wantGeneres, slices.Collect(tvshow.Genres()), "Genres() mismatch")
			assert.Equal(t, tt.wantTags, slices.Collect(tvshow.Tags()), "Tags() mismatch")
		})
	}
}
