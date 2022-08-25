package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	var max int = 0

	if a < 0 && b < 0 {
		max = a
	}
	var flag bool = false
	for i := b; i >= a; i-- {
		if i%7 == 0 {
			if i >= max {
				max = i
				flag = true
			}
		}

	}

	if max%7 == 0 && flag {
		fmt.Println(max)
	} else {
		fmt.Println("NO")
	}
}
