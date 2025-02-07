package nfo

import (
	"errors"

	"github.com/beevik/etree"
)

type Episode struct {
	title *etree.Element
}

func (*Episode) validNfoSubtype() {}

func (e *Episode) Title() string {
	return e.title.Text()
}

func (e *Episode) SetTitle(title string) {
	e.title.SetText(title)
}

var (
	pathEpisodeTitle = etree.MustCompilePath("/episodedetails/title")
)

var (
	ErrNoTitle        = errors.New("no title found")
	ErrMultipleTitles = errors.New("multiple titles found")
)

func parseEpisode(doc *etree.Document) (*Episode, error) {
	episode := &Episode{}

	titles := doc.FindElementsPath(pathEpisodeTitle)
	switch len(titles) {
	case 0:
		return nil, ErrNoTitle
	case 1:
		episode.title = titles[0]
	default:
		return nil, ErrMultipleTitles
	}
	return episode, nil
}
