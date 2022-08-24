package main

import "fmt"

func main() {

	var num int32

	fmt.Scan(&num)

	var leftPart int32 = num % 1000
	var rightPart int32 = (num - leftPart) / 1000

	var sumLeft int32 = getSumOfThreeNum(leftPart)
	var sumRight int32 = getSumOfThreeNum(rightPart)

	if sumLeft == sumRight {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func getSumOfThreeNum(n int32) int32 {

	var firstNum int32 = n / 100
	var remainder int32 = n % 100
	var secondNum int32 = remainder / 10
	var thirdNumber int32 = remainder % 10

	return firstNum + secondNum + thirdNumber
}
