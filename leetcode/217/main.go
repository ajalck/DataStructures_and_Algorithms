package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 1}
	is := containsDuplicate(arr)
	fmt.Println(is)
}

func containsDuplicate(nums []int) bool {
	for j := 0; j < len(nums)-1; j++ {
		for i := j + 1; i < len(nums); i++ {

			if nums[j] == nums[i] {
				return true
			}
		}
	}
	return false
}
