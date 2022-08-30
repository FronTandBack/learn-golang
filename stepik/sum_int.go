package main

import "fmt"

func main() {
	fmt.Println(sumInt(1, 2, 3, 4, 6))
}

func sumInt(a ...int) (int, int) {

	var total int
	for _, t := range a {
		total = total + t
	}
	return len(a), total
}
