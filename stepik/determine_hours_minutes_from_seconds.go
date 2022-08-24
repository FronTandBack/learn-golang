package main

import "fmt"

func main() {
	var k int

	fmt.Scan(&k)

	var hours = k / 3600
	var minutes = (k % 3600) / 60
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
