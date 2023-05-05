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
	

	fmt.Print(rs)




}
