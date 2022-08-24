package main

import (
	"fmt"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

// Contents returns the file's contents as a string.
// func Contents(filename string) (string, error) {
// 	f, err := os.Open(filename)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer f.Close() // f.Close will run when we're finished.

// 	var result []byte
// 	buf := make([]byte, 100)
// 	for {
// 		n, err := f.Read(buf[0:])
// 		result = append(result, buf[0:n]...) // append is discussed later.
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			return "", err // f will be closed if we return here.
// 		}
// 	}
// 	return string(result), nil // f will be closed if we return here.
// }

func main() {
	// fmt.Println(Contents("file.txt"))
	b()
}
