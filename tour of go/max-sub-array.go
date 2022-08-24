package main

import (
	"fmt"
)

func sum(numbers []int) int {

	var sum int = 0

	for _, v := range numbers {

		sum += v
	}

	return sum

}

func MaximumSubarraySum(numbers []int) int {

	// maxSubArr := []int{}
	if len(numbers) == 1 {
		return numbers[0]
	}
	max := 0
	// var start, end int
	for i := 1; i < len(numbers); i++ {

		for j := 0; j < i; j++ {
			if sum(numbers[j:i]) > max {
				max = sum(numbers[j:i])
				// start = j
				// end = i
			}
		}
	}

	// maxSubArr = numbers[start:end]

	fmt.Println(max)
	// fmt.Println(maxSubArr)
	return max
}

func main() {

	MaximumSubarraySum2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaximumSubarraySum2(numbers []int) int {
	res, sum := 0, 0
	for i := range numbers {
		sum += numbers[i]
		res = max(res, sum)
		sum = max(sum, 0)
	}
	return res
}
