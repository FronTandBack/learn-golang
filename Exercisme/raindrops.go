package main

import (
	"fmt"
	"strconv"
)

// has 3 as a factor, add 'Pling' to the result.
// has 5 as a factor, add 'Plang' to the result.
// has 7 as a factor, add 'Plong' to the result.

func main() {
	fmt.Println(Convert(34))
}

func Convert(number int) string {

	var res string = ""

	if number%3 == 0 {
		res += "Pling"
	}

	if number%5 == 0 {
		res += "Plang"
	}

	if number%7 == 0 {
		res += "Plong"
	}

	if res != "" {
		return res
	} else {
		return strconv.Itoa(number)
	}
}
