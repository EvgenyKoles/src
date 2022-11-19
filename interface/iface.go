package main

import (
	"fmt"
	"time"
)

//определяют поведение. интерфейс не может содержать данные, только методы
// структура хранит данные, интерфейс поведение

// интерфейс реализующий метод Fly
type Flyer interface {
	Fly()
	// Greed()
}

// func (b Bird) Greet(){
// 	fmt.Println("Hey there")
// }

type Bird struct {
	Name string
}

type Mig45 struct{
	Name string
}

//  метод на структуре выше
func (b Bird) Fly() {
	fmt.Println(b.Name + "is flying")
}

// func DoFly(f Flyer) {
// 	f.Fly()
// }


func (m Mig45) Fly() {
	fmt.Println("Mig Flied away")
}

func GoFly(f Flyer) {
	f.Fly()
	if b, ok := f.(Bird); ok {
		fmt.Println(b.Name)
	}
}



//--------------------------------------

type Printer interface {
	Print()
}

//создадим два обьекта реализующие интрефейс

type User struct{
	name string
	age int
	lastname string
}

type Document struct{
	name string
	DocumentType string
	date time.Time
}

//функция Print для структуры Document

func (d Document) Print() {
	fmt.Printf("Document name: %s, type: %s, date: %s \n", d.name, d.DocumentType, d.date)
}

func (u User) Print() {
	fmt.Printf("Hi I am %s %s and I am %d years old \n", u.name, u.lastname, u.age)
}

	// Предположим, нам нужно написать новый метод, выводящий подробности 
	// этих двух структур. 
	// Для этого можно использовать имеющийся интерфейс. 
	// Эта функция получает в качестве аргумента любые объекты, реализующие указанный интерфейс. 
	// Так что, если объект отвечает на методы, определенные в интерфейсе, значит его можно с ее помощью обработать.
func Process(obj Printer){
	obj.Print()
}


func main() {
	// duckPlane := Bird{"Duck plane"}
	// GoFly(duckPlane)

	u := User {name: "Ivan", age: 23, lastname: "Ivanovich"}
	doc := Document {name: "Passport", DocumentType: "csv", date: time.Now()}


	Process(u)
	Process(doc)
	

	

}


