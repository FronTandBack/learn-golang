package main

import "strings"

func RepeatStr(repititions int, value string) string {
  return strings.Repeat(value, repititions)
}

func main() {

	RepeatStr(4, "a");
}

