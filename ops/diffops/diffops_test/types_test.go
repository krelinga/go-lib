package diffops_test

import "github.com/krelinga/go-lib/ops/diffops"

type Foo struct {
	IsString     string
	IsInt        int
	IsMap        map[int]int
	IsMapMap     map[int]map[int]int
	IsSlice      []int
	IsSliceSlice [][]int
	IsPtrBar     *Bar
	IsBar        Bar
	IsSliceBar   []Bar
}

func (f *Foo) DiffOp(rhs any, s diffops.Sink) {
	diffops.CastRhs(f, rhs, s, func(rhs *Foo) {
		diffops.ByEqual[string]()(f.IsString, rhs.IsString, s.Field("IsString"))
		diffops.ByEqual[int]()(f.IsInt, rhs.IsInt, s.Field("IsInt"))
		diffops.MapOf[int](diffops.ByEqual[int]())(f.IsMap, rhs.IsMap, s.Field("IsMap"))
		diffops.MapOf[int](diffops.MapOf[int](diffops.ByEqual[int]()))(f.IsMapMap, rhs.IsMapMap, s.Field("IsMapMap"))
		diffops.SliceOf(diffops.ByEqual[int]())(f.IsSlice, rhs.IsSlice, s.Field("IsSlice"))
		diffops.SliceOf(diffops.SliceOf(diffops.ByEqual[int]()))(f.IsSliceSlice, rhs.IsSliceSlice, s.Field("IsSliceSlice"))
		diffops.ByMethod[*Bar]()(f.IsPtrBar, rhs.IsPtrBar, s.Field("IsPtrBar"))
		diffops.ByMethod[Bar]()(f.IsBar, rhs.IsBar, s.Field("IsBar"))
		diffops.SliceOf(diffops.ByMethod[Bar]())(f.IsSliceBar, rhs.IsSliceBar, s.Field("IsSliceBar"))
	})
}

type Bar struct {
	IsDouble float64
}

func (b Bar) DiffOp(rhs any, s diffops.Sink) {
	diffops.CastRhs(b, rhs, s, func(rhs Bar) {
		diffops.EqualsPlan(b.IsDouble, rhs.IsDouble, diffops.ByEqual[float64]())
	})
}
