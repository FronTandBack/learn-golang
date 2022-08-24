package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	var sliceOfDigits = getSliceOfDigits(num)
	var result = sumArray(sliceOfDigits...)
	fmt.Println(result)
}

func sumArray(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func getSliceOfDigits(num int) []int {
	var remainder int = 10
	var res int
	var slc []int = []int{}
	for num != 0 {
		res = num % remainder
		num = num / 10
		slc = append(slc, res)
	}

	return slc
}
