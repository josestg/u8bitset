package field_test

import (
	"github.com/josestg/u8bitset/internal/field"
	"math"
	"strings"
	"testing"
)

func TestField_SetBit(t *testing.T) {
	f := field.Field(0)
	n := uint64(0)
	assertEqual(t, field.Field(n), f)

	f = f.SetBit(0)
	n += 1 << 0
	assertEqual(t, field.Field(n), f)

	f = f.SetBit(1)
	n += 1 << 1
	assertEqual(t, field.Field(n), f)

	f = f.SetBit(50)
	n += 1 << 50
	assertEqual(t, field.Field(n), f)
}

func TestField_DelBit(t *testing.T) {
	f := field.Field(math.MaxUint64) // all bits set.
	n := uint64(math.MaxUint64)
	assertEqual(t, field.Field(n), f)

	for i := uint8(0); i < 64; i++ {
		f = f.DelBit(i)
		n -= 1 << i
		assertEqual(t, field.Field(n), f)
	}

	assertEqual(t, field.Field(0), f)
}

func TestField_IsSet(t *testing.T) {
	f := field.Field(0)
	assertFalse(t, f.IsSet(0))
	f = field.Field(1<<0 + 1<<1 + 1<<16)
	assertTrue(t, f.IsSet(0))
	assertTrue(t, f.IsSet(1))
	assertFalse(t, f.IsSet(2))
	assertFalse(t, f.IsSet(15))
	assertTrue(t, f.IsSet(16))
	assertFalse(t, f.IsSet(17))
}

func TestField_Empty(t *testing.T) {
	f := field.Field(0)
	assertTrue(t, f.Empty())
	f = field.Field(50)
	assertFalse(t, f.Empty())
}

func TestField_Cardinal(t *testing.T) {
	f := field.Field(0)
	assertTrue(t, f.Cardinal() == 0)

	f = field.Field(1<<0 + 1<<1 + 1<<16)
	assertTrue(t, f.Cardinal() == 3)

	f = field.Field(math.MaxUint64)
	assertTrue(t, f.Cardinal() == 64)
}

func TestField_String(t *testing.T) {
	f := field.Field(0)
	s := strings.Repeat("0", 64)
	assertTrue(t, f.String() == s)

	f = field.Field(math.MaxUint64)
	s = strings.Repeat("1", 64)
	assertTrue(t, f.String() == s)
}

func TestDifference(t *testing.T) {
	a := field.Field(0)
	a = a.SetBit(0)
	a = a.SetBit(1)
	a = a.SetBit(2)
	a = a.SetBit(3)

	b := field.Field(0)
	b = b.SetBit(0)
	b = b.SetBit(2)
	b = b.SetBit(4)

	c := field.Difference(a, b)
	assertFalse(t, c.IsSet(0))
	assertTrue(t, c.IsSet(1))
	assertFalse(t, c.IsSet(2))
	assertTrue(t, c.IsSet(3))

	d := field.Difference(b, a)
	assertFalse(t, d.IsSet(0))
	assertFalse(t, d.IsSet(1))
	assertFalse(t, d.IsSet(2))
	assertFalse(t, d.IsSet(3))
	assertTrue(t, d.IsSet(4))

	e := field.Difference(a, a)
	assertTrue(t, e.Empty())

	f := field.Difference(b, b)
	assertTrue(t, f.Empty())
}

func assertEqual(t *testing.T, exp, act field.Field) {
	t.Helper()
	if exp != act {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}

func assertTrue(t *testing.T, act bool) {
	t.Helper()
	if !act {
		t.Errorf("Expected true, got false")
	}
}

func assertFalse(t *testing.T, act bool) {
	t.Helper()
	if act {
		t.Errorf("Expected false, got true")
	}
}
