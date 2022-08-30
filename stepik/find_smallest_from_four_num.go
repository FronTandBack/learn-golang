package main

import "fmt"

func main() {
	minimumFromFour()
}

func minimumFromFour() int {
	var slc []int = make([]int, 4)

	for i := 0; i < 4; i++ {

		fmt.Scan(&slc[i])
	}
	var i = 1
	for i < len(slc) {
		var j = i
		for j >= 1 && slc[j] < slc[j-1] {
			slc[j], slc[j-1] = slc[j-1], slc[j]

			j--
		}

		i++
	}

	return slc[0]
}
