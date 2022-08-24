package main

import "fmt"

func main() {

	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	if a < b {
		var sum int
		for i := a; i <= b; i++ {
			sum = sum + i
		}

		fmt.Println(sum)
	}
}
