package main

import "fmt"

func main() {
	arr := []int{1, 12, 23, 43, 65, 65, 67, 87, 98, 102}
	fmt.Println("Enter the search element :")
	var element int
	fmt.Scanf("%d", &element)
	is, index := binarySearch(arr, element)
	fmt.Printf("Value is present %v index is %d", is, index)
}
func binarySearch(arr []int, value int) (bool, int) {
	high := len(arr)
	low := 1
	for low <= high {
		mid := (high + low) / 2
		if value == arr[mid] {
			return true, mid
		} else if value < arr[mid] {
			high = mid - 1
		} else if value > arr[mid] {
			low = mid + 1
		}
	}
	return false, -1
}
