package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {

	var m map[string]int

	m = make(map[string]int)

	for _, lA := range listArt {

		// var str []string = strings.Split(s, " ")
		// fmt.Println(str)
		pos := bytes.IndexByte([]byte(lA), byte(' '))
		num, err := strconv.Atoi(lA[pos+1:])

		ch := string(lA[0])
		if err == nil {
			m[ch] += num
		}

	}

	var res string
	var sum int = 0
	for _, lC := range listCat {
		res += "(" + lC + " : " + strconv.Itoa(m[lC]) + ") - "
		sum += m[lC]
	}

	// for _, val := range m {
	// 	sum += val
	// }

	if sum == 0 {
		return ""
	}
	res = res[:len(res)-3]
	return res
}

func StockList2(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	quantity := make(map[string]int)
	for _, stock := range listArt {
		splitted := strings.Split(stock, " ")
		firstCh := string(splitted[0][0])
		n, _ := strconv.Atoi(splitted[1])
		quantity[firstCh] += n
	}

	result := make([]string, len(listCat))
	for i, code := range listCat {
		result[i] = fmt.Sprintf("(%s : %v)", code, quantity[code])
	}

	return strings.Join(result, " - ")
}

func main() {
	var b = []string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"F", "G", "L", "K", "Z"}

	fmt.Println(StockList(b, c))
}
