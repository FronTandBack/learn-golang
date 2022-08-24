package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)

	var sliceOfDigits = getSliceOfDigits(num)
	var str string = ""
	for _, val := range sliceOfDigits {
		str += strconv.Itoa(val)
	}

	fmt.Println(str)
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

// Better solution
// var num, reversed int
// fmt.Scan(&num)

// for num > 0 {
// 	reversed  = reversed * 10 + num % 10
// 	num /= 10
// }

// fmt.Println(reversed)
