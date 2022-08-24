package main

import (
	"fmt"
	"strings"
)

type S struct {
	Name string
}

func printUpper(s *S) {
	s.Name = strings.ToUpper(s.Name)
	fmt.Println(s.Name)
}

func main() {
	s := &S{
		Name: "foo",
	}
	printUpper(s)
	fmt.Println(s.Name)
}
