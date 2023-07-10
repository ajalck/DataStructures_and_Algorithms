package main

import "fmt"

func main() {
	nums := []int{512, 867, 904, 997, 403}
	is := canBeIncreasing(nums)
	fmt.Println(is)
}
func canBeIncreasing(nums []int) bool {
	flag := 0
	for i := 1; i < len(nums); i++ {
		fmt.Println(i)
		if nums[i-1] >= nums[i] {
			fmt.Println(flag, i)
			if flag == 0 {
				flag++
				fmt.Println(nums[i])
			} else {
				return false
			}
			fmt.Println("index :",i)
			if i-2 >= 0 && nums[i-2] < nums[i] || i-1 == 0 {
				continue
			} else if i == len(nums) {
				return true

			}else if i+1 <= len(nums) && nums[i-1] < nums[i+1]{
				continue
			} else {
				return false
			}
		}
	}
	return true
}
