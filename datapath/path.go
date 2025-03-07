package datapath

import (
	"fmt"
	"strings"
)

// Not making Path a string type because I might want to
// extend it in the future in ways that are incompatible with raw strings.
type Path struct {
	str string
}

func (p Path) String() string {
	if p.str == "" {
		return "$"
	}
	return p.str
}

func (p Path) Field(name string) Path {
	return Path{
		str: fmt.Sprintf("%s.%s", p, name),
	}
}

func Field(name string) Path {
	return Path{}.Field(name)
}

func (p Path) Index(i int) Path {
	return Path{
		str: fmt.Sprintf("%s[%d]", p, i),
	}
}

func Index(i int) Path {
	return Path{}.Index(i)
}

func (p Path) Key(k any) Path {
	return Path{
		str: fmt.Sprintf("%s[%v]", p, k),
	}
}

func Key(k any) Path {
	return Path{}.Key(k)
}

func (p Path) TypeAssert(name string) Path {
	return Path{
		str: fmt.Sprintf("%s.(%s)", p, name),
	}
}

func TypeAssert(name string) Path {
	return Path{}.TypeAssert(name)
}

func (p Path) PtrDeref() Path {
	return Path{
		str: fmt.Sprintf("(*%s)", p),
	}
}

func PtrDeref() Path {
	return Path{}.PtrDeref()
}

func (p Path) Basename(bn string) string {
	return strings.Replace(p.String(), "$", bn, 1)
}

func (p Path) Method(name string) Path {
	return Path{
		str: fmt.Sprintf("%s.%s()", p, name),
	}
}

func Method(name string) Path {
	return Path{}.Method(name)
}
