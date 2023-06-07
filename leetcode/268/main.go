package main

import "fmt"

func main() {
	arr := []int{3, 0, 1}
	output := missingNumber(arr)
	fmt.Println(output)
}
func missingNumber(nums []int) int {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] == j {
			nums[i], nums[j] = nums[j], nums[i]
			i = j
			j++
			continue
		}
	}
	return j
}
