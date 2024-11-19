package geom

type Path interface {
	Element

	// The endpoints that are connected by the path.
	Endpoints() (*Point, *Point)

	pathIsAClosedType()
}