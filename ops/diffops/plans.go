package diffops

import "reflect"

func isSafe(in any) bool {
	v := reflect.ValueOf(in)
	if !v.IsValid() {
		return false
	}
	if v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		return !v.IsNil()
	}
	return true
}

func bothSafe(lhs, rhs any, s Sink) bool {
	lhsOk := isSafe(lhs)
	rhsOk := isSafe(rhs)
	if !lhsOk && !rhsOk {
		return false
	}
	if lhsOk != rhsOk {
		s.ValueDiff(lhs, rhs)
		return false
	}
	return true
}

func ByMethod[T DiffOper]() Plan[T] {
	return func(lhs, rhs T, s Sink) {
		if bothSafe(lhs, rhs, s) {
			lhs.DiffOp(rhs, s)
		}
	}
}

func ByEqual[T comparable]() Plan[T] {
	return func(lhs, rhs T, s Sink) {
		CastRhs(lhs, rhs, s, func(rhs T) {
			if lhs != rhs {
				s.ValueDiff(lhs, rhs)
			}
		})
	}
}

func Deref[T any](in Plan[T]) Plan[*T] {
	return func(lhs, rhs *T, s Sink) {
		if bothSafe(lhs, rhs, s) {
			// TODO: start here.  We can't deref RHS because we don't actually know that its a pointer.
			// Should I consider changing the Plan type to always require that LHS and RHS are the same type?
			in(*lhs, *rhs, s)
		}
	}
}

func SliceOf[T any](valueFn Plan[T]) Plan[[]T] {
	return func(lhs, rhs  []T, s Sink) {
		CastRhs(lhs, rhs, s, func(rhs []T) {
			for i := range lhs {
				if !s.WantMore() {
					return
				}
				if i >= len(rhs) {
					s.Key(i).Missing(lhs[i])
					continue
				}
				valueFn(lhs[i], rhs[i], s.Key(i))
			}
	
			for i := len(lhs); i < len(rhs) && s.WantMore(); i++ {
				s.Extra(rhs[i])
			}
		})
	}
}

func MapOf[K comparable, V any](valueFn Plan[V]) Plan[map[K]V] {
	return func(lhs, rhs map[K]V, s Sink) {
		CastRhs(lhs, rhs, s, func(rhs map[K]V) {
			for k, v := range lhs {
				if !s.WantMore() {
					return
				}
				rhsV, ok := rhs[k]
				if !ok {
					s.Key(k).Missing(v)
					continue
				}
				valueFn(v, rhsV, s.Key(k))
			}
	
			for k, v := range rhs {
				if !s.WantMore() {
					return
				}
				if _, ok := lhs[k]; !ok {
					s.Extra(v)
				}
			}
		})
	}
}

