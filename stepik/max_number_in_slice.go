package main

import "fmt"

func main() {
	var N int = 5

	// fmt.Scan(&N)
	var slc []int
	// if N >= 4 {
	slc = make([]int, N, N)
	// }
	for i := 0; i < len(slc); i++ {
		fmt.Scan(&slc[i])
	}
	var max int
	for _, item := range slc {
		if item > max {
			max = item
		}
	}
	fmt.Println(max)
}
