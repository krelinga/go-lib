package nfo

import (
	"fmt"
	"strconv"

	"github.com/beevik/etree"
)

type withTitle struct {
	e *etree.Element
}

func (wt *withTitle) Title() string {
	return wt.e.Text()
}

func (wt *withTitle) SetTitle(title string) {
	wt.e.SetText(title)
}

func (wt *withTitle) init(in *etree.Document, path etree.Path) error {
	found := in.FindElementsPath(path)
	switch len(found) {
	case 0:
		return ErrNoTitle
	case 1:
		wt.e = found[0]
	default:
		return ErrMultipleTitles
	}
	return nil
}

type WithTitle interface {
	Nfo

	Title() string
	SetTitle(title string)
}

type withDimensions struct {
	width, height *etree.Element
}

func (wd *withDimensions) Width() int {
	i, err := strconv.Atoi(wd.width.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func (wd *withDimensions) Height() int {
	i, err := strconv.Atoi(wd.height.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func (wd *withDimensions) init(in *etree.Document, widthPath, heightPath etree.Path) error {
	widths := in.FindElementsPath(widthPath)
	switch len(widths) {
	case 0:
		return ErrNoWidth
	case 1:
		wd.width = widths[0]
		if i, err := strconv.Atoi(wd.width.Text()); err != nil || i <= 0 {
			return fmt.Errorf("%w: %s", ErrInvalidWidth, wd.width.Text())
		}
	default:
		return ErrMultipleWidths
	}

	heights := in.FindElementsPath(heightPath)
	switch len(heights) {
	case 0:
		return ErrNoHeight
	case 1:
		wd.height = heights[0]
		if i, err := strconv.Atoi(wd.height.Text()); err != nil || i <= 0 {
			return fmt.Errorf("%w: %s", ErrInvalidHeight, wd.height.Text())
		}
	default:
		return ErrMultipleHeights
	}

	return nil
}
