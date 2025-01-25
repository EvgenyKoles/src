package main

import (
	"fmt"
)

func main() {

	nums := []int{1,1,1,0}
	
	//nums := []int{1,1,1,0}
	//maxProfit(nums)

	fmt.Print(canJump(nums))
	
}




func canJump(nums []int) bool {
	max :=0
	k := len(nums)-1;
	
	// if len(nums)-1 == 0{
	// 	return true
	// }
	// if nums[0] == 1 && nums[1]==0{
	// 	return true
	// }
	// if nums[0] == 0 {
	// 	return false
	// }
	
	for i:= 0; i < len(nums); i++ {
		if nums[i]> max{
			max = nums[i]+1
		}
		fmt.Println("k= ", k, ",nums[i] = ", nums[i], "max = ", max, "i= ", i)
		if nums[i] >= k {
			return true
		}
		
		if nums[i] == 0 && (nums[i-1]==1 || nums[i-1]==0) && max < 2{
			return false
		}
		if max == 0 && nums[i] == 0{
			return false
		}
		k--
		max--
	}
    
	return false
}