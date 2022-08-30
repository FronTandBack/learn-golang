package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str := []rune(strings.Join(strings.Split(text, " "), ""))
	if isPalindrome(str) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")

	}
}

func isPalindrome(str []rune) bool {
	lastIdx := len(str) - 1
	for i := 0; i < lastIdx/2 && i < (lastIdx-i); i++ {
		if str[i] != str[lastIdx-i] {
			return false
		}
	}
	return true
}
