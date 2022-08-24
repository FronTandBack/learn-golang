package main

import (
	"fmt"
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {

	m := make(map[string]string)
	str := s1 + s2

	for _, e := range str {

		m[string(e)] = string(e)

	}

	res := []string{}
	for _, v := range m {
		res = append(res, v)
	}

	sort.Strings(res)

	// joining the string by separator
	s := strings.Join(res, "")

	return s
}

func main() {
	fmt.Println(TwoToOne("aretheyhere", "yestheyarehere"))
}
