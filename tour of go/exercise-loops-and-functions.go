package main

func Sqrt(x float64) float64 {
	z := float64(1)

	for z != x {

		z -= (z*z - x) / (2 * z)
	}

	return z

}

func main() {
	Sqrt(2)
}
