package main

import (
	"fmt"
)

func SquareSum(numbers []int) int {

	var l int = len(numbers)

	sum := 0
	for i := 0; i < l; i++ {

		square := numbers[i] * numbers[i]
		sum += square
	}

	return sum
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	res := SquareSum(arr)

	fmt.Println(res)
}
