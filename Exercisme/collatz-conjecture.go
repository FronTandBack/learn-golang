package main

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {

	if n == 1 {
		return 0, nil
	}

	steps := 0
	for n != 1 {

		if isEven(n) {
			n = n / 2
		} else if isOdd(n) {
			n = 3*n + 1
		}
		steps++
	}

	if steps == 0 {
		return 0, errors.New("No result!")
	}

	return steps, nil
}

func isOdd(n int) bool {

	return !isEven(n)
}

func isEven(n int) bool {

	if n%2 == 0 {
		return true
	}
	return false
}

func main() {
	(CollatzConjecture(10))
}
