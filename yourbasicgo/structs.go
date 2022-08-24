package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	var a Student
	a.Name = "Alice"

	var pa *Student   // pa == nil
	pa = new(Student) // pa == &Student{"", 0}
	pa.Name = "Alice" // pa == &Student{"Alice", 0}

	b := Student{ // b == Student{"Bob", 0}
		Name: "Bob",
	}

	pb := &Student{ // pb == &Student{"Bob", 8}
		Name: "Bob",
		Age:  8,
	}

	c := Student{"Cecilia", 5} // c == Student{"Cecilia", 5}
	d := Student{}

	fmt.Println(b, pb, c, d)

}
