package main

import (
	"fmt"
	"unicode"
)

func main() {

	fmt.Println(IsIsogram("Alphabet"))
}

func IsIsogram(word string) bool {

	var isoletter map[rune]bool

	isoletter = make(map[rune]bool)

	for _, w := range word {

		if (w < 'a' || w > 'z') && (w < 'A' || w > 'Z') {
			continue
		}
		w = unicode.ToLower(w)
		if !isoletter[w] {
			isoletter[w] = true
		} else {
			return false
		}
	}
	return true
}
