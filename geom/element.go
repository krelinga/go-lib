package geom

import "iter"

type Element interface {
	BoundingBox() BoundingBox

	getByTag(*tag) Element
	getTagIndex() iter.Seq2[*tag, Element]
}
