package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
	"unicode"
	
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	rs := []rune(text)
	

	fmt.Println(string(rs[len(rs)-2])== ".")
	//fmt.Println(rs[utf8.RuneCountInString(text)-1])

	//strings.HasSuffix(text, string(rs[len(rs)-3]))
	

	fmt.Println(unicode.IsUpper(rs[0]))


	 if string(rs[len(rs)-2]) == "." && unicode.IsUpper(rs[0]) {
	 fmt.Print("Right")
	
	 }else {
	 	fmt.Print("Wrong")}



}
