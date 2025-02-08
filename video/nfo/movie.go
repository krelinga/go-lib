package nfo

import (
	"fmt"
	"iter"
	"strconv"

	"github.com/beevik/etree"
)

type Movie struct {
	withTitle
	width, height *etree.Element
	genres        []*etree.Element
	tags          []*etree.Element
}

func (*Movie) validNfoSubtype() {}

func (m *Movie) Width() int {
	i, err := strconv.Atoi(m.width.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func (m *Movie) GetHeight() int {
	i, err := strconv.Atoi(m.height.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func (m *Movie) Genres() iter.Seq[string] {
	return func(yield func(v string) bool) {
		for _, genre := range m.genres {
			if !yield(genre.Text()) {
				return
			}
		}
	}
}

func (m *Movie) Tags() iter.Seq[string] {
	return func(yield func(v string) bool) {
		for _, tag := range m.tags {
			if !yield(tag.Text()) {
				return
			}
		}
	}
}

var (
	pathMovieTitle  = etree.MustCompilePath("/movie/title")
	pathMovieWidth  = etree.MustCompilePath("/movie/fileinfo/streamdetails/video/width")
	pathMovieHeight = etree.MustCompilePath("/movie/fileinfo/streamdetails/video/height")
	pathMovieGenre  = etree.MustCompilePath("/movie/genre")
	pathMovieTag    = etree.MustCompilePath("/movie/tag")
)

func readMovie(doc *etree.Document) (*Movie, error) {
	movie := &Movie{}

	if err := movie.withTitle.init(doc, pathMovieTitle, ErrNoTitle, ErrMultipleTitles); err != nil {
		return nil, err
	}

	widths := doc.FindElementsPath(pathMovieWidth)
	switch len(widths) {
	case 0:
		return nil, ErrNoWidth
	case 1:
		movie.width = widths[0]
		if i, err := strconv.Atoi(movie.width.Text()); err != nil || i <= 0 {
			return nil, fmt.Errorf("%w: %s", ErrInvalidWidth, movie.width.Text())
		}
	default:
		return nil, ErrMultipleWidths
	}

	heights := doc.FindElementsPath(pathMovieHeight)
	switch len(heights) {
	case 0:
		return nil, ErrNoHeight
	case 1:
		movie.height = heights[0]
		if i, err := strconv.Atoi(movie.height.Text()); err != nil || i <= 0 {
			return nil, fmt.Errorf("%w: %s", ErrInvalidHeight, movie.height.Text())
		}
	default:
		return nil, ErrMultipleHeights
	}

	movie.genres = doc.FindElementsPath(pathMovieGenre)
	movie.tags = doc.FindElementsPath(pathMovieTag)

	return movie, nil
}
