package main

import "fmt"

// func generate(numRows int) [][]int {

// 	triangle := make([][]int, numRows)
// 	for i := 0; i < numRows; i++ {
// 		for j := 1; j < i; j++ {
// 			tr := make([]int, 0)
// 			tr = append(tr, i+j)
// 			triangle[i-1] = append(triangle[i], tr...)
// 		}
// 	}

// 	fmt.Println(triangle)

// 	return [][]int{}
// }

func main() {
	var rows int
	var temp int = 1
	fmt.Print("Enter number of rows : ")
	// fmt.Scan(&rows)
	rows = 5
	outer := make([][]int, 0)
	inner := make([]int, 0)
	for i := 0; i < rows; i++ {

		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k <= i; k++ {

			if k == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}

			// fmt.Printf(" %d", temp)
			inner = append(inner, temp)
		}
		outer = append(outer, inner)
		inner = inner[:len(inner)-1]

		fmt.Println("")

	}
}
