package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	output := twoSum(arr, target)
	fmt.Println(output)
}
func twoSum(nums []int, target int) []int {
	var i, j int
	j = 0
	for i = j + 1; i < len(nums); i++ {
		if nums[j]+nums[i] == target {
			break
		}
		if i == len(nums)-1 {
			j++
			i = j
		}
	}
	indices := []int{j, i}
	return indices
}
