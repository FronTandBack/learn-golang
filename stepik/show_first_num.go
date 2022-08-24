package main

import "fmt"

func main() {
	var num int

	// fmt.Scan(&num)
	num = 9546
	var remainder int = 10
	var res int
	for num != 0 {
		res = num % remainder
		num = num / 10
	}

	fmt.Println(res)
}
