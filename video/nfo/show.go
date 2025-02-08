package nfo

import "github.com/beevik/etree"

type Show struct {
	// TODO
	document *etree.Document
}

func (*Show) validNfoSubtype() {}

func (s *Show) getDocument() *etree.Document {
	return s.document
}

func readShow(_ *etree.Document) (*Show, error) {
	show := &Show{
		document: nil,
	}
	return show, nil
}
