package nfo

import "github.com/beevik/etree"

type withTitle struct {
	e *etree.Element
}

func (wt *withTitle) Title() string {
	return wt.e.Text()
}

func (wt *withTitle) SetTitle(title string) {
	wt.e.SetText(title)
}

func (wt *withTitle) init(in *etree.Document, path etree.Path, errMissing, errMultiple error) error {
	found := in.FindElementsPath(path)
	switch len(found) {
	case 0:
		return errMissing
	case 1:
		wt.e = found[0]
	default:
		return errMultiple
	}
	return nil
}

type WithTitle interface {
	Nfo

	Title() string
	SetTitle(title string)
}
