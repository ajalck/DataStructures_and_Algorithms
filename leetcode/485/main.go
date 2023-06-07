package main

import "fmt"

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	max := findMaxConsecutiveOnes(nums)
	fmt.Println("max is :", max)
}
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count += 1
		} else {
			if max < count {
				max = count
			}
			count = 0
		}
	}
	if max < count {
		max = count
	}
	return max
}
