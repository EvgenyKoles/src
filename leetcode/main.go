package main

import (
	
	"fmt"
	//"slices"
	//"sort"
)

func main() {

	nums := []int{1,2,3,4}
	fmt.Println(productExceptSelf(nums))
   // fmt.Print(len(nums))
}

func productExceptSelf(nums []int) []int {

    ans := []int{}
    m := nums[0]

    for i := len(nums)-1; i >= 0; i-- {

        m = nums[i] * m
        ans = append(ans, m)
        fmt.Println("m = ", m, "  ","nums[i]= ", nums[i])
    }
    
    return ans
    
}
