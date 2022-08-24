package main

import "fmt"

func main() {
	var s []int // a nil slice
	s1 := []string{"foo", "bar"}
	s2 := make([]int, 2)    // same as []int{0, 0}
	s3 := make([]int, 2, 4) // same as new([4]int)[:2]

	fmt.Println(s, s1, s2)
	fmt.Println(len(s3), cap(s3)) // 2 4
}
