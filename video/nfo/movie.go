package nfo

import "github.com/beevik/etree"

type Movie struct {
	// TODO
}

func (*Movie) validNfoSubtype() {}

func parseMovie(_ *etree.Document) (*Movie, error) {
	return &Movie{}, nil
}