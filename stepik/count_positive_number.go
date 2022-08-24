package main

import "fmt"

func main() {

	var sizeOfSlice int32

	fmt.Scan(&sizeOfSlice)
	var slc []int = make([]int, sizeOfSlice, sizeOfSlice)
	for i := 0; i < int(sizeOfSlice); i++ {
		fmt.Scan(&slc[i])
	}

	var counter int
	for _, val := range slc {

		if val >= 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
