package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsUpper(s rune) bool {
	if !unicode.IsUpper(s) && unicode.IsLetter(s) {
		return false
	}

	return true
}

func IsLower(s rune) bool {

	if !unicode.IsLower(s) && unicode.IsLetter(s) {
		return false
	}

	return true
}

func solve(str string) string {

	upper := 0
	lower := 0
	for _, v := range str {

		if IsUpper(v) {
			upper++
		}

		if IsLower(v) {
			lower++
		}
	}

	if upper > lower {
		str = strings.ToUpper(str)
	} else {
		str = strings.ToLower(str)
	}

	return str
}

func main() {

	fmt.Println(solve("coooDDDDDDDDDDe"))
}
