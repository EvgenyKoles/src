package main

import (
	//"bufio"
	"fmt"
	//"strconv"
	//"os"
)

type Stringer interface {
	String() string
}

type batteryForTest struct {
	Name string
}

func (c batteryForTest) String() string {

	rs := []byte(c.Name)
	var a = []byte{}
	var count = 0

	for _, value := range rs {
		if string(value) == "0" {
			a = append(a, ' ')
			count++
		}
	}

	for i := 0; i < 10-count; i++ {
		a = append(a, 'X')
	}
	//return string(a)
	//fmt.Print("[",string(a),"]")
	return fmt.Sprint("[", string(a), "]")
}

func main() {

	var text string
	fmt.Scan(&text)

	batteryForTest := batteryForTest{
		Name: text,
	}

	fmt.Println(batteryForTest)

}
