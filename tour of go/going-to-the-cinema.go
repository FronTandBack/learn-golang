package main

// import "fmt"

import (
	"fmt"
	"math"
)

func Movie(card, ticket int, perc float64) (count int) {
	var systemA float64
	systemB := float64(card)
	for systemA <= math.Ceil(systemB) {
		count++
		systemA = float64(ticket * count)
		systemB += float64(ticket) * math.Pow(perc, float64(count))
	}
	return count
}
func Movie2(card, ticket int, perc float64) int {

	var t float64 = float64(ticket)
	var n int
	var systemB float64 = float64(card)
	var systemA float64
	//partly my solution based on ranked implementation
	for systemA <= math.Ceil(systemB) {
		n++
		systemA = float64(ticket * n)
		t = t * perc
		systemB = systemB + t

	}

	return int(n)

}

func main() {
	fmt.Println(Movie2(0, 10, 0.95))
}
