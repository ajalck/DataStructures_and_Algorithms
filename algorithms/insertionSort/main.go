/*
INSERTION SORT
>> Inserting elements to its corresponding places
>> Swapping swapping to left upto its corresponding position comes
>> Time complexity 
		best case - O(n)
		worst case - O(n^2)
*/
package main

import "fmt"

func main() {
	nums := []int{43, 65, -2, 24, 65, 7, 23, 47, 147, 61}
	for i := 1; i < len(nums); i++ {
		fmt.Printf("%d th element, %d --->",i,nums[i])
		j := i
		for j > 0 && nums[j] < nums[j-1] {
			fmt.Print(j,",")
			nums[j], nums[j-1] = nums[j-1], nums[j]
			j--
		}
		fmt.Println()
	}
	fmt.Println(nums)
}
