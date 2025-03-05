package ktest

type TestingT interface {
	Helper()
	Errorf(format string, args ...any)
}