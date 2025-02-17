package pod

import "reflect"

// TODO: rename everything this file to copy.go, and rename DeepCopy to Copy.
// TODO: ditch DeepCopyTo, and just have DeepCopy return a new copy.

func DeepCopy[P POD](in P) P {
	out := reflect.Zero(reflect.TypeOf(in)).Interface().(P)
	in.InternalDeepCopyTo(out)
	return out
}

func DeepCopyTo[P POD](in P, out P) {
	in.InternalDeepCopyTo(out)
}
