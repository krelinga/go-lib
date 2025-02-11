package podreflect

type POD interface {
	// Private to use as a marker interface
	isPodType()
}

type Struct struct {}

func (s Struct) isPodType() {}