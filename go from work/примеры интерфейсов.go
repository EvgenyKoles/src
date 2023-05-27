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
	


-------------------------------------------------
	



-------------------------------------------------