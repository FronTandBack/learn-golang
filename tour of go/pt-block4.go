package main

import (
	"fmt"
	"strings"
)

type S struct {
	T *T
}

type T struct {
	Name string
}

func printUpper(s S) {
	s.T.Name = strings.ToUpper(s.T.Name)
	fmt.Println(s.T.Name)
}

func modA(a []int) {

	a = append(a, 6)
	fmt.Println(a)
}
func main() {
	s := S{
		T: &T{
			Name: "foo",
		},
	}

	k := []int{3, 5, 6}
	modA(k)
	fmt.Println(k)
	printUpper(s)
	fmt.Println(s.T.Name)
}
