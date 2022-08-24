package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) string {
	fmt.Println(str)

	res := strings.Split(str, " ")

	// newstr := ""
	for _, word := range res {
		// newstr += " " + strings.ToUpper(word[0])
		fmt.Println(word[0])
	}
	return str
}

func main() {

	ToJadenCase("How can mirrors be real if our eyes aren't real")
}
