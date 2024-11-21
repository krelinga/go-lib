package geom

import (
	"iter"
)

// I first tried to use struct{} for the underlying type of tag, but
// it turns out that go doesn't guarantee that different zero-sized
// values will have distinct memory addresses.
type tag int

func getImpl[Seeking Element](c Element, t *tag) Seeking {
	found, _ := c.getByTag(t).(Seeking)
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

type CircleTag struct {
	t *tag
}

func (ct *CircleTag) Get(e Element) *Circle {
	return getImpl[*Circle](e, ct.t)
}

func (ct *CircleTag) set(t *tag) {
	ct.t = t
}

func toPublicTagArray[PT publicTag](pts []PT) []publicTag {
	out := make([]publicTag, len(pts))
	for i, pt := range pts {
		out[i] = pt
	}
	return out
}

type tagSource bool

const (
	fromChild tagSource = true
	fromSelf  tagSource = false
)

type tagIndexValue struct {
	e         Element
	tagSource tagSource
}

type tagIndex = map[*tag]tagIndexValue

type tagBase struct {
	tagIndex tagIndex
}

func (tb *tagBase) getByTag(tag *tag) Element {
	if tb.tagIndex == nil {
		return nil
	}
	value, ok := tb.tagIndex[tag]
	if !ok {
		return nil
	}
	return value.e
}

func (tb *tagBase) getTagIndex() iter.Seq2[*tag, Element] {
	return func(yield func(*tag, Element) bool) {
		for k, v := range tb.tagIndex {
			if !yield(k, v.e) {
				return
			}
		}
	}
}

func (tb *tagBase) getSelfTags() iter.Seq[*tag] {
	return func(yield func(*tag) bool) {
		for k, v := range tb.tagIndex {
			if v.tagSource == fromSelf {
				if !yield(k) {
					return
				}
			}
		}
	}
}

func (tb *tagBase) addChildTags(in iter.Seq2[*tag, Element]) {
	if tb.tagIndex == nil {
		tb.tagIndex = make(tagIndex)
	}
	for k, v := range in {
		tb.tagIndex[k] = tagIndexValue{v, fromChild}
	}
}

func (tb *tagBase) addSelfTags(e Element, in iter.Seq[*tag]) {
	if tb.tagIndex == nil {
		tb.tagIndex = make(tagIndex)
	}
	for t := range in {
		tb.tagIndex[t] = tagIndexValue{
			e:         e,
			tagSource: fromSelf,
		}
	}
}

// TODO: does this need to be a generic function?
func (tb *tagBase) addPublicTags(e Element, pts []publicTag) {
	ptr := new(tag)
	if tb.tagIndex == nil {
		tb.tagIndex = make(tagIndex)
	}
	for _, pt := range pts {
		pt.set(ptr)
		tb.tagIndex[ptr] = tagIndexValue{
			e:         e,
			tagSource: fromSelf,
		}
	}
}
