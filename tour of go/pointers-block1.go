package main

import "fmt"

type S struct {
	Name string
}

func main() {
	var s *S
	var t = S{Name: "Zomb"}
	var x *S

	x = &S{Name: "Foo"}
	s = &S{Name: "Viktor"}
	fmt.Println(s)
	s = &t
	fmt.Println(s)
	fmt.Println(x)
}
