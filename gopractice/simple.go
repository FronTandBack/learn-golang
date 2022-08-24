package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Outside a goruntine.")
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("Inside a goroutine", i)
			time.Sleep(3 * time.Second)
		}(i)
	}

	fmt.Println("Outside again")

	runtime.Gosched()
}
