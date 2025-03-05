package datapath_test

import (
	"testing"

	"github.com/krelinga/go-lib/datapath"
)

func TestDataPath(t *testing.T) {
	tests := []struct {
		name string
		in datapath.Path
		want string
	}{
		{
			name: "Field",
			in: datapath.Path{}.Field("foo"),
			want: "$.foo",
		},
		{
			name: "Index",
			in: datapath.Path{}.Index(42),
			want: "$[42]",
		},
		{
			name: "Key",
			in: datapath.Path{}.Key("bar"),
			want: "$[bar]",
		},
		{
			name: "TypeAssert",
			in: datapath.Path{}.TypeAssert("baz"),
			want: "$.(baz)",
		},
		{
			name: "PtrDeref",
			in: datapath.Path{}.PtrDeref(),
			want: "(*$)",
		},
		{
			name: "Complex",
			in: datapath.Path{}.Field("foo").Index(42).Key("bar").TypeAssert("baz").PtrDeref(),
			want: "(*$.foo[42][bar].(baz))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.String(); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}