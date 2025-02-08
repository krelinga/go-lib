package nfo

import (
	"errors"
	"fmt"
	"io"

	"github.com/beevik/etree"
)

// Nfo will always be an instance of *Movie, *TvShow, or *Episode.
type Nfo interface {
	getDocument() *etree.Document
	validNfoSubtype()
}

var (
	ErrBadXml           = errors.New("invalid xml")
	ErrBadRootNamespace = errors.New("root tag namespace should be empty")
	ErrBadRootTag       = errors.New("unexpected root tag")
)

// ReadFrom reads an NFO file from the given reader and returns the parsed data.
func ReadFrom(in io.Reader) (Nfo, error) {
	doc := etree.NewDocument()
	if _, err := doc.ReadFrom(in); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrBadXml, err)
	}

	root := doc.Root()
	if root.Space != "" {
		return nil, fmt.Errorf("%w: %s", ErrBadRootNamespace, root.Space)
	}
	switch root.Tag {
	case "movie":
		return readMovie(doc)
	case "tvshow":
		return readTvShow(doc)
	case "episodedetails":
		return readEpisode(doc)
	default:
		return nil, fmt.Errorf("%w: %s", ErrBadRootTag, root.Tag)
	}
}
