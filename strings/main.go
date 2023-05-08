package main

import (
	//"bufio"
	"fmt"
	//"unicode"
	//"os"
	"strings"
	//"runtime/trace"
	//"strings"
	//	"unicode"
)

func main() {

// 	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

// 	rs := []rune(text)
	

	


//  if string(rs[len(rs)-1]) == "." && unicode.IsUpper(rs[0]) {
// 	 fmt.Print("Right")
	
// 	 }else {
// 	 	fmt.Print("Wrong")}

 //--------------------------------------
	// text, _ := bufio.NewReader(os.Stdin).ReadString('\n')



	// rs := []rune(text)
	// // sr := []rune(text)

	// max := (len(rs)-2)/2
	// //fmt.Print(max)
	// var flag bool

	// for i, _ := range rs {

	// 	if i > max {
	// 		break			
	// 	}
	// 	if rs[i]==rs[(len(rs)-2)-i] {
	// 		flag = true
	// 		continue
	// 	} else {
	// 		flag = false
	// 		break
	// 	}

	// }

	// if flag {
	// 	fmt.Print("Палиндром")
	// } else {
	// 	fmt.Print("Нет")
	// }
	//-------------------------------------
	// var test string
	// fmt.Scan(&test)

	// var v string
	// fmt.Scan(&v)

	// //fmt.Println(test, v)
	// fmt.Print(strings.Index(test,v))


    //-------------------------------------
	// var test string
	// fmt.Scan(&test)
	// test2 := strings.Split(test, "")
 	// test3:= []string{}

	// for i, _ := range test2 {
	// 	if i%2!=0 {
	// 		test3= append(test3,test2[i])
	// 		continue
	// 	}
	// }

	// for i, _ := range test3{
	// 	fmt.Print(test3[i])
	// }

	//--------------------------------

	// var test string
	// fmt.Scan(&test)
	// test2 := strings.Split(test, "")
 	// test3:= []string{}
	
	

	// 	for i := 0; i < len(test2); i++ {
	// 		if strings.Count(test, test2[i])<2 {
	// 			//fmt.Println("+")
	// 			test3 = append(test3, test2[i])
	// 		}
	// 	}


	// 	// fmt.Println(test2)
	// 	// fmt.Print(test3)	

	// 	for i, _ :=range test3{
	// 		fmt.Print(test3[i])
	// 	}

//------------------------


	var test string
	fmt.Scan(&test)
	//test2 := strings.Split(test, "")
 	//test3:= []string{}
	

	if len(test)<5 || !checkAlphabet(test) {
		fmt.Print("Wrong password")
	} else {
		fmt.Print("Ok")
	}


}
	func checkAlphabet (t string) bool {
		alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
		flag := true

		for i, _ := range t {
			if !strings.Contains(alphabet,string(t[i])) {
				flag = false
				break
			}
		}
		if !flag {
			return false
		}else {
			return true
		}
	}







