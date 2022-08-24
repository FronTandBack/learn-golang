package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(Score("cabbage"))

}

func Score(word string) int {
	var scrabble = map[string]int{
		"A, E, I, O, U, L, N, R, S, T": 1,
		"D, G":                         2,
		"B, C, M, P":                   3,
		"F, H, V, W, Y":                4,
		"K":                            5,
		"J, X":                         8,
		"Q, Z":                         10,
	}

	var total int = 0
	for _, ch := range word {
		for letters, _ := range scrabble {
			var idx int = strings.Index(letters, strings.ToUpper(string(ch)))
			if idx > -1 {
				total += int(scrabble[letters])
			}
		}
	}

	return total

}
