package main


func main() {
	type steps int 
	arr := []int{1, 2, 3, 1}
	var k steps=3
	rotate(arr,int(k))
}

func rotate(nums []int, k int)  {
    for i:=0;i<k;i++{
        for j:=0;j<len(nums)-1;j++{
        nums[0],nums[j+1]=nums[j+1],nums[0]
        }
    }
}
