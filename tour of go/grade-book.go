package main

import "fmt"

func GetGrade(a, b, c int) rune {
	// Code here

	score := (a + b + c) / 3

	if 90 <= score && score <= 100 {
		return 'A'
	} else if 80 <= score && score < 90 {
		return 'B'
	} else if 70 <= score && score < 80 {
		return 'C'
	} else if 60 <= score && score < 70 {
		return 'D'
	} else if 0 <= score && score < 60 {
		return 'F'
	}

	return 'O'
}

func GetGrade2(a, b, c int) rune {
	switch mean := (a + b + c) / 3; {
	case mean < 60:
		return 'F'
	case mean < 70:
		return 'D'
	case mean < 80:
		return 'C'
	case mean < 90:
		return 'B'
	default:
		return 'A'
	}

}

func main() {
	fmt.Println(GetGrade(95, 90, 93))
}
