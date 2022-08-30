package main

import "fmt"

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}

func one(xPtr *int) {
	*xPtr = 1
}
