package main

import (
	"bufio"
	"fmt"
	"os"
	//"runtime/trace"
	///"strings"
	//	"unicode"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')



	rs := []rune(text)
	// sr := []rune(text)

	max := (len(rs)-2)/2
	//fmt.Print(max)
	var flag bool

	for i, _ := range rs {

		if i > max {
			break			
		}
		if rs[i]==rs[(len(rs)-2)-i] {
			flag = true
			continue
		} else {
			flag = false
			break
		}

	}

	if flag {
		fmt.Print("Палиндром")
	} else {
		fmt.Print("Нет")
	}


}
