package main

import "fmt"

func main() {
	arr := []int{-2, 0, 10, -19, 4, 6, -8}
	is := checkIfExist(arr)
	fmt.Println(is)
}
func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if i == j {
				continue
			}
			if arr[i] == (2*arr[j]) || arr[j] == (2*arr[i]) {
				fmt.Println(i, j)
				return true
			}
		}
	}
	return false
}
