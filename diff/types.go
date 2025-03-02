package diff

import (
	"reflect"

	"github.com/krelinga/go-lib/valid"
)

type Path string

// Entry will always be an instance of one of the following types:
// - Extra
// - Missing
// - Different
type Entry interface {
	Path() Path
}

type Extra struct {
	path Path
	rhs  any
}

func (e Extra) Path() Path {
	return e.path
}

func (e Extra) Rhs() any {
	return e.rhs
}

type Missing struct {
	path Path
	lhs  any
}

func (m Missing) Path() Path {
	return m.path
}

func (m Missing) Lhs() any {
	return m.lhs
}

type Different struct {
	path Path
	lhs  any
	rhs  any
}

func (d Different) Path() Path {
	return d.path
}

func (d Different) Lhs() any {
	return d.lhs
}

func (d Different) Rhs() any {
	return d.rhs
}

type AnyDiffer interface {
	valid.Validator

	diffType() reflect.Type
	anyDiff(lhs, rhs any) []Entry
}

type TypedDiffer[T any] interface {
	AnyDiffer

	typedDiff(lhs, rhs T) []Entry
}

// Start here:
// - make AnyDiffer.diffType() more of a query method: pass in a type and get a bool back
// - we need a default differ for the Any type, that's probably a better way to start this.
// - it may be that we need remove TypedDiffer altogether, and just have AnyDiffer, especially for the default case.
// - change Struct.AllExportedFields to be Struct.Explicit, and have it default to false.
//   That way the default behavior can be to diff all exported fields, which is probably what we want.
// - consider adding a Kind() method to Different, so that we can make it easier for callers to check for
//   differences based on type vs. different values of the same type?
// - Or maybe that is better off as a different Entry implementation?
// - Consider refining Path ... I suspect it might be easier for some callers to have the individual components vs. the whole thing.
// - Consider adding an interface to allow types to define their default differ?
// - If I do the above, consider adding a Diff() method at the top of this package that takes two values of the same type, randomly
//   select the default differ for either value, and does the diff.  Randomness is important because it will help us catch cases
//   where the default differ is not symmetric.
// - If I can make generic differs work, then I can do things like:
//   - have Default as a var, which diffs any two entries with some reasonable defaults.
//   - have Skip as a var, which just skips the diff and always returns no differences.
//   - have PtrEqual as a var, which ensures that two pointers point to the same address.
