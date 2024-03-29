Дата и время 
	
// func Parse(layout, value string) (Time, error)
// парсит дату и время в строковом представлении
firstTime, err := time.Parse("2006/01/02 15-04", "2020/05/15 17-45")
if err != nil {
	panic(err)
}

// LoadLocation находит временную зону в справочнике IANA
// https://www.iana.org/time-zones
loc, err := time.LoadLocation("Asia/Yekaterinburg")
if err != nil {
	panic(err)
}

// func ParseInLocation(layout, value string, loc *Location) (Time, error)
// парсит дату и время в строковом представлении с отдельным указанием временной зоны
secondTime, err := time.ParseInLocation("Jan 2 06 03:04:05pm", "May 15 20 05:45:10pm", loc)
if err != nil {
	panic(err)
}

fmt.Println(firstTime.Format("02-01-2006 15:04:05"))  // 15-05-2020 17:45:00
fmt.Println(secondTime.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:10
	----
stdLongMonth      = "January"
stdMonth          = "Jan"
stdNumMonth       = "1"
stdZeroMonth      = "01"
stdLongWeekDay    = "Monday"
stdWeekDay        = "Mon"
stdDay            = "2"
stdUnderDay       = "_2"
stdZeroDay        = "02"
stdHour           = "15"
stdHour12         = "3"
stdZeroHour12     = "03"
stdMinute         = "4"
stdZeroMinute     = "04"
stdSecond         = "5"
stdZeroSecond     = "05"
stdLongYear       = "2006"
stdYear           = "06"
stdPM             = "PM"
stdpm             = "pm"
stdTZ             = "MST"
stdISO8601TZ      = "Z0700"  // prints Z for UTC
stdISO8601ColonTZ = "Z07:00" // prints Z for UTC
stdNumTZ          = "-0700"  // always numeric
stdNumShortTZ     = "-07"    // always numeric
stdNumColonTZ     = "-07:00" // always numeric
-----------------------
	
current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

// func (t Time) Date() (year int, month Month, day int)
fmt.Println(current.Date()) // 2020 May 15

// func (t Time) Year() int
fmt.Println(current.Year()) // 2020

// func (t Time) Month() Month
fmt.Println(current.Month()) // May

// func (t Time) Day() int
fmt.Println(current.Day()) // 15

// func (t Time) Clock() (hour, min, sec int)
fmt.Println(current.Clock()) // 17 45 12

// func (t Time) Hour() int
fmt.Println(current.Hour()) //17

// func (t Time) Minute() int
fmt.Println(current.Minute()) // 45

// func (t Time) Second() int
fmt.Println(current.Second()) // 12

// func (t Time) Unix() int64
fmt.Println(current.Unix()) // 1589546712

// func (t Time) Weekday() Weekday
fmt.Println(current.Weekday()) // Friday

// func (t Time) YearDay() int
fmt.Println(current.YearDay()) // 136



-------------------
Конвертирование структуры Time в строку

// func (t Time) Format(layout string) string
current := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
fmt.Println(current.Format("02-01-2006 15:04:05")) // 15-05-2020 17:45:12
	
------------------------

Сравнение структур Time
	
firstTime := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)
secondTime := time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

// func (t Time) After(u Time) bool
// true если позже
fmt.Println(firstTime.After(secondTime)) // true

// func (t Time) Before(u Time) bool
// true если раньше
fmt.Println(firstTime.Before(secondTime)) // false

// func (t Time) Equal(u Time) bool
// true если равны
fmt.Println(firstTime.Equal(secondTime)) // false



---------------------------------------------------
	
Методы, изменяющие структуру Time
	
now := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

// func (t Time) Add(d Duration) Time
// изменяет дату в соответствии с параметром - "продолжительностью"
future := now.Add(time.Hour * 12) // перемещаемся на 12 часов вперед

// func (t Time) AddDate(years int, months int, days int) Time
// изменяет дату в соответствии с параметрами - количеством лет, месяцев и дней
past := now.AddDate(-1, -2, -3) // перемещаемся на 1 год, два месяца и 3 дня назад

// func (t Time) Sub(u Time) Duration
// вычисляет время, прошедшее между двумя датами
fmt.Println(future.Sub(past)) // 10332h0m0s


---------------------------------------------------
	
	
Задачки
	
	На стандартный ввод подается строковое представление даты и времени в следующем формате:

1986-04-16T05:20:00+06:00
Ваша задача конвертировать эту строку в Time, а затем вывести в формате UnixDate:

func main() {

	var n string
	fmt.Scan(&n)



	firstTime, err := time.Parse(time.RFC3339, n)
		if err != nil {
			panic(err)
		}
		fmt.Println(firstTime.Format(time.UnixDate))  
}



