package main

import (
	//"errors"
	"fmt"
	"sort"
	//"math"
)

	func main() {
		

	nums1 := []int{0}
	m := 0
	nums2 := []int{1} 
	//n := 1


	
	

	if m == 0{
		copy(nums1,nums2)
	}

	nums_part := nums1[:m]

    merged := append(nums_part,nums2... )	
	sort.Ints(merged)

	copy(nums1,merged)

	fmt.Print(nums1)

		
	
}