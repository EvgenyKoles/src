//package main

import (
	"fmt"
	//"runtime/trace"
	"os"

)

// func readTask(j interface{}) {
// 	switch v := j.(type) {
// 	case float64:
// 		fmt.Printf("%.4f", v)
// 	case string:
// 		fmt.Println(v + " golang")
// 	default:
// 		fmt.Printf("Я не знаю такого типа %T!\n", v)
// 	}

//     switch v2 := j.(type) {
// 	case float64:
// 		fmt.Printf("%.4f", v2)
// 	case string:
// 		fmt.Println(v2 + " golang2")
// 	default:
// 		fmt.Printf("Я не знаю такого типа2 %T!\n", v2)
// 	}

// }

func readTask() (value1, value2, operation interface{}) {

	value1 = 1
	value2 = 2.2123
	operation = "/"

	if _, ok := value1.(float64); ok {
	    switch v := operation.(type) {
	    case string:
		    if v == "+" {
			result := value1.(float64) + value2.(float64)
			fmt.Printf("%.4f", result)
		}
		    if v == "-" {
			result := value1.(float64) - value2.(float64)
			fmt.Printf("%.4f", result)
		}
		    if v == "*" {
			result := value1.(float64) * value2.(float64)
			fmt.Printf("%.4f", result)
		}
		    if v == "/" {
			result := value1.(float64) / value2.(float64)
			fmt.Printf("%.4f", result)
		}   else {
			fmt.Println("неизвестная операция")
		}
	    default:
		fmt.Printf("Я не знаю такого типа %T!\n", v)
		os.Exit(0)
	}
	} else {
		fmt.Printf("value=%v: %T", value1, value1)
        os.Exit(0)
	}

	
	
}

func main() {

	fmt.Print(readTask())
	os.Exit(0)
}


--------------------------------------------------------------------------------------------------------
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

-------------------------------------------------------------------------------------------------------------
//просто структура
package main
import (
	"fmt"
)

type animal interface {
	makeSound()
}

type cat struct{}
type dog struct{}

func (c *cat) makeSound() {
	fmt.Println("meow!")
}

func (c *dog) makeSound() {
	fmt.Println("gav gav")
}


func main() {

	var c, d animal= &cat{}, &dog{}

	c.makeSound()
	d.makeSound()

}


------------------------------------------------
	
функции умеют принимать в качестве типов данных для своих агрументов входящих - интерфейсы
	
	package main


import (
	"fmt"

)


type greeter interface {
	greet(string) string
}


type russian struct{}
type american struct{}


func (r *russian) greet(name string) string {//определяем greet(который в интерфейсе) что бы он принимал нашу структуру
	return fmt.Sprintf("Привет %s", name)
}

func (a *american) greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}


func sayHello(g greeter, name string){
	fmt.Println(g.greet(name))
}

func main() {

	sayHello(&russian{}, "Petya2")
	sayHello(&american{}, "Bill2")

}

----------------------------------------------
	
//func main() {

	value1, value2, operation := readTask()

	if _, ok := value2.(float64); ok {
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

-------------------------------------------------
	

//package main


import (
	//"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
	//"os"            // пакет используется для проверки ответа, не удаляйте его
)


func readTask() (value1, value2, operation interface{}) {
	return 5.0, 5.6, "/" //тут играемся с параметрами
	}

func main() {
    vi1, vi2, operation := readTask()
    vi := [2]interface{}{vi1, vi2}
    var v1, v2 float64
    vf := [2]*float64{&v1, &v2}
    var ok bool
    for i, v:= range vi{
        if *vf[i], ok = v.(float64); !ok {
            fmt.Printf("value=%v: %T", v, v)    
            return
        }
    }
    ops:= map[string]func() float64{
        "+": func() float64 {return v1+v2 },
        "-": func() float64 {return v1-v2 },
        "*": func() float64 {return v1*v2 },
        "/": func() float64 {return v1/v2 },
    }
    
    if oper, ok := operation.(string); ok {
        if fun, ok:= ops[oper]; ok{
            fmt.Printf("%.4f", fun())	
            return
        }
    }
    fmt.Print("неизвестная операция")
    
} 

-------------------------------------------------
// работа с пакетом error

//package main

import (
	"fmt"
	"unicode"
)

type customError int 
// создали тип данных
//для того чтобы тип данных customError поддерживал интерфейс error мы реализовали в нем метод Error() string

func (c customError) Error() string {
	return fmt.Sprintf("цифра, индекс %d", c)
}

func errorInString(str string) error {
	// Полезная работа со строкой проигнорирована
	for i, s := range str {
		if unicode.IsDigit(s) {
			return customError(i)
		}
	}
	return nil
}

func main2() {

	var err error

	err = errorInString("st3ringstring")
	// к этому моменту выходит что теперь у нас есть
	// переменная err с типом интерфейс error и значение этой перменной 6

	if err != nil {
		fmt.Printf("Ошибка обработана: %v\n", err.Error())
	}

	if cError, ok := err.(customError); ok {
		fmt.Printf("Контекст: %d\n", cError)
	}

	// Output:
	// Ошибка обработана: цифра, индекс 6
	// Контекст: 6
}----------------------------------------------------------

//проверка типов

var i interface{} = "hello"

s := i.(string)
fmt.Println(s)

s, ok := i.(string)
fmt.Println(s, ok)

f, ok := i.(float64)
fmt.Println(f, ok)

f, ok = i.(float64) // panic
fmt.Println(f)

->
hello
hello true
0 false
0

var i interface{} = 12

if v, ok := i.(int); ok {
	fmt.Println(v+13) // Суммирование не произойдет, если ok == false
}



-------------------------------------------------------

// переопределили печать
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Stringer interface {
	String() string
}

type Animal struct {
	Name string
	Age  uint
}

type AnimalWithStringer struct {
	Name string
	Age  uint
}

func (a animalWithStringer) String() string {

	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}


func main() {

	animal := Animal{
		Name: "Gopher",
		Age:  2,
	}

	animalWithStringer := AnimalWithStringer{
		Name: "Gopher",
		Age:  2,
	}

	fmt.Println(animal)
	fmt.Println(animalWithStringer)
}


-----------------------------