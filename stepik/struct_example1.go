package main

import (
	"fmt"
	"math"
)

type Vector struct {
	x float64
	y float64
	z float64
}

func createVector(x float64, y float64, z float64) Vector {
	return Vector{x, y, z}
}

func length(v Vector) int {
	return int(math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z))
}

func main() {
	var a Vector = createVector(6, 3, 2)
	var b Vector = createVector(1, 2, 4)
	fmt.Print(length(a), " ")
	fmt.Print(length(b))
}
