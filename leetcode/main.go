package main

import (
	
	"fmt"
	"slices"
	//"sort"
)

func main() {

	nums := []int{3,0,6,1,5}
	fmt.Println(hIndex(nums))
	//fmt.Print(len(nums))
}

func hIndex(citations []int) int {
    
    slices.Sort(citations)
    fmt.Print(citations)
    h := len(citations)
  
    fmt.Println(" ")
    for i := range citations {
        
    
        if citations[i] <= len(citations) - (i) {
            h = citations[i]
            
            
        }else if h < len(citations) - (i){
            h = len(citations) - (i)
        }


        fmt.Println("h = ",h , "len - i = " ,len(citations) - (i))
    }

    return h
}