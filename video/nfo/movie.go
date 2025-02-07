package nfo

import (
	"fmt"
	"strconv"

	"github.com/beevik/etree"
)

type Movie struct {
	title         *etree.Element
	width, height *etree.Element
}

func (*Movie) validNfoSubtype() {}

var (
	pathMovieTitle  = etree.MustCompilePath("/movie/title")
	pathMovieWidth  = etree.MustCompilePath("/movie/fileinfo/streamdetails/video/width")
	pathMovieHeight = etree.MustCompilePath("/movie/fileinfo/streamdetails/video/height")
)

func parseMovie(doc *etree.Document) (*Movie, error) {
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

	return movie, nil
}
