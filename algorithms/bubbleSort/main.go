/*
BUBBLE SORT
>> largest element comes last or smallest comes first
>> Swapping adjacent elements by comparing each other
>> Time complexity 
		best case - O(n)
		worst case - O(n^2)
*/
package main

import "fmt"

func main() {
	nums := []int{43, 65,-2, 24, 65, 7, 23, 47, 147, 61}
	count:=0
	for{
		fmt.Print(count,"-->")
		for j := 0; j < len(nums)-1-count; j++ {
			fmt.Print(j,",")
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		fmt.Println()
		count++
		if count==len(nums)-2{
			break
		}
	}
	fmt.Println("sorted array is ",nums)
}
