package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, my name is John"
	seg := countSegments(s)
	fmt.Println("no of segments is :", seg)
}
func countSegments(s string) int {
	strs := strings.Split(s, " ")
	count := 0
	for _, str := range strs {
		if str != "" {
			count += 1
		}
	}
	return count
}
