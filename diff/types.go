package diff

import (
	"reflect"

	"github.com/krelinga/go-lib/datapath"
	"github.com/krelinga/go-lib/valid"
)

// Result will always be an instance of one of the following types:
// - Extra
// - Missing
// - ValueDiff
// - TypeDiff
type Result interface {
	Path() datapath.Path

	resultIsAClosedType()
}

type Extra struct {
	path datapath.Path
	rhs  any
}

func (e Extra) Path() datapath.Path {
	return e.path
}

func (e Extra) Rhs() any {
	return e.rhs
}

func (Extra) resultIsAClosedType() {}

type Missing struct {
	path datapath.Path
	lhs  any
}

func (m Missing) Path() datapath.Path {
	return m.path
}

func (m Missing) Lhs() any {
	return m.lhs
}

func (Missing) resultIsAClosedType() {}

type ValueDiff struct {
	path datapath.Path
	lhs  any
	rhs  any
}

func (d ValueDiff) Path() datapath.Path {
	return d.path
}

func (d ValueDiff) Lhs() any {
	return d.lhs
}

func (d ValueDiff) Rhs() any {
	return d.rhs
}

func (ValueDiff) resultIsAClosedType() {}

type TypeDiff struct {
	path datapath.Path
	lhs  reflect.Type
	rhs  reflect.Type
}

func (d TypeDiff) Path() datapath.Path {
	return d.path
}

func (d TypeDiff) Lhs() reflect.Type {
	return d.lhs
}

func (d TypeDiff) Rhs() reflect.Type {
	return d.rhs
}

func (TypeDiff) resultIsAClosedType() {}

type Results []Result

func (e Results) Equal() bool {
	return len(e) == 0
}

type AnyDiffer interface {
	supports(reflect.Type) error
	anyDiff(lhs, rhs any) []Result
}

type TypedDiffer[T any] interface {
	AnyDiffer
	valid.Validator

	typedDiff(lhs, rhs T) []Result
}

type DefaultDifferer[T any] interface {
	DefaultDiffer() TypedDiffer[T]
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