-------------------------------------------------
задача
На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:
2020-05-15 08:00:00
Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.
Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день, а затем вывести на стандартный вывод в том же формате.
	
	
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

)

func main() {

	n, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//fmt.Scan(&n)

	n = strings.TrimSuffix(n, "\n") // так прошло в задании на Lin
	//datestring3 := strings.TrimSuffix(datestring, "\r\n") на винде
	
	// var input string
	// fmt.Scan(&input)

	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, n) // парсим входные input данные согласно шаблону layout
	if err != nil {
		panic(err)
	}

	if t.Hour() >= 13 { // если больше 13го
		t = t.Add(24 * time.Hour)
	}

	fmt.Println(t.Format(layout))

}

-------------------------------------------------
	
	type Duration
		
Создается экземпляр типа Duration одной из следующих функций:

now := time.Now()
past := now.AddDate(0, 0, -1)
future := now.AddDate(0, 0, 1)

// func Since(t Time) Duration
// вычисляет период между текущим моментом и заданным временем в прошлом
fmt.Println(time.Since(past).Round(time.Second)) // 24h0m0s

// func Until(t Time) Duration
// вычисляет период между текущим моментом и заданным временем в будущем
fmt.Println(time.Until(future).Round(time.Second)) // 24h0m0s

// func ParseDuration(s string) (Duration, error)
// преобразует строку в Duration с использованием аннотаций:
// "ns" - наносекунды,
// "us" - микросекунды,
// "ms" - миллисекунды,
// "s" - секунды,
// "m" - минуты,
// "h" - часы.
dur, err := time.ParseDuration("1h12m3s")
if err != nil {
	panic(err)
}
fmt.Println(dur.Round(time.Hour).Hours()) // 1


метод Round, округляющий значение до ближайшего целого с заданной точностью.

У типа Duration помимо метода Round, который мы рассмотрели выше, есть ряд других методов, позволяющих вернуть часть значения: часы, минуты, секунды и пр.

func (d Duration) Hours() float64
func (d Duration) Minutes() float64
func (d Duration) Seconds() float64
func (d Duration) Milliseconds() int64
func (d Duration) Microseconds() int64
func (d Duration) Nanoseconds() int64
// func (t Time) Sub(u Time) Duration // вычисляет время, прошедшее между двумя датами fmt.Println(future.Sub(past)) // 10332h0m0s
	
-------------------
задачка 

На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).
Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

13.03.2018 14:00:15,12.03.2018 14:00:15

Sample Input:
13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:
24h0m0s


package main

import (
	"bufio"
	"fmt"
	"time"
	//"io"
	"os"
	//"strconv"
	"strings"

)

func main() {

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	line := input[:len(input)]

	s1, s2 := strings.Split(line, ",")[0], strings.Split(line, ",")[1]

	s2 = strings.TrimRight(s2, "\n")

	a := s2 [:len(s2)-1] // удалим один символ в конце строки. в решении на степике отработало без этого

	layout := "02.01.2006 15:04:05"

	firstTime, err := time.Parse(layout, s1)
	if err != nil {
		panic(err)
	}

	secondTime, err := time.Parse(layout, a)
	if err != nil {
		panic(err)
	}

	result := secondTime.Sub(firstTime)
	if result <0 {
	fmt.Println(firstTime.Sub(secondTime)) 
	} else {
	fmt.Println(secondTime.Sub(firstTime)) 
	}
}

-----------------------------------
	задачка
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).

Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.

Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

Sample Input:

12 мин. 13 сек.
Sample Output:

Fri May 15 19:28:18 UTC 2020
	
	
	package main

import (
	//"bufio"
	"fmt"
	"time"
	//"io"
	//"os"
	//"strconv"
	//"strings"

)

func main() {
	var min, sec int64
	fmt.Scanf("%d мин. %d", &min, &sec) // сканируем для времени и даты
	const now = 1589570165 //задано по задаче
	sumresult := (min*60 +sec ) + now // переводим минуты в секунды (60*min) складываем
	//fmt.Print(time.Unix(sumresult, 0))
	t := time.Unix(sumresult, 0)//преобразуем в тайм
	fmt.Println(t.Format(time.UnixDate)) //выводим в unixdate
}

	---лучшие решения
	package main

import (
	"fmt"
	"time"
)

func main() {
	var m, s time.Duration
	fmt.Scanf("%d мин. %d сек.", &m, &s)
	t := time.Unix(1589570165, 0).UTC().Add(m * time.Minute).Add(s * time.Second)
	fmt.Println(t.Format(time.UnixDate))
}
	
	
	
	