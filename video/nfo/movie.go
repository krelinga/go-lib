package nfo

import (
	"github.com/beevik/etree"
)

type Movie struct {
	withTitle
	withDimensions
	withGenres
	withTags
	withEdition
	document *etree.Document
}

func (*Movie) validNfoSubtype() {}

func (m *Movie) getDocument() *etree.Document {
	return m.document
}

var (
	pathMovieTitle   = etree.MustCompilePath("/movie/title")
	pathMovieWidth   = etree.MustCompilePath("/movie/fileinfo/streamdetails/video/width")
	pathMovieHeight  = etree.MustCompilePath("/movie/fileinfo/streamdetails/video/height")
	pathMovieGenre   = etree.MustCompilePath("/movie/genre")
	pathMovieTag     = etree.MustCompilePath("/movie/tag")
	pathMovieEdition = etree.MustCompilePath("/movie/edition")
)

func readMovie(doc *etree.Document) (*Movie, error) {
	movie := &Movie{
		document: doc,
	}

	if err := movie.withTitle.init(doc, pathMovieTitle); err != nil {
		return nil, err
	}
	if err := movie.withDimensions.init(doc, pathMovieWidth, pathMovieHeight); err != nil {
		return nil, err
	}
	movie.withGenres.init(doc, pathMovieGenre)
	movie.withTags.init(doc, pathMovieTag)
	movie.withEdition.init(doc, pathMovieEdition)

	return movie, nil
}
