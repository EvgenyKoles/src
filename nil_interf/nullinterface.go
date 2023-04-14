package main

import "fmt"

func main() {


// a := "sklsfkksflsf"
// b := "sadfkkasjfkjfaks"
// fmt.Println(a, b)


// var name string //console scan
// var age int
// var name2 string
// fmt.Print("Inter your name: ")
// fmt.Scan(&name, &name2)
// fmt.Print("Inter your age: ")
// fmt.Scan(&age)
// fmt.Print(name,name2, age)


// var a string
// fmt.Scan(&a)
// fmt.Println(string(a[len(a)-2]))


var a int
var h int
var m int
fmt.Scan(&a)

h = a/30
m = (a%30)*2

fmt.Print("It is ", h, " hours ", m, " minutes.")


}