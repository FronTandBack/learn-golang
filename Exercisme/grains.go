package main

import (
	"errors"
	"fmt"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("n must be between 1 and 64")
	}
	arr := make([]int, 0)
	arr = append(arr, 1)
	arr = append(arr, 2)
	for i := 2; i < 64; i++ {

		var d int = 2 * arr[i-1]
		arr = append(arr, d)
	}
	fmt.Println(arr)
	return (1 << uint(n-1)), nil
}

// Total returns the number of grains of rice on the entire mythical
// chessboard. Coded as a static value because making it a loop is silly when
// the formula is (2^n)-1 for the partial sum of the first n powers of 2.
func Total() uint64 {
	return 1<<64 - 1
}

func main() {

	// var test uint64 = 1<<64 - 1
	fmt.Println(Square(5))
}
