package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Valid(id string) bool {

	if len(id) <= 1 {
		return false
	}

	nums := strings.Split(id, " ")
	digits := strings.Join(nums[:], "")

	sum := 0
	for i, n := range digits {
		n, _ := strconv.Atoi(string(n))

		if i%2 == len(id)%2 {
			if n = n * 2; n > 9 {
				n -= 9
			}
		}

		sum += n

	}

	if sum%10 == 0 {
		return true
	}
	return false
}

//  0123456789101112131415
// "23232005776 6 3 5 5 4"

func Valid2(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	length := len(input)
	if length <= 1 {
		return false
	}
	sum := 0
	var digit int
	for i, r := range input {
		if !unicode.IsDigit(r) {
			return false
		}
		digit, _ = strconv.Atoi(string(r))
		if i%2 == length%2 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	if sum%10 != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(Valid("059"))
}
