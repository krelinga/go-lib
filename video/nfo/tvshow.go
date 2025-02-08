package nfo

import "github.com/beevik/etree"

type TvShow struct {
	withTitle
	withTags
	withGenres
	document *etree.Document
}

func (*TvShow) validNfoSubtype() {}

func (s *TvShow) getDocument() *etree.Document {
	return s.document
}

var (
	pathTvShowTitle = etree.MustCompilePath("/tvshow/title")
	pathTvShowGenre = etree.MustCompilePath("/tvshow/genre")
	pathTvShowTag   = etree.MustCompilePath("/tvshow/tag")
)

func readTvShow(doc *etree.Document) (*TvShow, error) {
	tvShow := &TvShow{
		document: doc,
	}

	if err := tvShow.withTitle.init(doc, pathTvShowTitle); err != nil {
		return nil, err
	}
	tvShow.withGenres.init(doc, pathTvShowGenre)
	tvShow.withTags.init(doc, pathTvShowTag)

	return tvShow, nil
}
