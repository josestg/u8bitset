package u8bitset

import (
	"github.com/josestg/u8bitset/internal/field"
	"strings"
)

// BitSet implementation for uint8.
type BitSet [4]field.Field // 256 bits.

// New creates a empty BitSet.
func New() *BitSet {
	return &BitSet{0, 0, 0, 0}
}

// Empty returns true if the set is empty.
func (bs *BitSet) Empty() bool {
	for i := 0; i < len(bs); i++ {
		if bs[i] != 0 {
			return false
		}
	}
	return true
}

// Cardinal returns the number of elements in the set.
func (bs *BitSet) Cardinal() uint8 {
	var count uint8
	for i := 0; i < len(bs); i++ {
		count += bs[i].Cardinal()
	}
	return count
}

// Add adds the given value to the set.
func (bs *BitSet) Add(val uint8) {
	f, p := bs.location(val)
	bs[f] = bs[f].SetBit(p)
}

// Has returns true if the given value is in the set.
func (bs *BitSet) Has(val uint8) bool {
	f, p := bs.location(val)
	return bs[f].IsSet(p)
}

// Del removes the given value from the set.
func (bs *BitSet) Del(val uint8) {
	f, p := bs.location(val)
	bs[f] = bs[f].DelBit(p)
}

// Reset resets the set to empty state.
func (bs *BitSet) Reset() {
	for i := 0; i < len(bs); i++ {
		bs[i] = 0
	}
}

// Union returns a new set that is the union of the two sets.
func (bs *BitSet) Union(other *BitSet) *BitSet {
	var union BitSet
	for i := 0; i < len(bs); i++ {
		union[i] = field.Union(bs[i], other[i])
	}
	return &union
}

// Intersection returns a new set that is the intersection of the two sets.
func (bs *BitSet) Intersection(other *BitSet) *BitSet {
	var both BitSet
	for i := 0; i < len(bs); i++ {
		both[i] = field.Intersection(bs[i], other[i])
	}
	return &both
}

// Difference returns a new set that exists in the current set, but not in the other set.
func (bs *BitSet) Difference(other *BitSet) *BitSet {
	var diff BitSet
	for i := 0; i < len(bs); i++ {
		diff[i] = field.Difference(bs[i], other[i])
	}
	return &diff
}

// Values returns a slice of all the values in the set.
func (bs *BitSet) Values() []uint8 {
	values := make([]uint8, 0, bs.Cardinal())
	for i := 0; i < len(bs); i++ {
		for j := uint8(0); j < 64; j++ {
			if bs[i].IsSet(j) {
				values = append(values, uint8(i*64)+j)
			}
		}
	}
	return values
}

// String returns a string representation of the set.
func (bs *BitSet) String() string {
	buf := make([]string, 0, 4)
	for i := 0; i < len(bs); i++ {
		buf = append(buf, bs[i].String())
	}
	return strings.Join(buf, "")
}

func (bs *BitSet) location(val uint8) (field, pos uint8) {
	field = val / 64
	pos = val % 64
	return
}
