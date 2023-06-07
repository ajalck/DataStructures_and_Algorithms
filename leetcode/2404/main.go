package main

import "fmt"

func main() {
	arr := []int{8154, 9139, 8194, 3346, 5450, 9190, 133, 8239, 4606, 8671, 8412, 6290}
	element := mostFrequentEven(arr)
	fmt.Println("element is :", element)
}
func mostFrequentEven(nums []int) int {
	numMap := make(map[int]int)
	for _, key := range nums {
		if key%2 == 0 {
			numMap[key]++
		}
	}
	//fmt.Println(numMap)
	var sk, lgfz int = -1, -1
	for key, fz := range numMap {
		if fz > lgfz {
			lgfz = fz
			sk = key
			//fmt.Printf("lgfz for key %d is %d\n", key, lgfz)
		} else if fz == lgfz {
			//fmt.Println("reached")
			if sk > key {
				sk = key
			}
		} 
	}
	if lgfz != -1 {
		return sk
	} else {
		return -1
	}
}
