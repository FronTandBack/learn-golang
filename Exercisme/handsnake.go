package main

import "fmt"

func Handshake(code uint) []string {

	// var secreHandsnake = map[int]string{
	// 	1:     "wink",
	// 	10:    "double blink",
	// 	100:   "close your eyes",
	// 	1000:  "jump",
	// 	10000: "Reverse the order of the operations in the secret handshake.",
	// }

	var binary []uint

	for code != 0 {
		binary = append(binary, code%2)
		code = code / 2
	}

	fmt.Println(binary)

	var test = []string{"my str"}
	return test
}

func Handshake2(handshake uint) []string {
	code := []string{}
	if handshake&1 == 1 {
		code = append(code, "wink")
	}
	if handshake&2 == 2 {
		code = append(code, "double blink")
	}
	if handshake&4 == 4 {
		code = append(code, "close your eyes")
	}
	if handshake&8 == 8 {
		code = append(code, "jump")
	}
	if handshake&16 == 16 {
		code = reverse(code)
	}
	return code
}
func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func main() {
	fmt.Println(Handshake2(19))
}
