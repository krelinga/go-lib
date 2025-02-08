package nfo

import (
	"io"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMovie(t *testing.T) {
	tests := []struct {
		name        string
		nfo         string
		nfoPath     string
		wantErr     error
		wantTitle   string
		wantWidth   int
		wantHeight  int
		wantGeneres []string
	}{
		{
			name:        "Happy Path",
			nfoPath:     "testdata/movie.nfo",
			wantTitle:   "Ghostbusters",
			wantWidth:   3840,
			wantHeight:  2160,
			wantGeneres: []string{"Comedy", "Fantasy"},
		},
		{
			name:    "No title",
			nfo:     `<movie></movie>`,
			wantErr: ErrNoTitle,
		},
		{
			name:    "Multiple titles",
			nfo:     `<movie><title>Title 1</title><title>Title 2</title></movie>`,
			wantErr: ErrMultipleTitles,
		},
		{
			name:    "No width",
			nfo:     `<movie><title>Title</title></movie>`,
			wantErr: ErrNoWidth,
		},
		{
			name: "Invalid width",
			nfo: `
			<movie>
				<title>Title</title>
				<fileinfo>
					<streamdetails>
						<video>
							<width>invalid</width>
						</video>
					</streamdetails>
				</fileinfo>
			</movie>`,
			wantErr: ErrInvalidWidth,
		},
		{
			name: "Multiple widths",
			nfo: `
			<movie>
				<title>Title</title>
				<fileinfo>
					<streamdetails>
						<video>
							<width>720</width>
							<width>720</width>
						</video>
					</streamdetails>
				</fileinfo>
			</movie>`,
			wantErr: ErrMultipleWidths,
		},
		{
			name: "No height",
			nfo: `
			<movie>
				<title>Title</title>
				<fileinfo>
					<streamdetails>
						<video>
							<width>720</width>
						</video>
					</streamdetails>
				</fileinfo>
			</movie>`,
			wantErr: ErrNoHeight,
		},
		{
			name: "Invalid height",
			nfo: `
			<movie>
				<title>Title</title>
				<fileinfo>
					<streamdetails>
						<video>
							<width>720</width>
							<height>invalid</height>
						</video>
					</streamdetails>
				</fileinfo>
			</movie>`,
			wantErr: ErrInvalidHeight,
		},
		{
			name: "Multiple heights",
			nfo: `
			<movie>
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
			</movie>`,
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

			outMovie, ok := out.(*Movie)
			if !ok {
				t.Fatal("Wrong type returned")
			}
			assert.Equal(t, tt.wantTitle, outMovie.Title(), "title")
			assert.Equal(t, tt.wantWidth, outMovie.Width(), "width")
			assert.Equal(t, tt.wantHeight, outMovie.GetHeight(), "height")
			assert.Equal(t, tt.wantGeneres, slices.Collect(outMovie.Genres()), "genres")
		})
	}
}
