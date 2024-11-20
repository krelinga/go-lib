package geom

// Desirable properties:
// - Type Preservation: If I pass in a Point, I should get a Point back.
// - Immutability: The original object should not be modified.
// - Chaining: it should be possible to combine multiple transformations
//             without needing to do a bunch of extra allocations.
// - Tag Preservation: tags created on the pre-transformed object should
//                     still work on the post-transformed object.
// - Consistent: transformations should work similarly across all types.

/*
API Options:
1) Global Transform() method.
  - naiive implementation fails at type preservation.
  - could use type assertions to get around this, but that's not ideal.
  - could use generics to work around this, but the problem there is that
    go doesn't allow type assertions on params constrained by generic types.

2) Method on each type.
  - how many types are we talking about here?
    -  Point, Path, Figure.
	- I think that's it.  There are several kinds of Paths, but we can
	  deal with all of them at the generic level and rely on type switches
	  to handle the specifics.
	- Right now we're surfacing Line ptrs directly, but I don't think that's
	  necessary.  We could just use the Path interface instead.
	- Currently, Point implements Path, but that isn't strictly necessary,
	  and might even be a good simplification to get rid of ... that would
	  eliminate a lot of edge cases.
	- so, if we make those changes, we only have the following types:
		- Point, where Transform() returns a *Point.
		- Figure, where Transform() returns a *Figure.
		- Path, where Transform() returns a Path (interface ptr).
	- there's also room in this plan for type-specific helper functions
	  that can be called from the public Transform() methods.


#2 sounds good, let's go with that.

Upon some further experimentation, it looks like there's a 3rd option that
involves using a generic Transform() method at the global scope that wraps
a weakly-typed call on an interface hierarchy with a cast.  Something like:

type foo interface {
	clone() foo
}

type intFoo int

func newIntFoo(x int) *intFoo {
	y := new(intFoo)
	*y = intFoo(x)
	return y
}

func (x *intFoo) clone() foo {
	cloned := new(intFoo)
	*cloned = intFoo(int(*x))
	return cloned
}

type stringFoo string

func newStringFoo(x string) *stringFoo {
	y := new(stringFoo)
	*y = stringFoo(x)
	return y
}

func (x *stringFoo) clone() foo {
	cloned := new(stringFoo)
	*cloned = stringFoo(string(*x))
	return cloned
}


func Clone[X foo](in X) X {
	cloned := in.clone()
	return cloned.(X)
}
*/

func clone[E Element](in E) E {
	return in.clone().(E)
}

type Transformation func(Element)

func Translate(dx, dy float64) Transformation {
	return func(e Element) {
		e.translate(dx, dy)
	}
}

func Transform[E Element](in E, fns ...Transformation) E {
	out := clone(in)
	for _, fn := range fns {
		fn(out)
	}
	return out
}
