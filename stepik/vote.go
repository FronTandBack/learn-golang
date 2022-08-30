package main

import "fmt"

func main() {
	fmt.Println(vote(1, 0, 0, 0))
}

func vote(x int, y int, z int, d int) int {

	// 0 0 0 = 0
	// 0 0 1 = 0
	// 1 1 1
	// 1 1 0
	//

	return (x + y + z + d) / 2

}
