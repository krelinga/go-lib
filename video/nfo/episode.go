package nfo

import (
	"errors"
	"strconv"

	"github.com/beevik/etree"
)

type Episode struct {
	withTitle
	withDimensions
}

func (*Episode) validNfoSubtype() {}

func (e *Episode) Width() int {
	i, err := strconv.Atoi(e.width.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func (e *Episode) GetHeight() int {
	i, err := strconv.Atoi(e.height.Text())
	if err != nil {
		panic(err)
	}
	return i
}

var (
	pathEpisodeTitle  = etree.MustCompilePath("/episodedetails/title")
	pathEpisodeWidth  = etree.MustCompilePath("/episodedetails/fileinfo/streamdetails/video/width")
	pathEpisodeHeight = etree.MustCompilePath("/episodedetails/fileinfo/streamdetails/video/height")
)

var (
	ErrNoTitle         = errors.New("no title found")
	ErrMultipleTitles  = errors.New("multiple titles found")
	ErrNoWidth         = errors.New("no width found")
	ErrMultipleWidths  = errors.New("multiple widths found")
	ErrInvalidWidth    = errors.New("invalid width, must be a positive integer")
	ErrNoHeight        = errors.New("no height found")
	ErrMultipleHeights = errors.New("multiple heights found")
	ErrInvalidHeight   = errors.New("invalid height, must be a positive integer")
)

func readEpisode(doc *etree.Document) (*Episode, error) {
	episode := &Episode{}

	if err := episode.withTitle.init(doc, pathEpisodeTitle); err != nil {
		return nil, err
	}
	if err := episode.withDimensions.init(doc, pathEpisodeWidth, pathEpisodeHeight); err != nil {
		return nil, err
	}

	return episode, nil
}
