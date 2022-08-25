package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)
	if a < b && a < c && b < c {

		var res = math.Sqrt((a * a) + (b * b))

		if res == c {
			fmt.Println("Прямоугольный")
		} else {
			fmt.Println("Непрямоугольный")
		}
	}
}
