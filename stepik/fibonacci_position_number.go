package main

import "fmt"

func main() {

	var num uint64
	var a, b uint64 = 0, 1
	fmt.Scan(&num)
	var numbers []uint64 = []uint64{}
	for i := 1; uint64(i) <= num+1; i++ {
		// var temp int = b
		// b = a + b
		// a = temp
		a, b = b, a+b
		numbers = append(numbers, a)
		if a == num {
			// fmt.Println(i)
			// return
		}

	}
	fmt.Println(numbers)

	var res int = binarySearch(numbers, num)

	if res != -1 {
		fmt.Println(res + 1)
	} else {
		fmt.Println(-1)
	}

	// fmt.Println(-1)
}

func binarySearch(numbers []uint64, searchNum uint64) int {

	var left int = 0
	var right int = len(numbers) - 1
	for left <= right {
		var mid int = left + (right-left)/2

		if numbers[mid] < searchNum {
			left = mid + 1
		} else if numbers[mid] > searchNum {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 0
// 1
// 1
// 2
// 3
// 5
// 8
// 13
