package main

import "fmt"

func main() {

	// var a, b int
	var len int
	var counter int
	fmt.Scan(&len)
	var n int
	var total int
	for fmt.Scan(&n); counter < len; fmt.Scan(&n) {

		if numberIsTwoDigit(n) {
			if n%8 == 0 {
				total += n
			}
		}
		counter++
	}

	fmt.Println(total)

}

func numberIsTwoDigit(n int) bool {

	if n > 9 && n < 100 {
		return true
	}
	return false
}
