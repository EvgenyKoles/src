package main

import (
	"fmt"
	//"time"
)

func main() {

	m := make (chan string, 5)
	
	task2(m,"8")

	
}

func task2 (c chan string, line string) {

	for i := 0; i < 5; i++ {
	c <- line+ " "
			
	}
	
}

