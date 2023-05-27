package main
import "fmt"

type Creature struct {  //создали структуру с полями, она имеет метод Greet
	Name     string
	Greeting string
}

func (c Creature) Greet() { // создали метод для этой структуры, в получатели присвоили экземпляр Creature для переменной С. Что бы мы могли обращатся к полям Creature при использовании функции
	fmt.Printf("%s says %s", c.Name, c.Greeting)
}

func main() {
	sammy := Creature{ //мы создали экземпляр Creature, указали значения для полей Name и Greeting.
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	
	Creature.Greet(sammy) //Здесь мы вызвали метод Greet, объединив имя типа и имя метода с помощью оператора . и предоставив экземпляр(строка 15) Creature в качестве первого аргумента. Пишем так: Структура.метод(экземпляр)
}
Output
Sammy says Hello!

------------------------------------------------
package main
import "fmt"

type Creature struct {
	Name     string
	Greeting string
}

func (c Creature) Greet() {
	fmt.Printf("%s says %s", c.Name, c.Greeting)
}

func main() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	sammy.Greet() // Этот пример идентичен предыдущему, но в этот раз мы использовали запись через точку для вызова метода Greet с помощью Creature, который хранится в переменной sammy как получатель. Пишем: экземпляр.метод. Экземпляр видит этот метод через структуру. Это сокращенная форма записи для вызова функции в первом примере.
}
	
Output
Sammy says Hello!

----------------------------------------------------------------

package main
import "fmt"

type Creature struct {
	Name     string
	Greeting string
}

func (c Creature) Greet() Creature {
	fmt.Printf("%s says %s!\n", c.Name, c.Greeting)
	return c //который будет возвращать Creature, чтобы мы могли использовать дополнительные методы для sammy
}

func (c Creature) SayGoodbye(name string) {
	fmt.Println("Farewell", name, "!")
}
Та
func main() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	R
	sammy.Greet().SayGoodbye("gophers")//Цепочка точек также указывает последовательность, в которой вызываются методы. Мы просто можем вызывать все методы через точку

	Creature.SayGoodbye(Creature.Greet(sammy), "gophers")// no common
}


Output
Sammy says Hello!!
Farewell gophers !
Sammy says Hello!!
Farewell gophers !

-----------------------------------------------------

//Тип интерфейса — это спецификация методов, используемых компилятором для гарантии того, что тип обеспечивает реализацию этих методов. Любой тип, который имеет методы с тем же именем, теми же параметрами и теми же возвращаемыми значениями, что и методы, которые находятся в определении интерфейса, реализует этот интерфейс и может привязываться к переменным с данным типом интерфейса.




package main

import (
	"fmt"
	"strings"
)

type Stringer interface {
  String() string
}

type Ocean struct { //определяется новый тип структуры под названием Ocean. Ocean, реализует интерфейс fmt.Stringer, поскольку Ocean определяет метод под названием String, который не принимает никаких параметров и возвращает строку
	Creatures []string
}

func (o Ocean) String() string {
	return strings.Join(o.Creatures, ", ")
}

func log(header string, s fmt.Stringer) { //мы используем fmt.Println и вызываем метод String из Ocean, когда он получает fmt.Stringer в качестве одного из своих параметров.
	fmt.Println(header, ":", s)
}

func main() {
	o := Ocean{ //новый экземпляр Ocean
		Creatures: []string{
			"sea urchin",
			"lobster",
			"shark",
		},
	}
	log("ocean contains", o) //передали его функции log, которая получает строку для вывода, после чего следует что-то, реализующее fmt.Stringer. Компилятор Go позволяет нам передавать o здесь, поскольку Ocean реализует все методы, запрашиваемые fmt.Stringer.
//Если Ocean не предоставляет метод String(), Go будет генерировать ошибку компиляции, поскольку метод log запрашивает fmt.Stringer в качестве аргумента.
//Go также необходимо убедиться, что метод String(), который был предоставлен, полностью соответствует методу, запрашиваемому интерфейсом fmt.Stringer
}


Output
ocean contains : sea urchin, lobster, shark

