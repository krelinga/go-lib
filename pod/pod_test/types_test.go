package pod_test

import "github.com/krelinga/go-lib/pod"

type Outer struct {
	String   string
	Int      int
	Map      map[string]int
	Inner    *Inner
	MapInner map[string]*Inner
}

func (o *Outer) InternalDiff(rhs pod.POD, reporter pod.DiffReporter) {
	typedRhs, ok := rhs.(*Outer)
	if !ok {
		reporter.TypeDiff(o, rhs)
		return
	}
	if o.String != typedRhs.String {
		reporter.ChildField("String").ValueDiff(o.String, typedRhs.String)
	}
	if o.Int != typedRhs.Int {
		reporter.ChildField("Int").ValueDiff(o.Int, typedRhs.Int)
	}
	pod.Diff(pod.WrapMapComp(o.Map), pod.WrapMapComp(typedRhs.Map), reporter.ChildField("Map"))
	pod.Diff(o.Inner, typedRhs.Inner, reporter.ChildField("Inner"))
	pod.Diff(pod.WrapMapPOD(o.MapInner), pod.WrapMapPOD(typedRhs.MapInner), reporter.ChildField("MapInner"))
}

func (o *Outer) InternalDeepCopyTo(out pod.POD) {
	// TODO: Implement
}

func (o *Outer) InternalValidate(reporter pod.ErrorReporter) {
	// TODO: implement
}

type Inner struct {
	String string
}

func (i *Inner) InternalDiff(rhs pod.POD, reporter pod.DiffReporter) {
	typedRhs, ok := rhs.(*Inner)
	if !ok {
		reporter.TypeDiff(i, rhs)
		return
	}
	if i.String != typedRhs.String {
		reporter.ChildField("String").ValueDiff(i.String, typedRhs.String)
	}
}

func (i *Inner) InternalDeepCopyTo(out pod.POD) {
	// TODO: Implement
}

func (i *Inner) InternalValidate(reporter pod.ErrorReporter) {
	// TODO: implement
}
