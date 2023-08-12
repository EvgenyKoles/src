package main

import "fmt"

type animal interface {
    breathe()
    walk()
	speak()
    
}

type lion struct {
    age int
}

func (l lion) breathe() {
    fmt.Println("Lion breathes")
}

func (l lion) walk() {
    fmt.Println("Lion walk")
}

func (m lion)speak() {  
	fmt.Println("Lion speak")
}
type dog struct{
    noize string

}
type cat struct{
	age int
    name string
}
func (kat cat)walk(){
    fmt.Println("Cat goes")
}
func (kat cat) breathe() {
    fmt.Println("Cat breathe")
}
func (kat cat) speak(){
    fmt.Println("Cat speak", kat.name)
}

func main() {
    
    // var k animal
    // var a animal
    a := lion{age: 10}
    k := cat{age: 10, name: "Katya"}
    
    animal.breathe(a)
    animal.speak(k)
    animal.walk(k)
}