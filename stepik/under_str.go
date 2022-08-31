package main

import "fmt"

func main() {

	var str, inStr string

	// fmt.Scan(&str)
	// fmt.Scan(&inStr)
	str = "awesoepme"
	inStr = "ep"
	// strLenght := len(str)
	// inStrLength := len(inStr)
	// for i := 0; i < strLenght-inStrLength; i++ {
	// 	var searchStr = str[i:(i + len(inStr))]
	// 	if searchStr == inStr {
	// 		fmt.Println(i)
	// 		break
	// 	}
	// }

	fmt.Println(subStr(inStr, str))
}

// awesome
// ep
// 0 == 2 false
// counter = s2[0] == s1[0] = false
func subStr(s2 string, s1 string) int {
	var counter int = 0 // pointing s2
	var i int = 0
	for i < len(s1) {
		if counter == len(s2) {
			break
		}
		if s2[counter] == s1[i] {
			counter++
		} else {
			// Special case where character preceding
			// the i'th character is duplicate
			if counter > 0 {
				i -= counter
			}
			counter = 0
		}
		i++
	}

	if counter < len(s2) {
		return -1
	}
	return i - counter

}
