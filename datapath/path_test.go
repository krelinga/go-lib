package datapath_test

import (
	"testing"

	"github.com/krelinga/go-lib/datapath"
)

func TestDataPath(t *testing.T) {
	tests := []struct {
		name         string
		in           datapath.Path
		want         string
		wantBasename string
	}{
		{
			name:         "Field",
			in:           datapath.Path{}.Field("foo"),
			want:         "$.foo",
			wantBasename: "base.foo",
		},
		{
			name:         "Index",
			in:           datapath.Path{}.Index(42),
			want:         "$[42]",
			wantBasename: "base[42]",
		},
		{
			name:         "Key",
			in:           datapath.Path{}.Key("bar"),
			want:         "$[bar]",
			wantBasename: "base[bar]",
		},
		{
			name:         "TypeAssert",
			in:           datapath.Path{}.TypeAssert("baz"),
			want:         "$.(baz)",
			wantBasename: "base.(baz)",
		},
		{
			name:         "PtrDeref",
			in:           datapath.Path{}.PtrDeref(),
			want:         "(*$)",
			wantBasename: "(*base)",
		},
		{
			name:         "Method",
			in:           datapath.Path{}.Method("Foo"),
			want:         "$.Foo()",
			wantBasename: "base.Foo()",
		},
		{
			name:         "Complex",
			in:           datapath.Path{}.Method("biff").Field("foo").Index(42).Key("bar").TypeAssert("baz").PtrDeref(),
			want:         "(*$.biff().foo[42][bar].(baz))",
			wantBasename: "(*base.biff().foo[42][bar].(baz))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.String(); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
			if got := tt.in.Basename("base"); got != tt.wantBasename {
				t.Errorf("got %q, want %q", got, tt.wantBasename)
			}
		})
	}
}
