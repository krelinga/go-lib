package diff

import (
	"errors"
	"reflect"
	"testing"
)

type fooer interface {
	Foo() string
}

type myInt int

func (myInt) Foo() string { return "" }

type Bar[T any] struct {
	Val T
}

type myPtrInt int

func (*myPtrInt) Foo() string { return "" }

func TestOptionsDb(t *testing.T) {
	zeroOptions := &options{}
	tests := []struct {
		name        string
		init        func(*optionsDb) error
		t           reflect.Type
		opts        []Option
		want        error
		wantLookups map[reflect.Type]*options
	}{
		{
			name: "builtin type",
			t:    reflect.TypeFor[int](),
			want: errEmptyPkgPath,
			wantLookups: map[reflect.Type]*options{
				reflect.TypeFor[int](): nil,
			},
		},
		{
			name: "anonymous type",
			t:    reflect.TypeOf(struct{ int }{}),
			want: errEmptyPkgPath,
			wantLookups: map[reflect.Type]*options{
				reflect.TypeOf(struct{ int }{}): nil,
			},
		},
		{
			name: "already registered",
			init: func(db *optionsDb) error {
				return db.register(reflect.TypeFor[myInt]())
			},
			t:    reflect.TypeFor[myInt](),
			want: errAlreadyRegistered,
			wantLookups: map[reflect.Type]*options{
				reflect.TypeFor[myInt](): zeroOptions,
			},
		},
		{
			name: "invalid method",
			t:    reflect.TypeFor[myInt](),
			opts: []Option{WithMethods("Bar")},
			want: errInvalidMethod,
			wantLookups: map[reflect.Type]*options{
				reflect.TypeFor[myInt](): nil,
			},
		},
		{
			name: "valid method interface",
			t:    reflect.TypeFor[fooer](),
			opts: []Option{
				WithMethods("Foo"),
			},
			want: nil,
			wantLookups: map[reflect.Type]*options{
				reflect.TypeFor[fooer](): {
					methods: []string{"Foo"},
				},
			},
		},
		{
			name: "valid method non-interface",
			t:    reflect.TypeFor[myInt](),
			opts: []Option{
				WithMethods("Foo"),
			},
			want: nil,
			wantLookups: map[reflect.Type]*options{
				reflect.TypeFor[myInt](): {
					methods: []string{"Foo"},
				},
			},
		},
		{
			name: "generic already registered",
			init: func(db *optionsDb) error {
				return db.register(reflect.TypeFor[Bar[int]]())
			},
			t:    reflect.TypeFor[Bar[string]](),
			want: errAlreadyRegistered,
			wantLookups: map[reflect.Type]*options{
				reflect.TypeFor[Bar[int]]():    zeroOptions,
				reflect.TypeFor[Bar[string]](): zeroOptions,
			},
		},
		{
			name: "method with pointer reciever",
			t:    reflect.TypeFor[*myPtrInt](),
			opts: []Option{WithMethods("Foo")},
			want: nil,
			wantLookups: map[reflect.Type]*options{
				reflect.TypeFor[*myPtrInt](): {
					methods: []string{"Foo"},
				},
			},
		},
		{
			name: "pointer to pointer",
			t:    reflect.TypeFor[**myPtrInt](),
			want: errPtrToPtr,
			wantLookups: map[reflect.Type]*options{
				reflect.TypeFor[**myPtrInt](): nil,
			},
		},
		{
			name: "pointer to interface",
			t:    reflect.TypeFor[*fooer](),
			want: errPtrToInterface,
			wantLookups: map[reflect.Type]*options{
				reflect.TypeFor[*fooer](): nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := make(optionsDb)
			if tt.init != nil {
				if err := tt.init(&db); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
			if err := db.register(tt.t, tt.opts...); !errors.Is(err, tt.want) {
				t.Errorf("optionsDb.register() = %v, want %v", err, tt.want)
			}
			for typ, want := range tt.wantLookups {
				if got := db.lookup(typ); !reflect.DeepEqual(got, want) {
					t.Errorf("optionsDb.lookup(%v) = %v, want %v", typ, got, want)
				}
			}
		})
	}
}
