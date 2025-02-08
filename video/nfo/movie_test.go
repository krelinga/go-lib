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

func TestReadMovie(t *testing.T) {
	tests := []struct {
		name        string
		nfo         string
		nfoPath     string
		wantErr     error
		wantTitle   string
		wantWidth   int
		wantHeight  int
		wantGeneres []string
		wantTags    []string
	}{
		{
			name:        "Happy Path",
			nfoPath:     "testdata/movie.nfo",
			wantTitle:   "Ghostbusters",
			wantWidth:   3840,
			wantHeight:  2160,
			wantGeneres: []string{"Comedy", "Fantasy"},
			wantTags: []string{
				"new york city",
				"environmental protection agency",
				"library",
				"supernatural",
				"paranormal phenomena",
				"loser",
				"slime",
				"gatekeeper",
				"nerd",
				"giant monster",
				"haunting",
				"hybrid",
				"possession",
				"mythology",
				"horror spoof",
				"paranormal investigation",
				"urban setting",
				"super power",
				"receptionist",
				"world trade center",
				"ghost",
				"duringcreditsstinger",
				"ghostbusters",
			},
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
			out, err := ReadFrom(r)
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
			assert.Equal(t, tt.wantHeight, outMovie.Height(), "height")
			assert.Equal(t, tt.wantGeneres, slices.Collect(outMovie.Genres()), "genres")
			assert.Equal(t, tt.wantTags, slices.Collect(outMovie.Tags()), "tags")
		})
	}
}

func TestSetMovieTitle(t *testing.T) {
	nfoString := `
		<movie>
			<title>Title</title>
			<fileinfo>
				<streamdetails>
					<video>
						<width>720</width>
						<height>480</height>
					</video>
				</streamdetails>
			</fileinfo>
			<unrelated>Unrelated</unrelated>
		</movie>`
	nfo, err := ReadFrom(strings.NewReader(nfoString))
	require.NoError(t, err)
	movieNfo, ok := nfo.(*Movie)
	require.True(t, ok, "unexpected type")
	movieNfo.SetTitle("New Title")
	builder := &strings.Builder{}
	err = WriteTo(nfo, builder)
	require.NoError(t, err)
	wantNfoString := `
		<movie>
			<title>New Title</title>
			<fileinfo>
				<streamdetails>
					<video>
						<width>720</width>
						<height>480</height>
					</video>
				</streamdetails>
			</fileinfo>
			<unrelated>Unrelated</unrelated>
		</movie>`
	assert.Equal(t, wantNfoString, builder.String())
}
