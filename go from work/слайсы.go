	for i := 0; i < n; i++{
	}
0
1
2
3
4
5
6
7
8
9

	for i := 1; i <= n; i++{
	}
1
2
3
4
5
6
7
8
9
10

----------------------------------------------------------------------------------------

	var n, numbers, sum int
	fmt.Scan(&n)
	t := []int{}
	
	for i := 0; i < n; i++{
		fmt.Scan(&numbers)
		t = append (t, numbers)
	}
	


	for _, value := range t{
		if value > 9 && value <100{
			if value%8 == 0 {
				sum=value+sum
			}
	
			}
	}
	fmt.Print(sum)	
-------------------------------------------------------------------------------

package main

import (
	"fmt"
)

func main() {

	var m,k int
	t := []int{}
	var n int
	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n){
		t = append (t, n)
	}

	//fmt.Println(t)

	for _, value := range t{
		if m <= value {
			m = value
		} 
	}	

	for _, value2 := range t{
		if value2 == m {
			k++
		}
	}

	//fmt.Println(m)
	fmt.Println(k)



}

--------------------------------------------------------------------------
	// sum := 0
	// sumall :=0
	// for i := 1; i <= n; i++ {
	// 	for j :=1; j <= m; j++{
	// 		sum = i+j
	// 		sumall = sumall +sum
	// 		fmt.Println("---",sum)
	// 	}
	// }


---------------------------------



package main

import (
	"fmt"
)

func main() {

	var numbers int
	t := []int{}
	
	for i := 0; i <= 2; i++{
		fmt.Scan(&numbers)
		t = append (t, numbers)
	}
	
	var year int
	deposit := int(t[0])

		for i:=0; i<100; i++{
		if deposit >= t[2]{
			break
		}else {
			deposit = deposit + deposit*t[1]/100
			year++
			fmt.Println(deposit,"----")
		}
		//fmt.Println(t[0],"--")
	}

	fmt.Println(year)
 
}


--------------------------------------------------------
	
	SBoey30xQr
	5rMoIhM4b0Q9
	
	