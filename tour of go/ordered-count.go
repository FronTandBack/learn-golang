package main

import (
	"fmt"
	"strings"
)

type Tuple struct {
	Char  rune
	Count int
}

func contains(t []Tuple, s rune) bool {
	for _, item := range t {
		if item.Char == s {
			return true
		}
	}
	return false
}

func OrderedCount(text string) []Tuple {

	t := []Tuple{}
	for _, v := range text {
		n := strings.Count(text, string(v))

		if !contains(t, v) {
			t = append(t, Tuple{Char: v, Count: n})
		}

	}

	fmt.Println(t)

	return t
}

func main() {
	str := "a long string with many repeated characters"
	// numberOfa := strings.Count(str, "a")
	OrderedCount(str)
	// fmt.Printf("[%v] string has %d of characters of [a] ", str, numberOfa)

}
