package nfo

import (
	"iter"
	"strconv"

	"github.com/beevik/etree"
)

type Movie struct {
	withTitle
	withDimensions
	withGenres
	withTags
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

	if err := movie.withTitle.init(doc, pathMovieTitle); err != nil {
		return nil, err
	}
	if err := movie.withDimensions.init(doc, pathMovieWidth, pathMovieHeight); err != nil {
		return nil, err
	}
	movie.withGenres.init(doc, pathMovieGenre)
	movie.withTags.init(doc, pathMovieTag)

	return movie, nil
}
