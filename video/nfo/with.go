package nfo

import (
	"iter"
	"strconv"

	"github.com/beevik/etree"
)

type findExactlyOne struct {
	doc               *etree.Document
	path              etree.Path
	missing, multiple error
	validate          func(*etree.Element) error
	out               **etree.Element
}

func (feo findExactlyOne) find() error {
	found := feo.doc.FindElementsPath(feo.path)
	switch len(found) {
	case 0:
		return feo.missing
	case 1:
		*feo.out = found[0]
		if feo.validate != nil {
			return feo.validate(*feo.out)
		}
		return nil
	default:
		return feo.multiple
	}
}

func isPositiveInt(returnIfFailed error) func(*etree.Element) error {
	return func(e *etree.Element) error {
		i, err := strconv.Atoi(e.Text())
		if err != nil || i <= 0 {
			return returnIfFailed
		}
		return nil
	}
}

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
	return findExactlyOne{
		doc:      in,
		path:     path,
		missing:  ErrNoTitle,
		multiple: ErrMultipleTitles,
		out:      &wt.e,
	}.find()
}

// An NFO file with a title.
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
	err := findExactlyOne{
		doc:      in,
		path:     widthPath,
		missing:  ErrNoWidth,
		multiple: ErrMultipleWidths,
		validate: isPositiveInt(ErrInvalidWidth),
		out:      &wd.width,
	}.find()
	if err != nil {
		return err
	}

	return findExactlyOne{
		doc:      in,
		path:     heightPath,
		missing:  ErrNoHeight,
		multiple: ErrMultipleHeights,
		validate: isPositiveInt(ErrInvalidHeight),
		out:      &wd.height,
	}.find()
}

// An NFO file with dimensions.
type WithDimensions interface {
	Nfo
	Width() int
	Height() int
}

func elementIter(elems []*etree.Element) iter.Seq[string] {
	return func(yield func(v string) bool) {
		for _, elem := range elems {
			if !yield(elem.Text()) {
				return
			}
		}
	}
}

type withGenres struct {
	genres []*etree.Element
}

func (wg *withGenres) Genres() iter.Seq[string] {
	return elementIter(wg.genres)
}

func (wg *withGenres) init(in *etree.Document, path etree.Path) {
	wg.genres = in.FindElementsPath(path)
}

// An NFO file with genres.
type WithGeneres interface {
	Nfo
	Genres() iter.Seq[string]
}

type withTags struct {
	tags []*etree.Element
}

func (wt *withTags) Tags() iter.Seq[string] {
	return elementIter(wt.tags)
}

func (wt *withTags) init(in *etree.Document, path etree.Path) {
	wt.tags = in.FindElementsPath(path)
}

// An NFO file with tags.
type WithTags interface {
	Nfo
	Tags() iter.Seq[string]
}
