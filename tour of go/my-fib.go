package main

import "fmt"

func fibonacci() func() int {
	a := 0
	b := 1
	c := 0
	return func() int {

		c = a + b
		a = b
		b = c

		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
