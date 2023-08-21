# u8bitset

A simple implementation of BitSet for `uint8` type in Go.


## Usage

```go
package main

import (
	"fmt"
	"github.com/josestg/u8bitset"
)

func main() {
	set := u8bitset.New()
	str := "Hello, World!"
	for _, c := range str {
		set.Add(uint8(c))
	}

	other := "World!"
	otherSet := u8bitset.New()
	for _, c := range other {
		otherSet.Add(uint8(c))
	}

	fmt.Println(set.Has(uint8('H'))) // true
	fmt.Println(set.Has(uint8('h'))) // false
	fmt.Println(set.Cardinal())      // 10

	intersection := set.Intersection(otherSet).Values()
	fmt.Printf("Intersection: %s\n", intersection) // Intersection: !Wdorl

}
```
