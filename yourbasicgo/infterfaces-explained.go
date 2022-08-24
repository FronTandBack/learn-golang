package main

import (
	"fmt"
	"strconv"
)

type MyStringer interface {
	String() string
}

type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

type Point2 struct {
	x, y int
}

func (p Point2) String() string {
	return "My string!"
}

func main() {
	var x MyStringer

	fmt.Println(x == nil) // true

	x = (*Point)(nil)
	fmt.Println(x == nil) // false

	fmt.Printf("%v %T\n", x, x)
	x = Temp(24)
	fmt.Println(x.String()) // 24 °C

	x = &Point{1, 2}
	fmt.Println(x.String()) // (1,2)

	x = Point2{3, 4}
	fmt.Println(x)
	var l interface{}

	l = 2.4
	fmt.Println(l) // 2.4

	l = &Point{1, 2}
	fmt.Println(l) // (1,2)

}
