package datapath

import "fmt"

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

func (p Path) Index(i int) Path {
	return Path{
		str: fmt.Sprintf("%s[%d]", p, i),
	}
}

func (p Path) Key(k any) Path {
	return Path{
		str: fmt.Sprintf("%s[%v]", p, k),
	}
}

func (p Path) TypeAssert(name string) Path {
	return Path{
		str: fmt.Sprintf("%s.(%s)", p, name),
	}
}

func (p Path) PtrDeref() Path {
	return Path{
		str: fmt.Sprintf("(*%s)", p),
	}
}
