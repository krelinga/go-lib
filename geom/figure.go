package geom

// A figure is a closed shape.
type Figure interface {
	Element

	figureIsAClosedType()
}