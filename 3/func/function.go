package main

import (
	"fmt"
	"math"
)

// создаем струкктуру
type Example struct{
	flag bool
	counter int16
	pi float32
}

func main() {

	var e1 Example

	fmt.Printf("%+v\n", e1)

	e2 := Example{
		flag: true,
		counter: 2,
		pi: 3.14,
	}

	fmt.Printf("%+v\n",e2 )

	// короткая запись, но надо указать все поля
	e3 := Example{false, 4, 3.12}

	fmt.Printf("%+v\n",e3)


	fmt.Println("__________________")
	///metods



	// когда обьявляешь тип. можешь обьявить его методы 
	//  (по сути фукнкции для этого метода)

	// type myInt int 

	// func (m myInt) showYourSelf() {
	// 	fmt.Printf("%T %v\n", m, m)
	// }

	// func (m * myInt) add(i myInt) {
	// 	*m = *m + myInt(i)
	// }

v := Vertex{3, 4}


	v.Sqrt(3)
	fmt.Println(v.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	v.noScale(10)
	fmt.Println(v.Abs())

	fmt.Println(v.Average())

	v.fu(4)
	fmt.Println(v.Abs(), "Fu")

// sa := SecretAgent{Person: Person{"James", 007}, LicenseToKill: false}

// fmt.Printf("%T %+v\n", sa, sa)
// fmt.Println("secret inn", sa.GetName())
// fmt.Println(sa.GetName())

} //main


type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}
// * значит возьми пердыдущий
// без * возьми изначальный 
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Sqrt (f float64) {
	v.X = v.X + f*2
	v.Y = v.Y + f*2
}

func(p Vertex) dich () float64{
	return p.X *2 
}


func (v Vertex) noScale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Average () float64{
	return ((v.X + v.Y)/2)
}
func (v Vertex) fu (z float64)  {
	v.X =  v.X * v.Y * z
	v.Y = z *3
}
	

	
//встраивание


type Person struct{
	Name string
	inn int
}

type Stuff struct{
	inn int
}

type SecretAgent struct{
	Person
	Stuff
	LicenseToKill bool
}
// могу использовать этот метод, т.к  структура Person есть в SecretAgent
func (p Person) GetName() string {
	return p.Name
}

// можно написть фукнцию для имени но для структуры агента
func (s SecretAgent) GetName() string{
	return "Classified"
}






