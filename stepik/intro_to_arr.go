package main

import "fmt"

func main() {
	var num uint8
	var workArray [10]uint8

	for idx, _ := range workArray {
		fmt.Scan(&num)
		workArray[idx] = num
	}

	var j, k uint8

	for i := 0; i < 3; i++ {
		fmt.Scan(&j, &k)

		if (j >= 0 && j < 10) && (k >= 0 && k < 10) {
			var temp uint8 = workArray[j]
			workArray[j] = workArray[k]
			workArray[k] = temp
		}
	}
	fmt.Println(workArray)
	for _, item := range workArray {
		fmt.Print(item)
		fmt.Print(" ")
	}
}
