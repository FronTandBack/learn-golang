package main

import "fmt"

func main() {
	fmt.Println(fib(4))
}

func fib(n int) int {

	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
