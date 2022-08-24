package main

import (
	"fmt"
)

func FindNb(m int) int {
	// your code

	for n := 1; m > 0; n++ {
		m -= n * n * n
		if m == 0 {
			return n
		}
	}
	return -1

}

func main() {

	fmt.Println(FindNb(1071225))

}
