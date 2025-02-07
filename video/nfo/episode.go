package nfo

import "github.com/beevik/etree"

type Episode struct {
	// TODO
}

func (*Episode) validNfoSubtype() {}

func parseEpisode(_ *etree.Document) (*Episode, error) {
	return &Episode{}, nil
}
