// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var a, b int
// 	fmt.Scan(&a, &b)
// 	var temp int = b
// 	var sum string
// 	for a != 0 {
// 		var resA int = getNumbers(&a)
// 		for b != 0 {
// 			var resB int = getNumbers(&b)
// 			if resA == resB {
// 				sum = sum + strconv.Itoa(resA)
// 				break
// 			}
// 		}
// 		b = temp
// 	}
// 	fmt.Println(sum)

// 	for i := len(sum) - 1; i >= 0; i-- {
// 		fmt.Print(string(sum[i]))
// 		fmt.Print(" ")
// 	}
// }

// func getNumbers(num *int) int {
// 	var remainder int = 10
// 	var res int
// 	res = *num % remainder
// 	*num = *num / 10

// 	return res

// }

package main

import "fmt"

func main() {
	var a, b int
	a = 564
	b = 8954
	// fmt.Scan(&a, &b)
	for i := 10000; i > 0; i /= 10 {
		if a/i != 0 {
			for k := 10000; k > 0; k /= 10 {
				if b/k != 0 {
					if a/i%10 == b/k%10 {
						fmt.Print(a/i%10, " ")
					}
				}
			}
		}
	}
}
