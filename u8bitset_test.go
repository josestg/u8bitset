package u8bitset_test

import (
	"github.com/josestg/u8bitset"
	"strings"
	"testing"
)

func TestBitSet(t *testing.T) {
	set := u8bitset.New()
	equal(t, true, set.Empty())
	equal(t, uint8(0), set.Cardinal())
	equal(t, strings.Repeat("0", 256), set.String())

	set.Add(0)
	set.Add(1)
	set.Add(2)
	set.Add(65)
	set.Add(95)
	set.Add(177)
	set.Add(254)
	set.Add(255)

	equal(t, false, set.Empty())
	equal(t, uint8(8), set.Cardinal())
	equal(t, true, set.Has(0))
	equal(t, true, set.Has(1))
	equal(t, true, set.Has(2))
	equal(t, true, set.Has(65))
	equal(t, true, set.Has(95))
	equal(t, true, set.Has(177))
	equal(t, true, set.Has(254))
	equal(t, true, set.Has(255))
	equal(t, false, set.Has(3))
	equal(t, false, set.Has(253))

	values := []uint8{0, 1, 2, 65, 95, 177, 254, 255}
	sliceEqual(t, values, set.Values())

	set.Del(0)
	set.Del(65)
	set.Del(177)

	values = []uint8{1, 2, 95, 254, 255}
	sliceEqual(t, values, set.Values())

	equal(t, false, set.Empty())
	equal(t, uint8(5), set.Cardinal())
	equal(t, false, set.Has(0))
	equal(t, true, set.Has(1))
	equal(t, true, set.Has(2))
	equal(t, false, set.Has(65))
	equal(t, true, set.Has(95))
	equal(t, false, set.Has(177))
	equal(t, true, set.Has(254))
	equal(t, true, set.Has(255))
	equal(t, false, set.Has(3))
	equal(t, false, set.Has(253))

	set.Reset()
	equal(t, true, set.Empty())
	equal(t, uint8(0), set.Cardinal())
	equal(t, strings.Repeat("0", 256), set.String())
}

func TestBitSet_Intersection(t *testing.T) {
	a := u8bitset.New()
	b := u8bitset.New()

	a.Add(0)
	a.Add(65)
	a.Add(177)
	a.Add(254)

	b.Add(65)
	b.Add(95)
	b.Add(177)
	b.Add(255)

	c := a.Intersection(b)
	d := b.Intersection(a)
	values := []uint8{65, 177}
	sliceEqual(t, values, c.Values())
	sliceEqual(t, values, d.Values())
	sliceEqual(t, c.Values(), d.Values())
}

func TestBitSet_Union(t *testing.T) {
	a := u8bitset.New()
	b := u8bitset.New()

	a.Add(0)
	a.Add(65)
	a.Add(177)
	a.Add(254)

	b.Add(65)
	b.Add(95)
	b.Add(177)
	b.Add(255)

	c := a.Union(b)
	d := b.Union(a)
	values := []uint8{0, 65, 95, 177, 254, 255}
	sliceEqual(t, values, c.Values())
	sliceEqual(t, values, d.Values())
	sliceEqual(t, c.Values(), d.Values())
}

func TestBitSet_Difference(t *testing.T) {
	a := u8bitset.New()
	b := u8bitset.New()

	a.Add(0)
	a.Add(65)
	a.Add(177)
	a.Add(254)

	b.Add(65)
	b.Add(95)
	b.Add(177)
	b.Add(255)

	c := a.Difference(b)
	d := b.Difference(a)
	values := []uint8{0, 254}
	sliceEqual(t, values, c.Values())
	values = []uint8{95, 255}
	sliceEqual(t, values, d.Values())
}

func equal[T comparable](t *testing.T, exp, act T) {
	t.Helper()
	if exp != act {
		t.Errorf("Expected %v, got %v", exp, act)
	}
}

func sliceEqual[T comparable](t *testing.T, exp, act []T) {
	t.Helper()
	if len(exp) != len(act) {
		t.Errorf("Expected %v, got %v", exp, act)
	}
	for i := 0; i < len(exp); i++ {
		if exp[i] != act[i] {
			t.Errorf("Expected %v, got %v", exp, act)
		}
	}
}
