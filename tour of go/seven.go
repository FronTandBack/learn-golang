package main

import (
	"fmt"
	"strconv"
)

func Seven(n int64) []int {

	steps := 0
	for true {
		o := n % 10
		n = n - o
		n = n / 10

		sa := n
		n = n - (2 * o)
		if n < 0 {
			n = sa
			break
		}
		l := strconv.FormatInt(n, 10)
		steps++
		if len(l) == 2 && n%7 == 0 {
			break
		} else if len(l) == 2 {
			break
		} else if len(l) == 1 && n%7 == 0 {
			break
		} else if n <= 0 {
			break
		}
		// fmt.Println(n)

	}

	// fmt.Println(steps, n)
	// var arr []int = [steps, n]
	// var arr[]int, int64 = {steps, n}
	data := []int{int(n), steps}

	return data
}

func Seven2(n int64) []int {
	r := 0
	for ; n >= 100; n = n/10 - 2*(n%10) {
		r += 1
	}
	return []int{int(n), r}
}

func main() {
	fmt.Println(Seven(132498))
	fmt.Println(Seven2(132498))
}
