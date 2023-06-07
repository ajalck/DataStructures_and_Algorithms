package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	arr := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}
	newArr := plusOne(arr)
	fmt.Println("new array is ", newArr)
}
func plusOne(digits []int) []int {
	var digistr string
	for i := range digits {
		digistr += strconv.Itoa(digits[i])
	}
	value := new(big.Int)
	value, ok := value.SetString(digistr, 10)
	if !ok {
		fmt.Println("error occured")
	}
	value.Add(value, big.NewInt(1))
	var newarr []int
	newval := value.String()
	for i := range newval {
		j, _ := strconv.Atoi(string(newval[i]))
		newarr = append(newarr, j)
	}
	return newarr
}
