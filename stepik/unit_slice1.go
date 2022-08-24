package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)
	var slc []int
	if N >= 4 {
		slc = make([]int, N, N)
	}
	for i := 0; i < len(slc); i++ {
		fmt.Scan(&slc[i])
	}

	fmt.Println(slc[3])

}
