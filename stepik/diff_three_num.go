package main

import "fmt"

func main() {
	// a != b b != c a != c
	var num int

	fmt.Scan(&num)

	var firstNum int = num / 100
	var remainder int = num % 100
	var secondNum int = remainder / 10
	var thirdNumber int = remainder % 10

	if (firstNum != secondNum) && (secondNum != thirdNumber) && (firstNum != thirdNumber) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
