package main

import "fmt"

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	cur := make([]int, 0)
	helper(n, k, 1, cur, &res)
	return res

}

func helper(n, k, start int, cur []int, res *[][]int) {
	if k == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := start; i <= n-k+1; i++ {
		cur = append(cur, i)
		helper(n, k-1, i+1, cur, res)
		cur = cur[:len(cur)-1]
	}
}

func main() {
	fmt.Println(combine(4, 2))
}
