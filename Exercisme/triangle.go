// // This is a "stub" file.  It's a little start on your solution.
// // It's not a complete solution though; you have to write some code.

// // Package triangle should have a package comment that summarizes what it's about.
// // https://golang.org/doc/effective_go.html#commentary
// package main

// import "fmt"

// // Notice KindFromSides() returns this type. Pick a suitable data type.
// type Kind string

// const (
// 	NaT Kind = "not a triangle" // not a triangle
// 	Equ Kind = "equilateral"    // equilateral
// 	Iso Kind = "isosceles"      // isosceles
// 	Sca Kind = "scalene"        // scalene
// )

// func KindFromSides(a, b, c float64) Kind {

// 	if a+b+c <= 0 {
// 		return NaT
// 	}

// 	if (a+b < c) || (a+c < b) || (b+c < a) {
// 		return NaT
// 	}

// 	var k Kind

// 	if a == b && b == c {
// 		k = Equ
// 	} else if a != b && b != c && a != c {
// 		k = Sca
// 	} else if (a == b && b != c) || (a == c && c != b) || (b == c && c != a) {
// 		k = Iso
// 	} else {
// 		k = NaT
// 	}

// 	return k
// }

// func main() {
// 	fmt.Println(KindFromSides(2, 2, 2))
// }

package main

import "math"

// Kind is the triangle kind
type Kind int

// Kind constants
const (
	NaT Kind = iota // Not a triangle
	Equ             // Equilateral
	Iso             // Isosceles
	Sca             // Scalene
)

// KindFromSides tells if a triangle is equilateral, isosceles or scalene
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a+b+c) || a+b <= c || a+c <= b || b+c <= a {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
