package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	var a [4]int                      // array with zero values
	var b [4]int = [4]int{0, 1, 2}    // partially initialized array
	var c [4]int = [4]int{1, 2, 3, 4} // array initialization
	d := [...]int{5, 6, 7, 0}         // ... - means that array size equals the number of elements in the array literal
	u := make([]int, 4, 5)            // slice of size 4 initialized with zero-valued array of size 5

	fmt.Printf("a: length: %d, capacity: %d, data: %v\n", len(a), cap(a), a)
	fmt.Printf("b: length: %d, capacity: %d, data: %v\n", len(b), cap(b), b)
	fmt.Printf("c: length: %d, capacity: %d, data: %v\n", len(c), cap(c), c)
	fmt.Printf("d: length: %d, capacity: %d, data: %v\n", len(d), cap(d), d)
	fmt.Printf("u: length: %d, capacity: %d, data: %v\n", len(u), cap(u), u)
}
