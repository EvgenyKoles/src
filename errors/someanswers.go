numbers, _ := bufio.NewReader(os.Stdin).ReadString('\n')
--------------------------------------------------------	
		//fmt.Print(m)

	// for key, value := range m {
	// 	fmt.Print(key, value, " ")
	// }

-------------------------------------------
	
package main

import (
	"fmt"
)

func main() {

	var numbers int
	t := []int{}
	
	for i := 0; i <= 3; i++{
		fmt.Scan(&numbers)
		t = append (t, numbers)
	}
	
	for _, value := range t{
		if value <= 10{
			continue
		}
		if value >= 100 {
			break
		}
		fmt.Println(value)	
	}

	
-------------------------------------------------------------

 conveert word to slice of numbers
// "strings"
// "strconv"
	// strs := make([]string, 6)
	// strs = strings.Split(v, "")
	// ints := make([]int,0,6)

	// for _, s := range strs {
	//  num, err := strconv.Atoi(s); 
	// 	if err == nil {
	// 		ints = append(ints, num)
	// }
	// }
	// a := ints[0]+ints[1]+ints[2]
	// b := ints[3]+ints[4]+ints[5]
	// if a==b {
	// fmt.Println("YES")	
	// } else {
	// 	fmt.Print("NO")
	// }
---------------------------------------------


package main

import (
	"fmt"
	"strings"
"strconv"
)

func main() {

	// var numbers int
	// t := []int{}
	
	// for i := 0; i <= 1; i++{
	// 	fmt.Scan(&numbers)
	// 	t = append (t, numbers)
	// }
	
	// fmt.Print(t)
	

	var v string
	fmt.Scan(&v)

	var p string
	fmt.Scan(&p)
	
	strs := make([]string, 6) 
	strs = strings.Split(v, "")
	ints := make([]int,0,6)

	for _, s := range strs {
	 num, err := strconv.Atoi(s); 
		if err == nil {
			ints = append(ints, num)
		}
	}



	strsP := make([]string, 6) 
	strsP = strings.Split(p, "")
	intsP := make([]int,0,6)

	for _, s := range strsP {
	 num, err := strconv.Atoi(s); 
		if err == nil {
			intsP = append(intsP, num)
		}
	}

	fmt.Print(ints)
	fmt.Print(intsP)


	for _, value := range ints {
		for _, value2 := range intsP{
			if value == value2 {
				fmt.Print(value, " ")
			}
		}
	}

 
}

---------------------------------------------
package main

import (
	"fmt"
)

func main() {

	var numbers float64
	// t := []int{}
	
	// for i := 0; i <= 1; i++{
	// 	fmt.Scan(&numbers)
	// 	t = append (t, numbers)
	// }
	
	fmt.Scan(&numbers)


	if numbers <=0 {
		fmt.Printf("число %2.2f numbers не подходит", numbers)
	}

	fmt.Printf("%.4f", numbers*numbers)
	

	
 
}
---------------------------------------------



package main

import (
	"fmt"
)

func main() {

	
	t2 := [6]uint8{}
	t := [10]uint8{}
	workArray := [10]uint8{}
	//make([]uint8, len(t), cap(t))
	
	for idx := range t {
        fmt.Scan(&t[idx])
    }
	
	for idx := range t2 {
        fmt.Scan(&t2[idx])
    }
	
	// fmt.Print(t)
	// fmt.Println(t2)
	// copy(t3, t)

	for i, value := range t {
		workArray[i] = value
	}
	
	workArray[t2[0]] = t[t2[1]]
	workArray[t2[1]] = t[t2[0]]

	workArray[t2[2]] = t[t2[3]]
	workArray[t2[3]] = t[t2[2]]

	workArray[t2[4]] = t[t2[5]]
	workArray[t2[5]] = t[t2[4]]
	
	for i := range workArray{
		fmt.Printf("%d ",workArray[i])
	}
	
}

--------------------------------------------


сумма чисел

package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {

	var v string
	fmt.Scan(&v)
	var sum int

	strs := make([]string, 3)
	strs = strings.Split(v, "")
	digit := make([]int,0,3)

	for _, s := range strs {
	 num, err := strconv.Atoi(s); 
		if err == nil {
			digit = append(digit, num)
		}
	}

	for _, value := range digit {
		sum = sum + value

	}

	fmt.Print(sum)


}	

--------------------------------------------


	var v string
	fmt.Scan(&v)
	

	strs := make([]string, 3)
	strs = strings.Split(v, "")
	digit := make([]int,0,3)

	for _, s := range strs {
	 num, err := strconv.Atoi(s);
		if err == nil {
			digit = append(digit, num)
		}
	}

	fmt.Print(digit)






--------------------------------------------


package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {

	var v string
	fmt.Scan(&v)
	

	strs := make([]string, 3)
	strs = strings.Split(v, "")
	digit := make([]int,0,3)

	for _, s := range strs {
	 num, err := strconv.Atoi(s);
		if err == nil {
			digit = append(digit, num)
		}
	}

	var koren int
	
	for _, value := range digit{
		koren = koren + value
	}
		
	if koren >9 {
		var a = koren%10
		var b = koren/10
		fmt.Print(a+b)
	} else {
		fmt.Print(koren)
	}

	

}	



--------------------------------------------

package main

import (
	"fmt"
	// "strings"
	// "strconv"
)

func main() {

	var firts int
	fmt.Scan(&firts)
	var second int
	fmt.Scan(&second)
	

	var flag = false

	for i := second; i >= firts; i-- {
		if i%7 == 0 {
			fmt.Print(i)
			flag = true
			break
		}
	}	

	if !flag {
		fmt.Print("NO")
	}
}
--------------------------------------------

package main

import (
	"fmt"
	// "strings"
	// "strconv"
)

func main() {

	var n int
	fmt.Scan(&n)

	Array := make([]int, 2) 
	Array[0]=0
	Array[1]=1

	var flag = false
	if n==1 {
		fmt.Print(1)
		flag = true
	}

	if n==3{
		fmt.Print(4)
		flag=true
	}
	if n==2{
		fmt.Print(3)
		flag=true
	}

	for i:=2; i<=n; i++{
		Array = append(Array, Array[i-2]+Array[i-1])
		if n == Array[i] {
			fmt.Print(i)
			flag = true
			break
		}
	}

	if !flag{
		fmt.Print(-1)
	}
}	

-------------------------------------------

func fibonacci(x int) int {

	slice := make([]int, 3) 

	slice[0] =1
	slice[1] =1
	slice[2] =2
	
	for i := 3; i < 100; i++ {
		slice = append(slice, slice[i-2]+slice[i-1])
	}
	//fmt.Println(slice)
	
	return slice[x-1]

}


--------------------------------------------
func sumInt(agrs ... int) (int, int) {
	var sum int

	for _, value := range agrs{
		sum = sum + value
	}
	//fmt.Print(agrs)
	return len(agrs), sum
}


--------------------------------------------
func main() {

	var a int = 100

	var b = &a
	*b++
	fmt.Println(*b, " ", a)

    var c int = 100
	var d = c
	d++

	fmt.Println(d,c)

}	

--------------------------------------------

func main() {
	var b = 2
	var c = 3
	test(&b,&c)
}	

func test(x1 *int, x2 *int) {
	fmt.Print(*x1 * *x2)

}
--------------------------------------------

func main() {
	r := Rectangle{0,0,10,10}
	r2 := Rectangle{x1: 0, x2: 0, y1: 5, y2: 2}
	//c := Circle{x: 0, y: 0, r: 5} можно создать так, и потом ссылаться на x,y
	
	fmt.Println(r.area())
	fmt.Println(r2.parameters())
	}
	
type Rectangle struct { // создаем структуру
    x1, y1, x2, y2 float64
}

func distance (x1,x2,y1,y2 float64) float64 { // создаем фукнцию дистанции
	a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 { // создаем функцию, которая принимает 
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func  (r2 *Rectangle) parameters() float64 {
	return RectangleArea(r2.x1, r2.y1, r2.x2, r2.x2)
}

func RectangleArea (x1, x2, y1,y2 float64) float64 {
	return x1*x2
}


------------------------------------------------------



func main() {

testStruct := new (Dich)

fmt.Println(testStruct.Shoot())
}
	
type Dich struct {
	on bool
	ammo int
	power int
} 

func (d *Dich) Shoot() bool{

	if d.on != false && d.ammo >0 {
		d.ammo--
		return true
	}else {
		return false
	}
}

func (d *Dich) RideBike() bool {

	if d.on != false && d.power >0 {
		d.power--
		return true
	}else {
		return false
	}
}


----------------------------------------------------------------------


func main() {
	
	var s string = "Это строка"


	fmt.Print(s[:4])
}

- Эт
	/*
		Попробуем изменить что-то встроке:
		s[3] = 12
		Ошибка компиляции: cannot assign to s[3], потому что строки - неизменяемые последовательности.
	*/
----------------------------------------------------------------------

func main() {
    fmt.Println(    
        // Содержится ли подстрока в строке    
        strings.Contains("test", "es"), 
        // результат: true

        // Кол-во подстрок в строке
        strings.Count("test", "t"),
        // результат: 2

        // Начинается ли строка с префикса       
        strings.HasPrefix("test", "te"), 
        // результат: true
		
        
        // Заканчивается ли строка суффиксом
        strings.HasSuffix("test", "st"), 
        // результат: true

        // Возвращает начальный индекс подстроки в строке, а при отсутствии вхождения возвращает -1
        strings.Index("test", "e"), 
        // результат: 1
		
        // объединяет массив строк через символ
        strings.Join([]string{"hello","world"}, "-"),
        // результат: "hello-world"

        // Повторяет строку n раз подряд
        strings.Repeat("a", 5), 
        // результат: "aaaaa"

        // Функция Replace заменяет любое вхождение old в вашей строке на new
        // Если значение n равно -1, то будут заменены все вхождения.
        // Общий вид: func Replace(s, old, new string, n int) string
        // Пример:
        strings.Replace("blanotblanot", "not", "***", 	-1),
        // результат: "bla***bla***"
 
        // Разбивает строку согласно разделителю
        strings.Split("a-b-c-d-e", "-"), 
        // результат: []string{"a","b","c","d","e"}

        // Возвращает строку c нижним регистром
        strings.ToLower("TEST"), 
        // результат: "test"

        // Возвращает строку c верхним регистром
        strings.ToUpper("test"), 
        // результат: "TEST"

        // Возвращает строку с вырезанным набором
        strings.Trim("tetstet", "te"),
        // результат: s
    )
}


----------------------------------------------------------------------
	Unicode
	
	
package main

import (
	"fmt"
	"unicode"
)

func main() {
    // функции ниже принимают на вход тип rune


    // проверка символа на цифру
	fmt.Println(unicode.IsDigit('1')) // true
    // проверка символа на букву
	fmt.Println(unicode.IsLetter('a')) // true 
    // проверка символа на нижний регистр
	fmt.Println(unicode.IsLower('A')) // false
    // проверка символа на верхний регистр
	fmt.Println(unicode.IsUpper('A')) // true
    // проверка символа на пробел 
    // пробел это не только ' ', но и:
    //  '\t', '\n', '\v', '\f', '\r' - подробнее читайте в документации
	fmt.Println(unicode.IsSpace('\t')) // true 

    // С помощью функции Is можно проверять на кастомный RangeTable:
    // например, проверка на латиницу:
 	fmt.Println(unicode.Is(unicode.Latin, 'ы')) // false


    // функции преобразований
	fmt.Println(string(unicode.ToLower('F'))) // f
	fmt.Println(string(unicode.ToUpper('f'))) // F
}
	
	
----------------------------------------------------------------------
	проверка на точку и на заглавную букву
	
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
	

	//fmt.Println(string(rs[len(rs)-3])== ".")
	//fmt.Println(rs[utf8.RuneCountInString(text)-1])
	//strings.HasSuffix(text, string(rs[len(rs)-3]))
	//fmt.Println(unicode.IsUpper(rs[0]))

	if string(rs[len(rs)-3]) == "." && unicode.IsUpper(rs[0]) {
	fmt.Print("Right")
	
	}else {
		fmt.Print("Wrong")}

}

			
			
----------------------------------------------------------------------
func main(){

	fmt.Print(T())
	

}

var k,p,v float64 = 1296,6,6

func T() float64{
	return 6/W()
}

func W() float64 {
	return math.Sqrt(k/M())
}

func M() float64{
	return p*v
}






----------------------------------------------------------------------


func main(){

	// с помощью встроенной функции make:
m1 := make(map[int]int)

// с помощью использования литерала отображения:
m2 := map[int]int{
    // Пары ключ:значение указываются при необходимости
    12: 2,
    1:  5,
}

fmt.Println(m1) // map[]
fmt.Println(m2) // map[1:5 12:2]
	

}

----------------------------------------------------------------------
	package main

import (

	"fmt"
	"time"
	//"os"

)

func main(){
 
	m := make(map[int]int)
	
	arr := make([]int, 10)
	for i := range arr {
   		fmt.Scan(&arr[i])

		if value, ok := m[arr[i]]; ok {
			fmt.Print(value, " ") 
			
		}else {
			m[arr[i]]=work(arr[i])
			fmt.Print(m[arr[i]], " ")
			
		}

	}

}
		
	func work(x int) int {
		time.Sleep(time.Second)
		if x >= 4 {
		  return x + 1
		} else {
		  return x - 1
		}
	}

	
----таже тема
	
	package main

import (
	"fmt"
	"time"
)

func main(){
 
	m := make(map[int]int)
	
	for i := 0; i < 10; i++ {
	   	fmt.Scan(&i)

		if value, ok := m[i]; ok {
			fmt.Print(value, " ") 
		}else {
			m[i]=work(i)
			fmt.Print(m[i], " ")
		}
	}
}
		
	func work(x int) int {
		time.Sleep(time.Second)
		if x >= 4 {
		  return x + 1
		} else {
		  return x - 1
		}
}

	
----------------------------------------------------------------------
	
пробежатся по мапе
	
		for key, value := range mapName {
    fmt.Println(key, value)
}

	
есть ли значение по ключу
	
	if value, inMap := m[1]; inMap {
	fmt.Println(value) // 10
}
	
	
	
----------------------------------------------------------------------
		
		package main

import (

	"fmt"
)

func main() {

	groupCity := map[int][]string{
		10:   []string{"Деревня", "Село"}, // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Большой город","Пригород","Дичь","Дичь2"}, // города с населением 100-999 тыс. человек
		1000: []string{"Миллионик","Мегамиллионник"}, // города с населением 1000 тыс. человек и более
	 }
	 

	cityPopulation := map[string]int{
		"Село" : 10,
		"Город" : 100,
		"Миллионик" : 1000,
		"Большой город" : 10000,
		"Пригород" : 100000,
	}
	 flag := false
	
	for k, _ := range cityPopulation {
		flag = false
		for i :=0; i < len(groupCity[100]); i++ { //обращение слайсу(который один из значений мапы)
			if k == groupCity[100][i]{
				//fmt.Println(k)
				flag = true
				break
			}
		}
		if !flag {
			delete(cityPopulation,k)
		}
	}
	  //fmt.Println(cityPopulation)

}


------------------------------------------------------------------
конвертация строк в слайсы байт(слайсы рун) и назад
	


package main

import (
    "fmt"
)

func main() {
    a := "str"

    b := []byte(a)

    c := string(b)

    fmt.Println(a) // str

    fmt.Println(b) // [115 116 114] - побайтовый срез

    fmt.Println(c) // str
}



package main

import (
    "fmt"
)

func main() {
    a := "строка"
    b := []rune(a) // срез рун
    c := string(b)
    fmt.Println(a) // строка
    fmt.Println(b) // [1089 1090 1088 1086 1082 1072] - срез рун
    fmt.Println(c) // строка
}


------------------------------------------------------------------
Конвертация чисел с плавающей запятой в строку 
Для этого есть функция FormatFloat:

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a float64 = 1.0123456789

	// 1 параметр - число для конвертации
	// fmt - форматирование
	// prec - точность (кол-во знаков после запятой)
	// bitSize - 32 или 64 (32 для float32, 64 для float64)
	fmt.Println(strconv.FormatFloat(a, 'f', 2, 64)) // 1.01

	// если мы хотим учесть все цифры после запятой, то можем в prec передать -1
	fmt.Println(strconv.FormatFloat(a, 'f', -1, 64)) // 1.0123456789

	// Возможные форматы fmt:
	// 'f' (-ddd.dddd, no exponent),
	// 'b' (-ddddp±ddd, a binary exponent),
	// 'e' (-d.dddde±dd, a decimal exponent),
	// 'E' (-d.ddddE±dd, a decimal exponent),
	// 'g' ('e' for large exponents, 'f' otherwise),
	// 'G' ('E' for large exponents, 'f' otherwise),
	// 'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
	// 'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).
	var b float64 = 2222 * 1023 * 245 * 2 * 52
	fmt.Println(strconv.FormatFloat(b, 'e', -1, 64)) // 5.791874088e+10
}
 Так же можно использовать пакет "fmt". Он обладает удобным методом Sprintf. Вот шпаргалка по всему пакету.

package main

import (
    "fmt"
)

func main() {
    fmt.Println(fmt.Sprint(20.19)) // Краткая форма

    a := 20.20
    fmt.Println(fmt.Sprintf("%f", a)) // Полная форма
}
Внимание! Использовать fmt для конвертации нежелательно из-за того что производительность ниже по сравнению с strconv.

Конвертация bool в string
Тут все просто:

var a = true
res := strconv.FormatBool(a)
fmt.Println(res)     	// true
fmt.Printf("%T", res)   // string

	
------------------------------------------------------------------
strconv.Atoi()			string -> int
strconv.Itoa()			int -> string
strconv.ParseInt()		string -> int64
strconv.ParseUint()		string -> uint64
strconv.FormatInt()		int64 -> string
strconv.FormatUint()	uint64 -> string
strconv.ParseBool()		string -> bool
strconv.ParseFloat()	string -> float64

------------------------------------------------------------------
	
------------------------------------------------------------------
	




