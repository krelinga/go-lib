package nfo

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadFrom(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		inputFile string
		wantErr   error
		typeTest  func(Nfo) bool
	}{
		{
			name:      "Valid movie XML",
			inputFile: "testdata/movie.nfo",
			wantErr:   nil,
			typeTest:  func(n Nfo) bool { _, ok := n.(*Movie); return ok },
		},
		{
			name:      "Valid show XML",
			inputFile: "testdata/show.nfo",
			wantErr:   nil,
			typeTest:  func(n Nfo) bool { _, ok := n.(*Show); return ok },
		},
		{
			name:      "Valid episode XML",
			inputFile: "testdata/episode.nfo",
			wantErr:   nil,
			typeTest:  func(n Nfo) bool { _, ok := n.(*Episode); return ok },
		},
		{
			name:    "Invalid root namespace",
			input:   `<ns:movie xmlns:ns="namespace"></ns:movie>`,
			wantErr: ErrBadRootNamespace,
		},
		{
			name:    "Invalid root tag",
			input:   `<invalid></invalid>`,
			wantErr: ErrBadRootTag,
		},
		{
			name:    "Invalid XML",
			input:   `<movie>`,
			wantErr: ErrBadXml,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r io.Reader
			if tt.inputFile != "" {
				var err error
				r, err = os.Open(tt.inputFile)
				if err != nil {
					t.Fatalf("failed to open file %s: %v", tt.inputFile, err)
				}
			} else {
				r = strings.NewReader(tt.input)
			}
			nfo, err := ReadFrom(r)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}

			if tt.typeTest != nil {
				assert.True(t, tt.typeTest(nfo), "unexpected type")
			}
		})
	}
}
