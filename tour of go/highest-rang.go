package main

import "fmt"

func InsertionSort(arr []int) []int {

	var i int = 1

	for i < len(arr) {
		var j int = i

		for j >= 1 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]

			j--
		}
		i++
	}

	return arr
}

func HighestRank2(nums []int) int {
	mii, maxK, maxV := map[int]int{}, 0, 0
	for _, v := range nums {
		mii[v]++
		if mii[v] > maxV || (mii[v] == maxV && v > maxK) {
			maxK = v
			maxV = mii[v]
		}
	}
	return maxK
}

func HighestRank(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	nums = InsertionSort(nums)
	count := 0
	max := 0
	var freaq int
	// var m map[int]int
	for i := 0; i < len(nums)-1; i++ {

		if nums[i] == nums[i+1] {
			count++
		} else {
			count = 0
		}

		if count >= max {
			max = count

			freaq = nums[i+1]

		}
	}
	fmt.Println(freaq)
	return freaq
}

func main() {
	var nums []int = []int{2, 2, 2, 6, 7, 19, 10, 10, 10, 30, 45, 43, 56, 6, 66, 8, 8, 8, 8, 8}

	HighestRank(nums)
	HighestRank2(nums)
}
