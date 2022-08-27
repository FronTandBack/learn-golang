package main

import (
	"fmt"
)

func main() {
	// var N int

	// fmt.Scan(&N)
	// var res int = 1
	// for i := 0; ; i++ {
	// 	res = int(math.Pow(2, float64(i)))
	// 	if res > N {
	// 		break
	// 	}
	// 	fmt.Printf("%d ", res)
	// }

	for i := 1; i <= 50; i *= 2 {
		fmt.Print(i, " ")
	}
}

// 1
// 2
// 2 * 2
// 2 * 2 * 2
// 2 * 2 * 2 * 2
