package main

import (
	"fmt"
	"math"
)

func IsNumber(n int) bool {

	powers := []int{}
	oldValue := n
	for n != 0 {
		mod := n % 10
		powers = append(powers, mod)
		n /= 10
	}

	res := 0.0
	for i := 0; i < len(powers); i++ {

		res += math.Pow(float64(powers[i]), float64(len(powers)))
	}

	return int(res) == oldValue
}

func main() {
	fmt.Println(IsNumber(153))
}
