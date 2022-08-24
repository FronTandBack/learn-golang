package main

import "fmt"

func main() {

	var n int
	var max int
	var counter int = 1
	var prev int
	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		if n >= max {
			prev = max
			max = n
			if prev == max {
				counter = counter + 1
			} else {
				counter = 1
			}
		}
	}
	fmt.Println(counter)
}
