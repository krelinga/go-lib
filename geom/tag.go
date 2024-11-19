package geom

import (
	"fmt"
	"iter"
	"maps"
)

// I first tried to use struct{} for the underlying type of tag, but
// it turns out that go doesn't guarantee that different zero-sized
// values will have distinct memory addresses.
type tag int

func getImpl[Seeking Element](c Element, t *tag) Seeking {
	found, ok := c.getByTag(t).(Seeking)
	fmt.Printf("found: %v, ok: %v\n", found, ok)
	return found
}

type publicTag interface {
	set(*tag)
}

type PointTag struct {
	t *tag
}

func (pt PointTag) Get(e Element) *Point {
	return getImpl[*Point](e, pt.t)
}

func (pt *PointTag) set(t *tag) {
	pt.t = t
}

type LineTag struct {
	t *tag
}

func (lt LineTag) Get(e Element) *Line {
	return getImpl[*Line](e, lt.t)
}

func (lt *LineTag) set(t *tag) {
	lt.t = t
}

type tagBase struct {
	tagIndex map[*tag]Element
}

func (tb *tagBase) getByTag(tag *tag) Element {
	if tb.tagIndex == nil {
		return nil
	}
	return tb.tagIndex[tag]
}

func (tb *tagBase) getTagIndex() iter.Seq2[*tag, Element] {
	return maps.All(tb.tagIndex)
}

func (tb *tagBase) addAllTags(in iter.Seq2[*tag, Element]) {
	if tb.tagIndex == nil {
		tb.tagIndex = make(map[*tag]Element)
	}
	for k, v := range in {
		tb.tagIndex[k] = v
	}
}

func addTags[PT publicTag](tb *tagBase, e Element, pts ...PT) {
	ptr := new(tag)
	fmt.Printf("ptr: %p\n", ptr)
	if tb.tagIndex == nil {
		tb.tagIndex = make(map[*tag]Element)
	}
	for _, pt := range pts {
		pt.set(ptr)
		tb.tagIndex[ptr] = e
	}
}
