package main

import "fmt"

func main() {

	var N int

	fmt.Scan(&N)

	var myslice []int = make([]int, N, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&myslice[i])
	}

	for idx, val := range myslice {
		if idx%2 == 0 {
			fmt.Print(val, " ")
		}
	}
}
