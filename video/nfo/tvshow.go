package nfo

import "github.com/beevik/etree"

type TvShow struct {
	// TODO
	document *etree.Document
}

func (*TvShow) validNfoSubtype() {}

func (s *TvShow) getDocument() *etree.Document {
	return s.document
}

func readTvShow(_ *etree.Document) (*TvShow, error) {
	tvShow := &TvShow{
		document: nil,
	}
	return tvShow, nil
}
