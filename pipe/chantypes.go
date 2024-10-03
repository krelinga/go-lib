package pipe

// A type constraint that matches any channel which is readable.
// See mergeImpl() for an example of how this is used.
//
// It may make sense to impelment a similar constraint for writable channels in the future.
type readable[dataType any] interface {
	chan dataType | <-chan dataType
}
