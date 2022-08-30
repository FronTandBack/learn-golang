package main

import "fmt"

func main() {
	var a, b int
	// fmt.Scan(&a, &b)
	a = 100
	b = 500
	b = b / 7 * 7
	if b >= a {
		fmt.Print(b)
	} else {
		fmt.Print("NO")
	}
}
