package nfo

import (
	"errors"
	"fmt"
	"io"

	"github.com/beevik/etree"
)

type Nfo interface{}

var (
	ErrBadXml           = errors.New("invalid xml")
	ErrBadRootNamespace = errors.New("root tag namespace should be empty")
	ErrBadRootTag       = errors.New("unexpected root tag")
)

func Parse(in io.Reader) (Nfo, error) {
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
		return parseMovie(doc)
	case "tvshow":
		return parseShow(doc)
	case "episodedetails":
		return parseEpisode(doc)
	default:
		return nil, fmt.Errorf("%w: %s", ErrBadRootTag, root.Tag)
	}
}
