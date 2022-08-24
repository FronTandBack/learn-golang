package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func ConvertFracts(a [][]int) string {
	d := 1
	for _, num := range a {
		g := gcd(num[0], num[1])
		num[0] /= g
		num[1] /= g
		d = d * num[1] / gcd(d, num[1])
	}
	res := ""
	for _, num := range a {
		res += fmt.Sprintf("(%d,%d)", num[0]*d/num[1], d)
	}
	return res
}

// 3563
// 2 * 3 * 4 = 3 * 2 * 2 = 12

// 130 1310  40 = (13 * 2 * 5) (131 * 2 * 5) (2 * 2 * 2 * 5)

func main() {
	var lst = [][]int{{69, 130}, {87, 1310}, {30, 40}}

	// gcd(18, 12)
	ConvertFracts(lst)
	// ConvertFracts(lst)
}
