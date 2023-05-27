{
    "firstName": "Иван",
    "lastName": "Иванов",
    "address": {
        "streetAddress": "Московское ш., 101, кв.101",
        "city": "Ленинград",
        "postalCode": 101101
    },
    "phoneNumbers": [
        "812 123-1234",
        "916 123-4567"
    ]
}



switch to go from fson https://mholt.github.io/json-to-go/

-----------------------------------------
	
	Начнем с функции для кодирования данных Marshal:

type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

s := myStruct{
	Name:   "John Connor",
	Age:    35,
	Status: true,
	Values: []int{15, 11, 37},
}

// Функция Marshal принимает аргумент типа interface{} (в нашем случае это структура)
// и возвращает байтовый срез с данными, кодированными в формат JSON.
data, err := json.Marshal(s)
if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("%s", data) // {"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}


-----------------------------------
	
	
MarshalIndent похож на Marshal, но применяет отступ (indent) для форматирования вывода. Каждый элемент JSON в выходных данных начинается с новой строки, начинающейся с префикса (prefix), за которым следует один или несколько отступов в соответствии с вложенностью:

type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

s := myStruct{
	Name:   "John Connor",
	Age:    35,
	Status: true,
	Values: []int{15, 11, 37},
}

data, err := json.MarshalIndent(s, "", "\t")
if err != nil {
	fmt.Println(err)
	return
}

fmt.Printf("%s", data)

//{
//	"Name": "John Connor",
//	"Age": 35,
//	"Status": true,
//	"Values": [
//		15,
//		11,
//		37
//	]
//}


неэкспортируемые поля (имена которых начинаются со строчной буквы) не участвуют в кодировании / декодировании
---------------------------------------
	
Ну и в завершении этого шага рассмотрим последнюю из трех функций Unmarshal, она принимает в качестве аргумента байтовый срез и указатель на объект, в который требуется декодировать данные. Рассмотрим это на уже знакомом примере:

data := []byte(`{"Name":"John Connor","Age":35,"Status":true,"Values":[15,11,37]}`)

type myStruct struct {
	Name   string
	Age    int
	Status bool
	Values []int
}

var s myStruct

if err := json.Unmarshal(data, &s); err != nil {
	fmt.Println(err)
	return
}


---------------------------------------
	
	проверка на json

	
	type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

m := user{Name: "John Connor", Email: "test email"}

data, _ := json.Marshal(m)

data = bytes.Trim(data, "{") // испортим json удалив '{'

// функция json.Valid возвращает bool, true - если json правильный
if !json.Valid(data) {
	fmt.Println("invalid json!") // вывод: invalid json!
}

fmt.Printf("%s", data) // вывод: "Name":"John Connor","Email":"test email","Status":false,"Language":null}
	
	
	
---------------------------------------
	задачи по json. на вход подается json
{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3,4]
        },
        {
            "LastName":"Ien",
            "FirstName":"ccc",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[5,2]
        }
    ]
}
	
Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы. Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:

{
    "Average": 14.1
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

)

func main() {

	var data []byte
	data, err := ioutil.ReadAll(os.Stdin) // прочитали ввод
	if err != nil {
		fmt.Println("invalid input") 
	}

	type (
		Student struct { // массив входящий в group
			Rating []int
		}
		Group struct { // общее json
			Students []Student
		}
		Rating struct { // массив- "Rating":[5,2]
			Average float32 // тут назвали Average, т.к в выводе надо именно Average
		}
	)

		var result Group

	if err := json.Unmarshal(data, &result); err != nil { //распарсиваем что пришло на вход
		fmt.Println(err)
		return
	} else {
		//fmt.Print(string(data))
	}

	if !json.Valid(data) {
		fmt.Println("invalid json!") // вывод: invalid json!
	}
	var countofstudents int
	var countofnumbers int

	for i, _ := range result.Students{ // пробегаемся по студентам, чситаем колечество result.Students[] - {[1 2 3]} {[5 2]}
		countofstudents = i
	}	

	for _, value := range result.Students{ // пробегаемся по рейтингу считаем кол-во оценок
		countofnumbers = countofnumbers + len(value.Rating)
	}
	
	type Output struct { //создаем структуру для вывода
		Average   float32
	}

	k := Output { 
		Average: float32(countofnumbers)/float32(countofstudents+1),
		
	}

	data2, err := json.MarshalIndent(k, "", "    ")// переводим в json
	if err != nil {
		fmt.Println(err)
	return
	}
	
	fmt.Printf("%s", data2)
	
}



-------------------------------------
	
	
	
	

	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	