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
			name:       "Happy Path",
			nfoPath:    "testdata/episode.nfo",
			wantTitle:  "Asteroid Blues",
			wantWidth:  720,
			wantHeight: 480,
		},
		{
			name:    "No title",
			nfo:     `<episodedetails></episodedetails>`,
			wantErr: ErrNoTitle,
		},
		{
			name:    "Multiple titles",
			nfo:     `<episodedetails><title>Title 1</title><title>Title 2</title></episodedetails>`,
			wantErr: ErrMultipleTitles,
		},
		{
			name:    "No width",
			nfo:     `<episodedetails><title>Title</title></episodedetails>`,
			wantErr: ErrNoWidth,
		},
		{
			name: "Invalid width",
			nfo: `
			<episodedetails>
				<title>Title</title>
				<fileinfo>
					<streamdetails>
						<video>
							<width>invalid</width>
						</video>
					</streamdetails>
				</fileinfo>
			</episodedetails>`,
			wantErr: ErrInvalidWidth,
		},
		{
			name: "Multiple widths",
			nfo: `
			<episodedetails>
				<title>Title</title>
				<fileinfo>
					<streamdetails>
						<video>
							<width>720</width>
							<width>720</width>
						</video>
					</streamdetails>
				</fileinfo>
			</episodedetails>`,
			wantErr: ErrMultipleWidths,
		},
		{
			name: "No height",
			nfo: `
			<episodedetails>
				<title>Title</title>
				<fileinfo>
					<streamdetails>
						<video>
							<width>720</width>
						</video>
					</streamdetails>
				</fileinfo>
			</episodedetails>`,
			wantErr: ErrNoHeight,
		},
		{
			name: "Invalid height",
			nfo: `
			<episodedetails>
				<title>Title</title>
				<fileinfo>
					<streamdetails>
						<video>
							<width>720</width>
							<height>invalid</height>
						</video>
					</streamdetails>
				</fileinfo>
			</episodedetails>`,
			wantErr: ErrInvalidHeight,
		},
		{
			name: "Multiple heights",
			nfo: `
			<episodedetails>
				<title>Title</title>
				<fileinfo>
					<streamdetails>
						<video>
							<width>720</width>
							<height>480</height>
							<height>480</height>
						</video>
					</streamdetails>
				</fileinfo>
			</episodedetails>`,
			wantErr: ErrMultipleHeights,
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
			assert.ErrorIs(t, err, tt.wantErr)
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
