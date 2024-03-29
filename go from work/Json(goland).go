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
	
-- описать каждое поле json отдельно
	
	
package main

import (
	"encoding/json"
	"fmt"

)

func main() {

	m := myStruct{Name: "John Connor", Age: -1, Status: true}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", data) // {"name":"John Connor"}
}

type myStruct struct {
	// при кодировании / декодировании будет использовано имя name, а не Name
	Name string `json:"name"`
	// при кодировании / декодировании будет использовано то же имя (Age),
	// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
	// то при кодировании оно будет опущено
	Age int `json:",omitempty"`
	// при кодировании / декодировании поле всегда игнорируется
	Status bool `json:"-"`
}
Завершая рассмотрение этого вопроса нужно отметить следующее: неэкспортируемые поля (имена которых начинаются со строчной буквы) не участвуют в кодировании / декодировании.
	
	
-------------------------------------

	
	Типы Encoder и Decoder
	
	
	package main

import (
	"bytes"
	"encoding/json"
	"fmt"

)

func main() {

	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var (
		src = testStruct{Name: "John Connor", Age: 35} // структура с данными
		dst = testStruct{}                             // структура без данных
		buf = new(bytes.Buffer)                        // буфер для чтения и записи
	)

	enc := json.NewEncoder(buf) //создали объекты Encoder и Decoder
	dec := json.NewDecoder(buf)//в качестве аргумента буфер, который удовлетворяет
	// одновременно и интерфейсу io.Reader, и интерфейсу io.Writer,

	enc.Encode(src)
	dec.Decode(&dst)

	fmt.Print(dst) // {John Connor 35}
}
	
	
	
------нагляднее
	
	package main

import (
	"bytes"
	"encoding/json"
	"fmt"

)

type testStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Surname string `json:"Surname"`
 }
 
 func main() {
	var (
	   src = testStruct{Name: "John Connor", Age: 35} // структура с данными
	   dst = testStruct{}                             // структура без данных
	   buf = new(bytes.Buffer)                        // буфер для чтения и записи
	)
 
	enc := json.NewEncoder(buf)
	dec := json.NewDecoder(buf)
 
	fmt.Printf("buf before encode %v\n", buf)
	fmt.Printf("src before encode %+v\n", src)
	fmt.Printf("dst before encode %+v\n", dst)
	fmt.Println("-----------------------------")
 
	enc.Encode(src)

	fmt.Printf("buf after encode %v\n", buf)
	fmt.Printf("src after encode %+v\n", src)
	fmt.Printf("dst after encode %+v\n", dst)
	fmt.Println("-----------------------------")
 
	dec.Decode(&dst)

	fmt.Printf("buf after decode %v\n", buf)
	fmt.Printf("src after decode %+v\n", src)
	fmt.Printf("dst after decode %+v\n", dst)
 }
 
 
 
 вывод
 	 

buf before encode 
src before encode {Name:John Connor Age:35 Surname:}
dst before encode {Name: Age:0 Surname:}
-----------------------------
buf after encode {"name":"John Connor","age":35,"Surname":""}
src after encode {Name:John Connor Age:35 Surname:}
dst after encode {Name: Age:0 Surname:}
-----------------------------
buf after decode
src after decode {Name:John Connor Age:35 Surname:}
dst after decode {Name:John Connor Age:35 Surname:}



----------------------------------
 задача, из файла типа
    [
	{
		"global_id": 273478509,
		"system_object_id": "M.72.19.2",
		"signature_date": "24.03.2017 14:42:00",
		"Razdel": "M",
		"Kod": "72.19.2",
		"Name": "Научные исследования и разработки в области технических наук",
		"Idx": "M.72.19.2"
	},
	{
		"global_id": 277934797,
		"system_object_id": "S.95.29.41",
		"signature_date": "28.03.2017 09:36:10",
		"Razdel": "S",
		"Kod": "95.29.41",
		"Name": "Ремонт предметов и изделий из металла",
		"Idx": "S.95.29.41"
	}
]
	
	найти сумму в поле global_id
 
 
 
	
	package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	//"strings"

)

func main() {

	var data []byte
	file, _ := os.Open("new2.json") //прочитали из файла

	data, _ = ioutil.ReadAll(file)

	// если файл это один огромный массив тогда создаем структуру типа массив
	type AutoGenerated []struct { //у каждого элемента этого массива есть поля ниже
		GlobalID       int    `json:"global_id"`
		SystemObjectID string `json:"-"`
		SignatureDate  string `json:"-"`
		Razdel         string `json:"-"`
		Kod            string `json:"-"`
		Name           string `json:"-"`
		Idx            string `json:"-"`
		Nomdescr       string `json:"-"`
	}

	var result AutoGenerated

	if err := json.Unmarshal(data, &result); err != nil { //распарсиваем что пришло на вход
		fmt.Println(err)
		return
	} 
	var sum int
	for _, value := range result {
		sum = sum + value.GlobalID // у каждого элемента в файле(result) есть поле GlobalID
	}
	fmt.Print(sum)

}
	
	
	------------ топ ответов
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var items []struct {
		ID uint64 `json:"global_id"`
	}
	f, _ := os.Open("data-20190514T0100.json")
	json.NewDecoder(f).Decode(&items)
	var sum uint64
	for _, item := range items {
		sum += item.ID
	}
	fmt.Println(sum)
}
	
	
---------------	еще, через http
		
		package main

import (
   "encoding/json"
   "fmt"
   "net/http"
)

type Data []struct {
   GlobalID       int    `json:"global_id"`
   SystemObjectID string `json:"system_object_id"`
   SignatureDate  string `json:"signature_date"`
   Razdel         string `json:"Razdel"`
   Kod            string `json:"Kod,omitempty"`
   Name           string `json:"Name"`
   Idx            string `json:"Idx"`
   Nomdescr       string `json:"Nomdescr,omitempty"`
}

func main(){
   var StructData Data
   var urlData = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"
   var Result int
   resp, _ := http.Get(urlData)
   defer resp.Body.Close()
   r := json.NewDecoder(resp.Body)
   r.Decode(&StructData)
   for _, val := range StructData{
      Result += val.GlobalID
   }
   fmt.Println("Result:", Result)
}
	
---------------------------
	ссылки https://tutorialedge.net/golang/parsing-json-with-golang/
	https://pkg.go.dev/encoding/json#Marshal
	
	
	