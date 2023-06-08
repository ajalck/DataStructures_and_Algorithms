package main

import "fmt"

func main() {
	nums := []int{1, 2, 34, 3, 4, 5, 7, 23, 12}
	is := threeConsecutiveOdds(nums)
	fmt.Println(is)
}
func threeConsecutiveOdds(arr []int) bool {
	flag := 0
	for _, v := range arr {
		if v%2 != 0 {
			flag++
			fmt.Println(flag, v)
		} else {
			flag = 0
		}
		if flag == 3 {
			return true
		}

	}
	return false
}
