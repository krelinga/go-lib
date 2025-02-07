package nfo

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEpisode(t *testing.T) {
	tests := []struct {
		name       string
		nfo        string
		nfoPath    string
		wantErr    error
		wantTitle  string
		wantWidth  int
		wantHeight int
	}{
		{
			name:    "No title",
			nfo:     `<episodedetails></episodedetails>`,
			wantErr: ErrNoTitle,
		},
		{
			name:       "Single title",
			nfoPath:    "testdata/episode.nfo",
			wantTitle:  "Asteroid Blues",
			wantWidth:  720,
			wantHeight: 480,
		},
		{
			name:    "Multiple titles",
			nfo:     `<episodedetails><title>Title 1</title><title>Title 2</title></episodedetails>`,
			wantErr: ErrMultipleTitles,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r io.Reader
			if tt.nfoPath != "" {
				var err error
				r, err = os.Open(tt.nfoPath)
				if err != nil {
					t.Fatalf("Failed to open file: %v", err)
				}
			} else {
				r = strings.NewReader(tt.nfo)
			}
			out, err := Parse(r)
			assert.Equal(t, tt.wantErr, err, "parseEpisode() error")
			if err != nil {
				return
			}

			outEpisode, ok := out.(*Episode)
			if !ok {
				t.Fatal("Wrong type returned")
			}
			assert.Equal(t, tt.wantTitle, outEpisode.Title(), "title")
			assert.Equal(t, tt.wantWidth, outEpisode.Width(), "width")
			assert.Equal(t, tt.wantHeight, outEpisode.GetHeight(), "height")
		})
	}
}
