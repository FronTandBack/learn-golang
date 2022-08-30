package main

import (
	"fmt"
	"strconv"
)

func main() {

	var mynum int
	fmt.Scan(&mynum)

	fmt.Println(strconv.FormatInt(int64(mynum), 2))
}
