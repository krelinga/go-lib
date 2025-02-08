package nfo

import "github.com/beevik/etree"

type Show struct {
	// TODO
}

func (*Show) validNfoSubtype() {}

func readShow(_ *etree.Document) (*Show, error) {
	return &Show{}, nil
}
