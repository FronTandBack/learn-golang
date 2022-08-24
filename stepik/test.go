package main

import (
	"fmt"
)

func main() {
	var d int
	fmt.Scan(&d)
	var hours int = d / 30
	var minutes int = 2 * (d % 30)

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
