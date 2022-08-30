package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}

	s1 = append(s1, s2...)
	fmt.Println(s1)
}
