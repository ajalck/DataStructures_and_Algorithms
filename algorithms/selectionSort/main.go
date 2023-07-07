/*
SELECTION SORT
>> Selecting smallest element from unsorted part and putting it in the i th position in ascending order
>> Time complexity 
		best case - O(n)
		worst case - O(n^2)
*/

package main

import "fmt"

func main() {

	nums := []int{43, 65, -2, 24, 65, 7, 23, 47, 147, 61}

	for i := 0; i < len(nums)-1; i++ {
		min := i
		fmt.Printf("For %d position finding min value at--->",i)	
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		fmt.Print(min,"\n")
		nums[i], nums[min] = nums[min], nums[i]
	}
	fmt.Println("sorted array is :",nums)
}
