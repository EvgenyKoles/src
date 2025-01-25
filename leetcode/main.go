package main

import (
	//"errors"
	//"crypto/subtle"
	"fmt"
	//"sort"
	//"net"
	//"sort"
)

func main() {

	nums := []int{7,1,5,3,6,4}

	//maxProfit(nums)

	fmt.Print(maxProfit(nums))
	
}




func maxProfit(prices []int) int {
    


	sum := 0

	for i := 0; i < len(prices)-1; i++ {

		if prices[i] < prices[i+1] {
			
			sum += prices[i+1] - prices[i]
		}

	}
	
	return sum
} 