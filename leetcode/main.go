package main

import (
	//"errors"
	"fmt"
	//"sort"
	//"net"
	//"sort"
)

func main() {

	nums := []int{3,2,6,5,0,3}

	//maxProfit(nums)

	fmt.Print(maxProfit(nums))

}

func maxProfit(prices []int) int {
    
	max := 0
	min := prices[0]
	k := 0

	for i := 0; i < len(prices); i++ {

		if min > prices[i] && i < len(prices)-1{
			min = prices[i]
			max = prices[i]
		}
		if max < prices[i]{
			max = prices[i]

			if max - min > k {
				k = max - min
			}
		}
		//fmt.Println("max = ", max, " min = ", min, "k= ", k)
	}
	
	return k
}