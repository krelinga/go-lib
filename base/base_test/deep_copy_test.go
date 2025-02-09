package base_test

import (
	"testing"

	"github.com/krelinga/go-lib/base"
	"github.com/stretchr/testify/assert"
)

type testDeepCopierPtr struct {
	a, b int
	c    *testDeepCopierChildPtr
}

func (tdc *testDeepCopierPtr) DeepCopy() interface{} {
	return &testDeepCopierPtr{
		a: tdc.a,
		b: tdc.b,
		c: base.DeepCopy(tdc.c),
	}
}

type testDeepCopierChildPtr struct {
	d int
}

func (tdcc *testDeepCopierChildPtr) DeepCopy() interface{} {
	return &testDeepCopierChildPtr{
		d: tdcc.d,
	}
}

type testDeepCopierValue struct {
	a, b int
	c    testDeepCopierChildValue
}

func (tdc testDeepCopierValue) DeepCopy() interface{} {
	return testDeepCopierValue{
		a: tdc.a,
		b: tdc.b,
		c: base.DeepCopy(tdc.c),
	}
}

type testDeepCopierChildValue struct {
	d int
}

func (tdcc testDeepCopierChildValue) DeepCopy() interface{} {
	return testDeepCopierChildValue{
		d: tdcc.d,
	}
}

func TestDeepCopy(t *testing.T) {
	t.Run("Ptr", func(t *testing.T) {
		inParent := &testDeepCopierPtr{
			a: 1,
			b: 2,
			c: &testDeepCopierChildPtr{
				d: 3,
			},
		}
		outParent := base.DeepCopy(inParent)
		inParent.a = 4
		inParent.b = 5
		inParent.c.d = 6

		assert.Equal(t, 1, outParent.a)
		assert.Equal(t, 2, outParent.b)
		assert.Equal(t, 3, outParent.c.d)
		assert.NotSame(t, inParent, outParent)
		assert.NotSame(t, inParent.c, outParent.c)
	})

	t.Run("Value", func(t *testing.T) {
		inParent := testDeepCopierValue{
			a: 1,
			b: 2,
			c: testDeepCopierChildValue{
				d: 3,
			},
		}
		outParent := base.DeepCopy(inParent)
		inParent.a = 4
		inParent.b = 5
		inParent.c.d = 6

		assert.Equal(t, 1, outParent.a)
		assert.Equal(t, 2, outParent.b)
		assert.Equal(t, 3, outParent.c.d)
	})
}
