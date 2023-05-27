package main

import (
	"fmt"
)

func main() {

    // fmt.Print("Enter a number of students: ")
    // var input float64
    // fmt.Scanf("%f", &input)


    // output := input * 1

    // fmt.Println(input)    
	// fmt.Println(output)    
	// for i := 0; i < 2; i++ {
	
	// }

	//слайс
	a1 :=[] int {10,20,30}
	s1 :=[] int {1,2,3}

	a1 = append(a1, 3)//добавление одного элемента
	s1 = append(s1, a1...)//добавление слайса в слайс
	
	//обьявим переменную типа набор слайсов
	var slm [][]int
	slm = append(slm, a1,s1)

	fmt.Println(a1)
	fmt.Println(s1)

	fmt.Println(slm) //будет - [[10 20 30 3] [1 2 3 10 20 30 3]]


 //создать слайс нужной длины
	slice3 :=make([]int, 5)

	slice3 = append(slice3, []int{6,7,8}...)
	slice4 := slice3

	fmt.Println(slice3)

 //внутри слайса - ссылк ана массивб она копируетсся если просто присвоить

 	fmt.Println(slice3)
 	fmt.Println(slice4)
 	fmt.Println("----------")

	slice4 = append(slice4, []int {3,4,5}...) //переопределили слайс4. теперь ссылается на другую область памяти
	
	fmt.Println(slice3)
	fmt.Println(slice4)

	//копирование слайса
	var slice5 []int
	copy(slice3,slice4)	
	fmt.Println(slice5)

	slice6 :=make([]int, 11,11)
	copy(slice6,slice4)	//в скобках. куда и откуда
	fmt.Println(slice6)

	//часть слайса

	 fmt.Println("часть слайса", slice3[6:8], slice4[9:11])



	//pack


	fmt.Printf("slece elements is %v\nnewline\n", slice3)


	//maps
		//empty map, can't add value or key
	var mm map [string] string //key type, value type
	fmt.Println("uninitialized map", mm)

		//right way to create mao


	var mm2 map[string] string = map [string]string{}
	
	mm2["test"] = "ok"
	fmt.Println(mm2)


	//короткая инициализация

	var mm3 = make(map[string]string)
	mm3 ["dich"] = "losh"
	fmt.Println(mm3)



	//получение значения мапы

	firstname := mm3["dich"]
	fmt.Println("firstname", len(firstname), firstname)


	//если обратится к несуществующему ключу выдасть значение по умочанию
	
	//проверка на то что значение есть

	firstname, ok := mm3["dich"]

	fmt.Println(firstname, ok)

	fmt.Println("-------------")
	delete(mm3,"dich") // key delete

	firstname2, ok := mm3["dich"] //проверим еще раз
	fmt.Println(firstname2,ok)

	mm3["dich2"] = "losh2"
	fmt.Println(mm3)

	
	//порядок ключей в мапе в случайном порядке

	mm4 := mm3
	mm4["dich2"] = "loch3"  //обновится в обеих мапах

	fmt.Println(mm3,mm4)

	
		//управляющие структуры
	a:= true
	if !a {
		println("hello world")
	}

	b:=1
	if b ==1 && b >0 {
		println("дичь")
	}


	// условие при проверке мапов

	mm5:= map[string]string{
		"firstname" : "Vasily", 
		"secondname" : "Zenya"}


		//пишем условие потом ; и булевое значение что оно есть
	if _, ok := mm5["firstname"]; ok {  
		println("firstname exist-", ok)
	}else {
		println("no firstname")
	}

	//циклы

	for {
		println("sicle")
		break
	}


	println("*****************")

	s7 := []int{3,4,5,6,7,8}
	value :=0
	idx :=0

	//операции по слайсам

	for idx < 4 { // у s1 индекс 6, если он меньше 4
		if idx < 2 {
			idx++
			continue // если он меньше 2х то я иду дальше
		}
		value = s7[idx] // иначе присваиваем значение индесу и увеличиваем индекс
		idx++
		println("while-style loop, idx:", idx, "value:", value)
	}


	for i :=0; i < len(s7); i++ {
	fmt.Println("loop", i , s7[i])
	
	}
	
	//просто перебрать слайс

	for idx := range s7 {
		println("range slice by index", idx)
	}


	for idx, val := range s7 {
		println("range slice by index", idx, val)
	}

	println("---------------")

	//for по мапе. в качестве ключа ключ мапы, можно подавить ключ (_) если надо только по значению

	for key,val := range mm4{
		println("range map by key-val", key,val)
	}


	mm8:= map[string]string{
		"firstname" : "Petr", 
		"secondname" : "zenya"}

	mm8["flag"] = "ok"
	fmt.Println(mm8)

	switch mm8["firstname"] {
	case "vasily", "Evgeny" :
		println("switch - name is Vasily")
		//не переходит в другой вариант по - умолчанию. если хочешь дальше то fallthrough
	case "Petr" :
		if mm8 ["flag"] == "ok"{
			break //выходим
		}
		println("switch - name is Petr")
	fallthrough // переходим в следующий вариант
		default :
		println("switch - other name")
	}

	//or


	switch {
	case mm8 ["firstname"] == "vasily":
		println("switch2 - vasily")
	case mm8["firstname"] == "Petr":
		println("switch2 - pert")

	}

	//выйти из цикла в Go

	fmt.Println("-------------")

	Myloop:
		for key, val := range mm8 {
			println("switch is loop", key, val)
			switch{
			case key == "firstname" && val == "vasily":
				println("switch - break loop here")
				break Myloop // выходит из всего for
			}
			
		}


	//перебрать по строке

	str1:= "Привет Мир"

	fmt.Println("ru", str1, len(str1))


	for  index, ruValue := range str1{
		fmt.Printf("%#U at position %d\n", ruValue, index)
	}


	// пакеты
	
	//импортируем пакет


	fmt.Println("-------------")
	//обьявляем функцию
	//bool -  тип возвращаемого значения
	
	// func CheckStartingLevel (level int) bool {
	// 	return level == startingLevel
	// }


	//---------------------------------------------
	// d :=((a/10)%10)
	// b := (a%10)

	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(b)

// conveert word to slice of numbers
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

// 	var n int
// // считываем числа пока не будет введен 0
// for fmt.Scan(&n); n != 0; fmt.Scan(&n){
// 	fmt.Println(n)
// }	

//-------------------------------------------

func minimumFromFour() int {
	var a,b,c,d,min int
 fmt.Scan(&a,&b,&c,&d)

 if min <= a && a <b && a<c && a <d {
 return a	
 } else if min <= b && b <a && b<c && b < d{
	 return b
 }else if min <= c && c <b && c<a && c < d{
	 return c
 }else {
	 return d
 }
}




	
}
	