-----------------------------------
Получатели по указателю и интерфейсы


package main

import "fmt"

type Submersible interface { //интерфейс под названием Submersible, который требует типы с методом Dive()
	Dive()
}

type Shark struct { //определили тип Shark с полем Name и методом isUnderwater для отслеживания состояния Shark
	Name string
	isUnderwater bool
}

func (s Shark) String() string { //Также мы определили метод String() получателя по значению, чтобы он мог полностью выводить на экран состояние Shark. Используя fmt.Println путем применения интерфейса fmt.Stringer
	if s.isUnderwater {
		return fmt.Sprintf("%s is underwater", s.Name)
	}
	return fmt.Sprintf("%s is on the surface", s.Name)
}

func (s *Shark) Dive() { //определили метод Dive() для получателя по указателю для типа Shark, который изменяет возвращаемое методом isUnderwater значение на true.
	s.isUnderwater = true
}

func submerge(s Submersible) { //использовали функцию submerge, которая получает параметр Submersible. Далее нам надо определить метод Dive(он выше) который должен принимать Shark, что бы мы могли в submerge передать Shark. 
	s.Dive()
}

func main() {
	
	//мы определили переменную s, которая указывает на Shark, и немедленно вывели s с помощью fmt.Println.
	s := &Shark{
		Name: "Sammy",
	}

	fmt.Println(s)

	submerge(s)

	fmt.Println(s)
}


--------------------------------------
	
	проверка типов в интерфейсе
		
package main
import "fmt"

var i interface{} = 12

func main() {

if v, ok := i.(int); ok {
	fmt.Println(v+12) // Суммирование не произойдет, если ok == false
	}
    
}


more 
	func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Умножим на 2:", v*2)
	case string:
		fmt.Println(v + " golang")
	default:
		fmt.Printf("Я не знаю такого типа %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}

----------------------------

package main

import (
	"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
	"os"            // пакет используется для проверки ответа, не удаляйте его
)




func main() {

	value1, value2, operation := readTask()

	if _, ok := value2.(float64); ok 
		if _, ok := value1.(float64); ok {
			switch v := operation.(type) {
			case string:
				if operation == "+" {
					result := value1.(float64) + value2.(float64)
					fmt.Printf("%.4f", result)
					break
				}
				if operation == "-" {
					result := value1.(float64) - value2.(float64)
					fmt.Printf("%.4f", result)
					break
				}
				if operation == "*" {
					result := value1.(float64) * value2.(float64)
					fmt.Printf("%.4f", result)
					break
				}
				if operation == "/" {
					result := value1.(float64) / value2.(float64)
					fmt.Printf("%.4f", result)
					break
				} else {
					fmt.Println("неизвестная операция")
					//os.Exit(0)
					return
				}
			default:
				fmt.Printf("Я не знаю такого типа %T!\n", v)
				//os.Exit(0)
				return
			}
		} else {
			fmt.Printf("value=%v: %T", value1, value1)
			//os.Exit(0)
			return
		}

	} else {
		fmt.Printf("value=%v: %T", value2, value2)
		//os.Exit(0)
		return
	}
	os.Exit(0)
	return

} 


---------------------------------------



package main

import "fmt" //

type Stringer interface {
	String() string
}

type batteryForTest struct{
	Name string
}

func (c batteryForTest) String() string {

	rs := []byte(c.Name)
	var a = []byte{}
	var count = 0
	
	for _ , value  := range rs {
		if string(value) == "0" {
			a = append(a, ' ')
			count++
		}
	}
	for i := 0; i < 10-count; i++ {
		a = append(a, 'X')
	}
	
    return fmt.Sprint("[", string(a), "]")
	//fmt.Print("[",string(a),"]")
	//return fmt.Sprintf("%v (%d)", a.Name, a.Age)
	//return string(a)
}

func main() {
	var text string
	fmt.Scan(&text)

	batteryForTest := batteryForTest{
		Name: text,
	}
	fmt.Print(batteryForTest)
 }
 
 -----------------------------------------------------













-----------------------------------------------


