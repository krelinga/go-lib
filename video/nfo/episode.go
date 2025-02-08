package nfo

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/beevik/etree"
)

type Episode struct {
	title         *etree.Element
	width, height *etree.Element
}

func (*Episode) validNfoSubtype() {}

func (e *Episode) Title() string {
	return e.title.Text()
}

func (e *Episode) SetTitle(title string) {
	e.title.SetText(title)
}

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

	titles := doc.FindElementsPath(pathEpisodeTitle)
	switch len(titles) {
	case 0:
		return nil, ErrNoTitle
	case 1:
		episode.title = titles[0]
	default:
		return nil, ErrMultipleTitles
	}

	widths := doc.FindElementsPath(pathEpisodeWidth)
	switch len(widths) {
	case 0:
		return nil, ErrNoWidth
	case 1:
		episode.width = widths[0]
		if i, err := strconv.Atoi(episode.width.Text()); err != nil || i <= 0 {
			return nil, fmt.Errorf("%w: %s", ErrInvalidWidth, episode.width.Text())
		}
	default:
		return nil, ErrMultipleWidths
	}

	heights := doc.FindElementsPath(pathEpisodeHeight)
	switch len(heights) {
	case 0:
		return nil, ErrNoHeight
	case 1:
		episode.height = heights[0]
		if i, err := strconv.Atoi(episode.height.Text()); err != nil || i <= 0 {
			return nil, fmt.Errorf("%w: %s", ErrInvalidHeight, episode.height.Text())
		}
	default:
		return nil, ErrMultipleHeights
	}

	return episode, nil
}
