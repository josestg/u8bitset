package field

import "fmt"

// Field is the underlying storage that used to store the membership status each
// element in the set. A single Field can store the membership status of 64
// elements.
type Field uint64

// Empty returns true if all bits in the Field are 0.
func (f Field) Empty() bool {
	return f == 0
}

// SetBit sets the bit at pos in the Field to 1.
func (f Field) SetBit(pos uint8) Field {
	return Union(f, 1<<pos)
}

// DelBit sets the bit at pos in the Field to 0.
func (f Field) DelBit(pos uint8) Field {
	mask := Invert(1 << pos)
	return Intersection(f, mask)
}

// IsSet returns true if the bit at pos in the Field is 1.
func (f Field) IsSet(pos uint8) bool {
	return Intersection(f, 1<<pos) != 0
}

// String returns a string representation of the Field.
func (f Field) String() string {
	return fmt.Sprintf("%064b", f)
}

// Cardinal returns the number of bits that are set to 1 in the Field.
func (f Field) Cardinal() uint8 {
	var count uint8
	for i := uint8(0); i < 64; i++ {
		if f.IsSet(i) {
			count++
		}
	}
	return count
}

// Invert returns the inverse of a Field.
func Invert(f Field) Field { return ^f }

// Union returns the union of two Fields.
func Union(a, b Field) Field { return a | b }

// Intersection returns the intersection of two Fields.
func Intersection(a, b Field) Field { return a & b }

// Difference returns a new field that all bits that are set in A, but not set in B (A / B).
// Note that A/B != B/A.
func Difference(a, b Field) Field {
	ab := Intersection(a, b)
	return a ^ ab
}
