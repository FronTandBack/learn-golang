package main

import "fmt"

// 1: "корова"

// от 2 до 4: "коровы"

// от 5 до 20: "коров"

// от 21: если число заканчивается на "1", то "корова"

// если число заканчивается на 2, 3, 4, то "коровы"

func main() {
	var k int

	fmt.Scan(&k)

	if k == 1 {
		fmt.Printf("%d korova", k)
	} else if k >= 2 && k <= 4 {
		fmt.Printf("%d korovy", k)
	} else if k >= 5 && k <= 20 {
		fmt.Printf("%d korov", k)
	} else if k >= 21 {
		var endN int = k % 10
		if endN == 1 {
			fmt.Printf("%d korova", k)
		} else if endN == 2 || endN == 3 || endN == 4 {
			fmt.Printf("%d korovy", k)
		} else {
			fmt.Printf("%d korov", k)
		}
	}
}
