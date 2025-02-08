package nfo

import (
	"fmt"
	"iter"
	"strconv"

	"github.com/beevik/etree"
)

type Movie struct {
	title         *etree.Element
	width, height *etree.Element
	genres        []*etree.Element
	tags          []*etree.Element
}

func (*Movie) validNfoSubtype() {}

func (m *Movie) Title() string {
	return m.title.Text()
}

func (m *Movie) SetTitle(title string) {
	m.title.SetText(title)
}

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

	titles := doc.FindElementsPath(pathMovieTitle)
	switch len(titles) {
	case 0:
		return nil, ErrNoTitle
	case 1:
		movie.title = titles[0]
	default:
		return nil, ErrMultipleTitles
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
