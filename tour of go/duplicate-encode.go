package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CountChar(word string, c byte) int {

	total := 0
	nc := byte(unicode.ToLower(rune(c)))

	for _, v := range word {

		if unicode.ToLower(v) == rune(nc) {
			total++
		}
	}

	return total
}

func DuplicateEncode(word string) string {

	// dpEn := make([]string, 0)

	var str string = ""

	for i := 0; i < len(word); i++ {
		ch := word[i]
		// var totalChar int = strings.Count(word, ch)

		count := CountChar(word, ch)
		if count > 1 {
			// dpEn = append(dpEn, ")")
			str += ")"
		} else {
			// dpEn = append(dpEn)
			str += "("
		}
		// fmt.Println(CountChar(word, ch))
	}
	return str

}

func DuplicateEncode2(word string) (r string) {
	word = strings.ToLower(word)

	t := map[rune]int{}
	for _, c := range word {
		t[c]++
	}

	for _, c := range word {
		if t[c] == 1 {
			r += "("
		} else {
			r += ")"
		}
	}

	return
}

func main() {
	fmt.Println(DuplicateEncode2("Success"))
}
