package main

import "fmt"

type H struct {
	id *int
}

func one(xPtr *int) {
	*xPtr = 153
}

func main() {

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

	a := 10

	p := &a
	fmt.Println(p)  //0x10410020
	fmt.Println(*p) //10

	a = 11
	fmt.Println(p) //0x10410020
	fmt.Println(*p)

	q := 90
	b := H{id: &q}
	fmt.Println(b.id)  //0x1041002c
	fmt.Println(*b.id) //90

	q = 80
	fmt.Println(b.id)  //0x1041002c
	fmt.Println(*b.id) //80
}
