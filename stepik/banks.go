package main

import "fmt"

func main() {
	var x, p, y uint

	fmt.Scan(&x, &p, &y)

	var money uint
	var years int
	for x <= y {
		money = (x * p / 100)
		x = money + x
		years++
	}
	fmt.Println(years)
}
