package main

import (
	"fmt"
)

var (
	k, set = 0, 0
	once   bool
	ogs    string
)

func longestPalindrome(s string) string {
	if !once {
		ogs = s
		once = true
	}
	flag := 0
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = +1, -1 {
		if r[i] != r[j] {
			flag = 1
			break
		}
	}
	if flag == 0 {
		return s
	}
	fmt.Println(s, "------->")
	// for set+2 != len(ogs) {
	// 	if first {
	// 		fmt.Println("First called")
	// 		set++
	// 		first = false
	// 		return longestPalindrome(string(ogs[set:]))
	// 	} else if second {
	// 		fmt.Println("Second called")
	// 		first = true
	// 		return longestPalindrome(string(ogs[:len(ogs)-set]))
	// 	}
	// }
	for k < len(ogs) {
		for set+2 != len(ogs)-k {
			set++
			return longestPalindrome(ogs[k : len(ogs)-set])
		}
		k++
		set = -1
	}
	return "no palindrome found"
}

func main() {
	str := "cbbd"
	palindrome := longestPalindrome(str)
	fmt.Println("longest palindrome is :", palindrome)
}
