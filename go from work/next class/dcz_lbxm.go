for — единственная конструкция для циклов в Go. Вот несколько примеров.

С одним условием, аналог while в питоне:

i := 1
for i <= 3 {
    fmt.Println(i)
    i = i + 1
}
// 1
// 2
// 3
Классический for из трех частей (инициализация, условие, шаг):

for j := 7; j <= 9; j++ {
    fmt.Println(j)
}
// 7
// 8
// 9
Бесконечный цикл, выполняется до break для выхода из цикла или return для выхода из функции:

for {
    fmt.Println("loop")
    break
}
// loop
continue переходит к следующей итерации цикла:

for n := 0; n <= 5; n++ {
    if n%2 == 0 {
        continue
    }
    fmt.Println(n)
}
-----------------------------
	
If/else
Конструкция if-else в Go ведет себя без особых сюрпризов.

Вокруг условия не нужны круглые скобки, но фигурные скобки для веток обязательны:

if 7%2 == 0 {
    fmt.Println("7 is even")
} else {
    fmt.Println("7 is odd")
}
// 7 is odd
Можно использовать if без else:

if 8%4 == 0 {
    fmt.Println("8 is divisible by 4")
}
// 8 is divisible by 4
Единственный нюанс: перед условием может идти выражение. Объявленные в нем переменные доступны во всех ветках:

if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}
// 9 has 1 digit



-------------------------------------
	
	Switch
switch описывает условие с множеством веток.

В отличие от многих других языков, не нужно указывать break. Go выполняет только подходящую ветку и не проваливается в следующую:

i := 2
fmt.Print("Write ", i, " as ")
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
    fmt.Println("three")
}
// Write 2 as two
Чтобы провалиться в следующую ветку, есть специальное ключевое слово fallthrough. Его можно использовать только в качестве последней инструкции в блоке:

i := 2
fmt.Print("Write ", i, " as ")
switch i {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
    fallthrough
case 3:
    fmt.Println("Bye-bye")
}
// Write 2 as two
// Bye-bye
В одной ветке можно указать несколько выражений. Ветка default сработает, если остальные не подошли:

// одной строкой инициализировали day
// и сразу сделали switch по нему
switch day := time.Now().Weekday(); day {
case time.Saturday, time.Sunday:
    fmt.Println(day, "is a weekend")
default:
    fmt.Println(day, "is a weekday")
}
// Tuesday is a weekday
Выражения в ветках не обязательно должны быть константами. switch может работать как if:

t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("It's before noon")
default:
    fmt.Println("It's after noon")
}
// It's before noon

-------------------------
	
	Массивы
Массив в Go — это нумерованная последовательность элементов. Длина массива известна заранее и зафиксирована.

Массив a содержит 5 целых чисел. Тип и количество элементов — часть определения массива. По умолчанию элементы массива принимают нулевое значение, в данном случае — 0:

var arr [5]int
fmt.Println("empty:", arr)
// empty: [0 0 0 0 0]
Обращение к элементу массива — через квадратные скобки:

arr[4] = 100
fmt.Println("set:", arr)
// set: [0 0 0 0 100]
fmt.Println("get:", arr[4])
// get: 100
Встроенная функция len() возвращает количество элементов:

fmt.Println("len:", len(arr))
// len: 5
Можно инициализировать массив при объявлении:

arr := [5]int{1, 2, 3, 4, 5}
fmt.Println("init:", arr)
// init: [1 2 3 4 5]
Массивы одноразмерные, но их можно комбинировать, чтобы получить нужную размерность:

var arr [2][3]int
for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
        arr[i][j] = i + j
    }
}
fmt.Println("2d:", arr)
// 2d: [[0 1 2] [1 2 3]]


--------------------------------------
	Срезы
Срез (slice) — ключевая структура данных в Go. Это массив изменяемой длины, как list в питоне или Array в js. Обычно в программах на Go оперируют именно срезами, «чистые» массивы встречаются намного реже.

Срез определяется только типом элементов, но не их количеством. Чтобы создать срез ненулевой длины, используют встроенную функцию make(). Здесь мы создаем срез из трех пустых строк:

s := make([]string, 3)
fmt.Printf("empty: %#v\n", s)
// empty: []string{"", "", ""}
Шаблон %#v возвращает «внутреннее представление» значения, примерно как repr() в питоне.

s := make([]string, 3)
fmt.Printf("%#v", s)
	



С элементами среза можно работать точно так же, как с элементами массива:

s[0] = "a"
s[1] = "b"
s[2] = "c"

fmt.Println("set:", s)
// set: [a b c]

fmt.Println("get:", s[2])
// get: c
len() возвращает длину среза:

fmt.Println("len:", len(s))
// len: 3
Срез можно инициализировать при объявлении:

s := []string{"a", "b", "c"}
fmt.Println("init:", s)
// init: [a b c]
В отличие от массива, в срез можно добавлять новые элементы через встроенную функцию append(). Функция возвращает новый срез:

s := []string{"a", "b", "c"}

fmt.Println("src:", s)
// src: [a b c]

s = append(s, "d")
s = append(s, "e", "f")
fmt.Println("upd:", s)
// upd: [a b c d e f]
Всегда используйте значение, которое возвращает append(). Вот так делать не стоит:

append(s, "d")
fmt.Println("upd:", s)
Дело в том, что срез сам по себе не хранит данные, это ссылка на конкретный массив. Если в массиве нет места для нового элемента, append() создаст новый массив побольше, скопирует в него старые элементы, добавит новый элемент и вернет ссылку на новый массив. Если эту ссылку проигнорировать, новый срез вы потеряете.

Срез можно скопировать через встроенную функцию copy(). Здесь создаем пустой срез dst такой же длины, как s, и копируем в него элементы s:

src := []string{"a", "b", "c", "d", "e", "f"}
dst := make([]string, len(src))

copy(dst, src)
fmt.Println("copy:", dst)
// copy: [a b c d e f]
Срезы поддерживают... срезы (отсюда их название). Выражение slice[from:to] вернет срез от элемента с индексом from включительно до элемента с индексом to не включительно:

s := []string{"a", "b", "c", "d", "e", "f"}

sl1 := s[2:5]
fmt.Println("sl1:", sl1)
// sl1: [c d e]
Этот срез включает все элементы, кроме s[5]:

sl2 := s[:5]
fmt.Println("sl2:", sl2)
// sl2: [a b c d e]
А этот срез включает элементы от s[2] и до конца:

sl3 := s[2:]
fmt.Println("sl3:", sl3)
// sl3: [c d e f]

-----
	Срез можно получить из массива:

a := [...]int{0, 1, 2, 3}
s := a[1:3]
// [1 2]
Можно создать срез нужной capacity без заполнения значениями по умолчанию:

s := make([]int, 0, 5)
fmt.Println(s, len(s), cap(s))
// [] 0 5
Все срезы массива и дочерние срезы ссылаются на одни и те же данные: 

package main

import "fmt"

func main() {
    a := [...]int{0, 1, 2, 3, 4, 5}
    s1 := a[1:]
    s2 := s1[1:3]

    fmt.Println(a, s1, s2)
    // [0 1 2 3 4 5] [1 2 3 4 5] [2 3]

    s2[0] = 8

    fmt.Println(a, s1, s2)
    // [0 1 8 3 4 5] [1 8 3 4 5] [8 3]
}
-----
	package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	a2 := [4]int{1, 2, 3}
	fmt.Println("Массивы 1:", a1, a2)
	fmt.Println("Типы:", reflect.TypeOf(a1), reflect.TypeOf(a2))              // Посмотрим типы
	fmt.Println("Типы Одинаковые?", reflect.TypeOf(a1) == reflect.TypeOf(a2)) // Убедимся, что типы разные
	/* fmt.Println("Массивы одинаковые?", a1 == a2) 			  // Свалится с ошибкой из-за разных типов
	// если бы вдруг заработало, то, наверное, было бы false*/
	fmt.Println("")

	b1 := [3]int{1, 2, 3}
	b2 := [3]int{1, 2, 3}
	fmt.Println("Массивы 2:", b1, b2)
	fmt.Println("Типы:", reflect.TypeOf(b1), reflect.TypeOf(b2))
	fmt.Println("Типы Одинаковые?", reflect.TypeOf(b1) == reflect.TypeOf(b2))
	fmt.Println("Массивы одинаковые?", b1 == b2) // А при одинаковых размерах работает

}

Массивы 1: [1 2 3] [1 2 3 0]
Типы: [3]int [4]int
Типы Одинаковые? false

Массивы 2: [1 2 3] [1 2 3]
Типы: [3]int [3]int
Типы Одинаковые? true
Массивы одинаковые? true

Program exited.
--------
	
	
	
package main

import "fmt"

func main() {

	slice := make([]int, 0)
	slice1 := []int{}
	var slice2 []int
	// Выводим тип, представление, нулёвость и длину
	fmt.Printf("%T\t%v\t%v\t%v\n", slice, slice, slice == nil, len(slice))
	fmt.Printf("%T\t%v\t%v\t%v\n", slice1, slice1, slice1 == nil, len(slice))
	fmt.Printf("%T\t%v\t%v\t%v\n", slice2, slice2, slice2 == nil, len(slice))
	slice = append(slice, 1)
	slice1 = append(slice1, 1)
	slice2 = append(slice2, 1)
	fmt.Printf("%T\t%v\t%v\t%v\n", slice, slice, slice == nil, len(slice))
	fmt.Printf("%T\t%v\t%v\t%v\n", slice1, slice1, slice1 == nil, len(slice))
	fmt.Printf("%T\t%v\t%v\t%v\n", slice2, slice2, slice2 == nil, len(slice))
}



[]int	[]	false	0
[]int	[]	false	0
[]int	[]	true	0
[]int	[1]	false	1
[]int	[1]	false	1
[]int	[1]	false	1

Program exited.
	
----------------------------------------------

Срезы и строки
Строку можно преобразовать в срез байт и обратно:

str := "го!"
bytes := []byte(str)

fmt.Println(bytes)
// [208 179 208 190 33]

fmt.Println(str == string(bytes))
// true
Строку можно преобразовать в срез unicode-символов (Go называет их рунами). Одна руна может занимать несколько байт (что и произошло с рунами г и о):

runes := []rune(str)

fmt.Println(runes)
// [1075 1086 33]

fmt.Println(str == string(runes))
// true
Внутри Go строка реализована как массив байт, а не рун. Так что обращение по индексу к элементу строки вернет соответствующий байт, а не руну:

str := "го!"

bytes := []byte(str)
fmt.Println(bytes[1])
// 179 - второй байт

runes := []rune(str)
fmt.Println(runes[1])
// 1086 - вторая руна

fmt.Println(str[1])
// 179, не 10864




---------------------------------
	Карты
Карта (map), так же известная как словарь (dict), хеш-таблица (hash table) или ассоциативный массив (associative array) — это неупорядоченный набор пар «ключ-значение».

Чтобы создать пустую карту, используют make():

m := make(map[string]int)
Задать пары «ключ-значение»:

m["key"] = 7
m["other"] = 13
Вывести содержимое карты:

fmt.Println("map:", m)
// map: map[key:7 other:13]
Получить значение по ключу:

val := m["key"]
fmt.Println("val:", val)
// val: 7
len() возвращает количество записей (пар «ключ-значение») в карте:

fmt.Println("len:", len(m))
// len: 2
delete() удаляет запись по ключу:

delete(m, "other")
fmt.Println("map:", m)
// map: map[key:7]
Обращение к записи по ключу возвращает необязательное второе значение: признак, есть такой ключ в карте или нет. Обращение по несуществующему ключу не приведет к ошибке, но вернет этот признак со значением false:

_, ok := m["other"]
fmt.Println("has other:", ok)
// has other: false
Пустой идентификатор _ указывает, что нам не интересно само значение по ключу, важен только признак «есть/нет» (ok).

Карту можно инициализировать при объявлении:

n := map[string]int{"foo": 1, "bar": 2}
fmt.Println("map:", n)
// map: map[bar:2 foo:1]


----------------------

Напишите программу, которая считает, сколько раз каждая цифра встречается в числе. Гарантируется, что на вход подаются только положительные целые числа, не выходящие за диапазон int.

Sample Input:

12823
Sample Output:

1:1 2:2 3:1 8:1



package main

import (
	"fmt"
	"strconv"
	"strings"

)

func main() {
	var number int
	fmt.Scan(&number)

	counter := make(map[int]int)

	var s string = strconv.FormatUint(uint64(number), 10)

	for i := 0; i < len(s); i++ {
		r,_ := strconv.Atoi(string(s[i]))
		counter[r] = strings.Count(s, string(s[i]))
	}

	fmt.Print(counter)

	
}


----------------------------------------------------

Обход коллекции
range обходит элементы коллекций. Посмотрим, как использовать его на срезах, картах и строках.

Просуммировать элементы среза (или массива):

nums := []int{2, 3, 4}
sum := 0
for _, num := range nums {
    sum += num
}
fmt.Println("sum:", sum)
// sum: 9
range на массивах и срезах возвращает индекс и значение для каждого элемента. В примере выше мы не использовали индекс, поэтому заглушили его пустым идентификатором _. Но иногда индекс может и пригодиться:

nums := []int{2, 3, 4}
for idx, num := range nums {
    if num == 3 {
        fmt.Println("index:", idx)
    }
}
// index: 1
range на карте итерирует по записям:

m := map[string]string{"a": "apple", "b": "banana"}
for key, val := range m {
    fmt.Printf("%s -> %s\n", key, val)
}
// a -> apple
// b -> banana
Или только по ключам:

m := map[string]string{"a": "apple", "b": "banana"}
for key := range m {
    fmt.Println("key:", key)
}
// key: a
// key: b
range на строках итерирует по unicode-символам (рунам). Первое значение — порядковый номер байта, с которого начинается руна (руна может занимать несколько байт). Второе значение — числовой код самой руны:

for idx, char := range "ого" {
    fmt.Println(idx, char, string(char))
}
// 0 1086 о
// 2 1075 г
// 4 1086 о




-----------------------------------
Напишите программу, которая принимает на вход фразу и составляет аббревиатуру по первым буквам слов:

Today I learned → TIL
Высшее учебное заведение → ВУЗ
Кот обладает талантом → КОТ
Если слово начинается не с буквы, игнорируйте его:

Ар 2 Ди #2 → АД
Разделителями слов считаются только пробельные символы. Дефис, дробь и прочие можно не учитывать:

Анна-Мария Волхонская → АВ
	
	
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	// "strconv"
	"strings"

)

func main() {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	
	t := strings.Fields(str)
	
	abbr := []rune{}

	for _, value := range t {
		rs := []rune(value)
		if unicode.IsLetter(rs[0]){
		abbr = append(abbr,  unicode.ToUpper(rs[0]))
		}
	}
	fmt.Print(string(abbr))
}



------------------------------


Функции
Функция — центральная конструкция языка. Вот несколько примеров.

Эта функция принимает два целых числа и возвращает их сумму:

func sum(a int, b int) int {
    return a + b
}
Результат возвращается явно, через return.

Если у нескольких идущих подряд параметров один и тот же тип, их можно «схлопнуть» и указать тип только для последнего:

func sum(a, b, c int) int {
    return a + b + c
}
Функция может возвращать несколько значений. В этом примере — частное и остаток от деления одного числа на другое:

func divide(divisible, divisor int) (int, int) {
    quotient := divisible / divisor
    remainder := divisible % divisor
    return quotient, remainder
}
Результат вызова можно сразу разложить по переменным:

q, r := divide(10, 3)
fmt.Println("10 / 3 =", q)
// 10 / 3 = 3
fmt.Println("10 % 3 =", r)
// 10 % 3 = 1
Или проигнорировать одно из значений с помощью пустого идентификатора _:

_, r = divide(42, 2)
if r == 0 {
    fmt.Println("42 is divisible by 2")
}
// 42 is divisible by 2

------------------------------
	
	Вариативные функции
Функция может принимать произвольное количество аргументов в «хвосте» (как *args в питоне или ...args в js). Например, так ведет себя fmt.Println(). В Go такие функции называют вариативными (variadic).

Эта функция суммирует целые числа:

func sum(nums ...int) {
    fmt.Print(nums, " -> ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
Можно передавать индивидуальные аргументы, как у обычной функции:

sum(1, 2)
// [1 2] -> 3
sum(1, 2, 3)
// [1 2 3] -> 6
А можно передать срез, преобразовав его в список аргументов с помощью ...:

nums := []int{1, 2, 3, 4}
sum(nums...)
// [1 2 3 4] -> 10

-------------------------------------------
	
Анонимные функции
Go поддерживает анонимные функции. Работают они как обычные, но не имеют названия (как лямбды в питоне или стрелочные функции в js).

Чаще всего анонимные функции используют, чтобы вернуть из функции другую функцию. В примере ниже intSeq() возвращает функцию-генератор, которая при каждом вызове выдает очередное значение счетчика i. Генератор использует переменную, определенную во внешней функции — то есть образует замыкание (closure):

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
Результат вызова intSeq() — функцию-генератор — мы записываем в переменную next. У next собственное значение счетчика i, которое увеличивается при каждом вызове:

next := intSeq()

fmt.Println(next())
// 1
fmt.Println(next())
// 2
fmt.Println(next())
// 3
Если создать еще один генератор — он будет обладать собственным счетчиком i:

gen := intSeq()
fmt.Println(gen())
// 1
fmt.Println(gen())
// 2
Иногда анонимную функцию передают как аргумент другой функции. Пример из пакета sort :

func Search(n int, f func(int) bool) int
Search() находит наименьшее i из диапазона [0, n), для которого функция-предикат f(i) вернет true. В качестве предиката удобно использовать анонимную функцию:

a := []int{1, 2, 4, 8, 16, 32, 64, 128}
x := 53

// ближайший сверху к `x` элемент среза `a`
closest := sort.Search(len(a), func(i int) bool { return a[i] >= x })

fmt.Println(a[closest], "is the closest to", x)
// 64 is the closest to 53




-------------------------------------------
Напишите функцию filter(), которая фильтрует срез целых чисел с помощью функции-предиката и возвращает отфильтрованный срез. Функция-предикат вызывается для каждого элемента исходного среза. Если она возвращает true, элемент попадает в отфильтрованный срез. Если возвращает false — не попадает.
Считайте исходный срез из стандартного ввода с помощью готовой функции readInput(). Затем выполните на нем filter(). В качестве предиката используйте функцию, которая возвращает true только для четных чисел. Напечатайте отфильтрованный срез.
Гарантируется, что на вход подаются только целые числа.
	
	package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

)

func filter(predicate func(int) bool, iterable []int) []int {
	// отфильтруйте `iterable` с помощью `predicate`
	// и верните отфильтрованный срез
	nums := []int{}

	iscount := func(i int) bool {
		return i%2 ==0
	}

	for i := 0; i < len(iterable); i++ {
		if iscount(iterable[i]) {
			nums = append(nums, iterable[i])
		}
	}
	return nums
}

func main() {
	src := readInput()
	res := filter(func(i int) bool {return i%2 ==0}, src)
	fmt.Println(res)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}

	
	
-------------------------------------------
Указатели
Указатель (pointer) содержит адрес памяти, который ссылается на конкретное значение.

Тип *T — указатель на значение типа T. Если указатель не инициализирован, он равен nil (аналог None в питоне и null в js).

// iptr - указатель на значение типа int
// пока что он пустой
var iptr *int
fmt.Println(iptr)
// <nil>
Оператор & возвращает указатель на конкретное значение:

i := 42
iptr = &i

// теперь iptr ссылается на i
fmt.Println(iptr)
// 0xc000118000
0xc000118000 — адрес памяти. По этому адресу находится значение 42.

Оператор * обращается к значению, на которое ссылается указатель. Оно доступно как для чтения, так и для записи:

// прочитать значение i через указатель iptr
fmt.Println(*iptr)
// 42

// установить значение i через указатель iptr
*iptr = 21
fmt.Println(i)
// 21
Вот схема непростых отношений между значением и указателем на него:

  &i                 i
┌──────────────┐   ┌──────────┐
│ 0xc000118000 │ → │ 42       │
└──────────────┘   └──────────┘
  *int               int
&i — указатель на i. Его тип — *int, а значение — 0xc00011800. По адресу 0xc000118000 находится значение i — 42.

  iptr = &i          *iptr
┌──────────────┐   ┌──────────┐
│ 0xc000118000 │ → │ 42       │
└──────────────┘   └──────────┘
iptr — указатель на i. Вызывая *iptr, мы обращаемся по адресу 0xc000118000 и получаем значение i — 42.

Остаток урока мы посвятим указателям. Если они вам «не зашли», можно перейти к следующему уроку и вернуться сюда, когда будет подходящий настрой.
	
	
	
----------
	Указатели в параметрах функции
	
Указатели в параметрах функции позволяют изменять переданные значения. Вот как это работает.

Функция addval() принимает параметр типа int — конкретное число. При вызове Go передает не оригинальное число n, а его копию — nval:

func addval(nval int, delta int) {
    nval += delta
}

n := 42
addval(n, 3)
fmt.Println(n)
// 42
addval() модифицировала копию оригинального числа nval, так что само n не изменилось.

Функция addptr() принимает параметр типа *int — указатель на число. При вызове Go передает адрес в памяти, по которому находится оригинальное число n — указатель nptr. Функция изменяет значение n по указателю через оператор *:

func addptr(nptr *int, delta int) {
    *nptr += delta
}

n := 42
addptr(&n, 3)
fmt.Println(n)
// 45
Благодаря оператору &, в функцию передано не значение n, а указатель на него. addptr() изменила оригинальное значение n через переданный указатель.
	
-------------
	
Указатель или значение?
Базовый принцип такой:

Если функция только читает переменную, но не изменяет — передавайте значение.
Если функция изменяет значение — передавайте указатель.
Функция math.Max() возвращает максимальное из двух чисел:

a := 5
b := 3
max := math.Max(a, b)
a и b не изменяются, поэтому функция принимает обычные значения.

Функция fmt.Scanf() считывает значения из стандартного ввода и записывает их в переданные переменные. Поэтому принимает указатели:

var a, b int
fmt.Scanf("%d-%d", &a, &b)
Правило работает для скалярных значений (логических, чисел, строк) и массивов. Со срезами и картами другая история.
	
	
----------
	

	Указатель или значение? Срезы и карты
Функция стандартной библиотеки sort.Ints() сортирует срез:

func Ints(nums []int)
Обратите внимание — функция ничего не возвращает, она изменяет элементы оригинального среза. Но почему тогда nums передан как значение ([]int), а не как указатель (*[]int)?

Дело в том, что срез сам по себе не содержит данные массива. Срез — это легковесная структура данных, одно из полей которой — указатель на конкретный массив. Поэтому nums внутри функции — это копия, но не всего массива, а этой легковесной структуры с указателем на массив. Обращаясь к элементу среза, функция переходит по указателю и модифицирует оригинальный элемент массива.

  slice             array
┌──────────────┐   ┌─┬─┬───┬─┐
│ 0xc000118000 │ → │0│1│...│n│
└──────────────┘   └─┴─┴───┴─┘
Если функция изменяет отдельные элементы среза, передавайте его как значение:

func sortSlice(nums []int) {
    sort.Ints(nums)
}

nums := []int{5, 1, 3, 9}
sortSlice(nums)
fmt.Println(nums)
// [1 3 5 9]
Этот подход не сработает, если изменить сам срез (добавить или удалить элементы):

func appendByVal(nums []int, n int) {
    nums = append(nums, n)
}

nums := []int{42}
appendByVal(nums, 43)
fmt.Println(nums)
// ожидание: [42 43]
// реальность: [42] 
nums внутри функции — это копия оригинального среза. Изменив nums через append(), функция поменяла копию, а оригинал не изменился.

Чтобы изменить срез в целом, можно использовать указатель:
func appendByPtr(nums *[]int, n int) {
    *nums = append(*nums, n)
}

nums := []int{42}
appendByPtr(&nums, 43)
fmt.Println(nums)
// [42 43]
Но в обычных функциях такой подход нечасто встретишь. Лучше вернуть новый срез, чем переопределять старый по указателю:

func appendAndReturn(nums []int, n int) []int {
    nums = append(nums, n)
    return nums
}

nums := []int{42}
nums = appendAndReturn(nums, 43)
fmt.Println(nums)
// [42 43]
Итого:

если функция не меняет срез — передавать значение;
если функция меняет отдельные элементы, но не сам срез — передавать значение;
если функция меняет сам срез — передавать значение и возвращать новое значение.
Вариант «передавать указатель на срез» остается для методов, о которых мы поговорим на следующем уроке.

Для карт принцип такой же.
	
	
	
-------------------
	// shuffle перемешивает элементы nums in-place.
func shuffle(nums []int) {
    rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	//fmt.Println(nums)
}
	
	
------------------
	Структуры	
------------------
	
	Структуры
Структура (struct) группирует поля в единую запись. В Go нет классов и объектов, так что структура — наиболее близкий аналог объекта в питоне и js.

Объявим тип person на основе структуры с полями name и age:

type person struct {
    name string
    age  int
}
Так создается новая структура типа person:

bob := person{"Bob", 20}
fmt.Println(bob)
// {Bob 20}
Можно явно указать названия полей:

alice := person{name: "Alice", age: 30}
fmt.Println(alice)
// {Alice 30}
Если не указать поле, оно получит нулевое значение:

fred := person{name: "Fred"}
fmt.Println(fred)
// {Fred 0}
Оператор & возвращает указатель на структуру:

annptr := &person{name: "Ann", age: 40}
fmt.Println(annptr)
// &{Ann 40}
В Go иногда создают новые структуры через функцию-конструктор с префиксом new:

func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}
Функция возвращает указатель на локальную переменную — это нормально. Go распознает такие ситуации, и выделяет память под структуру в куче (heap) вместо стека (stack), так что структура продолжит существовать после выхода из функции.

john := newPerson("John")
fmt.Println(john)
// &{John 42}
Если функция-конструктор возвращает саму структуру, а не указатель — удобно использовать префикс make вместо new:

func makePerson(name string) person {
    p := person{name: name}
    p.age = 42
    return p
}
В реальности чаще не заморачиваются и всегда используют префикс new вне зависимости от того, что возвращает конструктор — значение или указатель на него. Но на курсе я буду соблюдать это разделение: make — значение, new — указатель.

Доступ к полям структуры — через точку:

sean := person{name: "Sean", age: 50}
fmt.Println(sean.name)
// Sean
Чтобы получить доступ к полям структуры через указатель, не обязательно разыменовывать его через *. Эти два варианта эквивалентны:

sven := &person{name: "Sven", age: 50}
fmt.Println((*sven).age)
// 50
fmt.Println(sven.age)
// 50
Поля структуры можно изменять:

sven.age = 51
fmt.Println(sven.age)
// 51


-----------------------------
	
если использовать значение — будет создана копия оригинальной структуры. 
Если указатель — копии не будет, будет ссылка на оригинал. 
	
	
Составные структуры
Структуры могут включать другие структуры:

type person struct {
    firstName string
    lastName  string
}

type book struct {
    title  string
    author person
}
b := book{
    title: "The Majik Gopher",
    author: person{
        firstName: "Christopher",
        lastName:  "Swanson",
    },
}
fmt.Println(b)
// {The Majik Gopher {Christopher Swanson}}
Если вложенная структура не представляет самостоятельной ценности, можно даже не объявлять отдельный тип:

type user struct {
    name  string
    karma struct {
        value int
        title string
    }
}
u := user{
    name: "Chris",
    karma: struct {
        value int
        title string
    }{
        value: 100,
        title: "^-^",
    },
}
fmt.Printf("%+v\n", u)
// {name:Chris karma:{value:100 title:^-^}}
Благодаря шаблону %+v, Printf() печатает структуру вместе с названиями полей.

Поле структуры может ссылаться на другую структуру:

type comment struct {
    text   string
    author *user
}
chris := user{
    name: "Chris",
}
c := comment{
    text:   "Gophers are awesome!",
    author: &chris,
}
fmt.Printf("%+v\n", c)
// {text:Gophers are awesome! author:0xc0000981e0}

	
	

----------------------------
	Методы
	
	Go позволяет определять методы на типах.

Метод отличается от обычной функции специальным параметром — получателем. В определении метода получатель указывается сразу после ключевого слова func. В данном случае — получатель типа rect:

type rect struct {
    width, height int
}

func (r rect) area() int {
    return r.width * r.height
}
Метод вызывается для получателя через точку, как в питоне или js:

r := rect{width: 10, height: 5}
fmt.Println("rect area:", r.area())
// rect area: 50
Получателем может быть не значение заданного типа, а указатель на это значение:

type circle struct {
    radius int
}

func (c *circle) area() float64 {
    return math.Pi * math.Pow(float64(c.radius), 2)
}
cptr := &circle{radius: 5}
fmt.Println("circle area:", cptr.area())
// circle area: 78.54
При вызове метода Go автоматически преобразует значение получателя в указатель или указатель в значение, как того требует определение метода. Любой из перечисленных вариантов будет работать:

rptr := &r
r.area()
rptr.area()

c := *cptr
c.area()
cptr.area()
	!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
Считается хорошим тоном во всех методах использовать или только значение, или только указатель, но не смешивать одно с другим. Обычно используют указатель: так Go не приходится копировать всю структуру, а метод может ее изменить.

// Если метод принимает получателя как значение, изменить его не получится
func (r rect) scale(factor int) {
    r.width *= factor
    r.height *= factor
}

fmt.Println("rect before scaling:", r)
// rect before scaling: {10 5}

r.scale(2)

fmt.Println("rect after scaling:", r)
// rect after scaling: {10 5}

// Если метод принимает получателя как указатель, его можно изменить
func (c *circle) scale(factor int) {
    c.radius *= factor
}

fmt.Println("circle before scaling:", c)
// circle before scaling: {5}

c.scale(2)

fmt.Println("circle after scaling:", c)
// circle after scaling: {10}
Вопрос «значение или указатель для получателя» так популярен, что у сообщества есть отдельный гайдлайн на эту тему.
	!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	-----------------	
	
	
	
	Определяемые типы
На предыдущем шаге мы создали структурный тип и методы для него. Но новый тип не обязательно создавать на основе структуры — можно использовать любые базовые типы.

Создадим тип «ИНН» на основе строки:

type inn string
Тип inn (он называется определяемым типом, defined type) получил свойства базового типа string. Добавим ему новое поведение с помощью метода:

func (id inn) isValid() bool {
    if len(id) != 12 {
        return false
    }
    for _, char := range id {
        if !unicode.IsDigit(char) {
            return false
        }
    }
    return true
}
inn1 := inn("111201284667")
fmt.Println("inn", inn1, "is valid:", inn1.isValid())
// inn 111201284667 is valid: true

inn2 := inn("ohmyinn12345")
fmt.Println("inn", inn2, "is valid:", inn2.isValid())
// inn ohmyinn12345 is valid: false
Это чем-то похоже на наследование, но механизм более примитивный. Если создать новый определяемый тип на основе inn — он унаследует структуру и свойства inn, но не методы:

type otherid inn
other := otherid("111201284667")
fmt.Println("other inn", other, "is valid:", other.isValid())
// ОШИБКА: other.isValid undefined

----------------Композиция--------------
	
	
package main

import (
	"fmt"
	"log"
	"net/url"

)

// counter представляет целочисленный счетчик
type counter struct {
	value uint
}

// increment увеличивает счетчик на единицу
func (c *counter) increment() {
	c.value++
}

// incrementDelta увеличивает счетчик на дельту
func (c *counter) incrementDelta(delta uint) {
	c.value += delta
}

// usage представляет статистику использования сервиса
type usage struct {
	service string
	counter counter
}

// makeUsage создает usage
func makeUsage(service string) usage {
	return usage{service, counter{}}
}

// pageviews представляет просмотры страницы
type pageviews struct {
	url     *url.URL
	counter counter
}

// makePageviews создает pageviews
func makePageviews(uri string) pageviews {
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal(err)
	}
	return pageviews{u, counter{}}
}

func main() {
	usage := makeUsage("find")
	usage.counter.increment()
	usage.counter.increment()
	usage.counter.increment()
	usage.counter.increment()
	fmt.Printf("%s usage: %d\n", usage.service, usage.counter.value)
	// find usage: 3

	pv := makePageviews("/doc/find")
	pv.counter.incrementDelta(100)
	fmt.Printf("%s views: %d\n", pv.url, pv.counter.value)
	// /doc/find views: 100
}

	
//---------------Встраивание---------------
	
//Все хорошо, но несколько неудобно было писать usage.counter.increment() на предыдущем шаге. По-хорошему, usage  расширяет counter — //отношение между ними больше похоже на наследование, чем на композицию. В Go в таких случаях используют встраивание (embedding). //Посмотрим, как оно работает.

//Есть тип «счетчик», такой же, как на предыдущем шаге:

type counter struct {
    value uint
}
func (c *counter) increment() {
    c.value++
}
func (c *counter) incrementDelta(delta uint) {
    c.value += delta
}
//Мы хотим замерять использование сервисов. Встроим счетчик в тип «использование сервиса»:

type usage struct {
    service string
    counter
}
func makeUsage(service string) usage {
    return usage{service, counter{}}
}
//Благодаря встраиванию, поля и методы счетчика доступны прямо на usage, без обращения к полю counter:

usage := makeUsage("find")
usage.increment()
usage.increment()
usage.increment()
fmt.Printf("%s usage: %d\n", usage.service, usage.value)
// find usage: 3
//Аналогично с типом «просмотры страниц»:

type pageviews struct {
    url *url.URL
    counter
}

func makePageviews(uri string) pageviews {
    u, err := url.Parse(uri)
    if err != nil {
        log.Fatal(err)
    }
    return pageviews{u, counter{}}
}
//Поля и методы счетчика доступны прямо на pageviews:

pv := makePageviews("/doc/find")
pv.incrementDelta(100)
fmt.Printf("%s views: %d\n", pv.url, pv.value)
///doc/find views: 100



----------------------Задача-----------------------
Напишите программу, которая проверяет корректность пароля. Корректным считается пароль, который удовлетворяет любому из условий:

содержит буквы и цифры;
длина не менее 10 символов.
Следуйте указаниям по тексту программы. Не меняйте сигнатуры функций, определение типа password и переменную validator в main().

Гарантируется, что на вход программы подается строка без пробелов.

Sample Input:
hellowor1d
Sample Output:
true
	
	
	
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"

)

// validator проверяет строку на соответствие некоторому условию
// и возвращает результат проверки
type validator func(s string) bool

// digits возвращает true, если s содержит хотя бы одну цифру
// согласно unicode.IsDigit(), иначе false
func digits(s string) bool {
	f := false
	for _, v := range s {
		if unicode.IsDigit(v) {
			f = true
			return f
		}
	}
	//fmt.Println(f,"--")
	return f
	
}

// letters возвращает true, если s содержит хотя бы одну букву
// согласно unicode.IsLetter(), иначе false
func letters(s string) bool {
	f := false
	for _, v := range s {
		if unicode.IsLetter(v) {
			f = true
			return f
		}
	}
	//fmt.Println(f, "----")
	return f
}

// minlen возвращает валидатор, который проверяет, что длина
// строки согласно utf8.RuneCountInString() - не меньше указанной
func minlen(length int) validator {
	f := func(s string) bool {
		return utf8.RuneCountInString(s) >= length 
	}
	return f
}

// and возвращает валидатор, который проверяет, что все
// переданные ему валидаторы вернули true
func and(funcs ...validator) validator {
	f := true
	return func(s string) bool {
		for _, v := range funcs {
			if !v(s){
			f = v(s)
			}
		}
		return f
	}

}

// or возвращает валидатор, который проверяет, что хотя бы один
// переданный ему валидатор вернул true
func or(funcs ...validator) validator {
	f := false
	return func(s string) bool {
		for _, v := range funcs {
			if v(s){
			f = v(s)
			}
		}
		return f
	}
}

// password содержит строку со значением пароля и валидатор
type password struct {
	value string
	validator
}

// isValid() проверяет, что пароль корректный, согласно
// заданному для пароля валидатору
func (p *password) isValid() bool {
	if p.validator(p.value) == true {
		//fmt.Print("true")
		return true
	}
	//fmt.Print("false")
	return false
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

func main() {
	var s string
	fmt.Scan(&s)
	// валидатор, который проверяет, что пароль содержит буквы и цифры,
	// либо его длина не менее 10 символов
	validator := or(and(digits, letters), minlen(10))
	p := password{s, validator}
	fmt.Println(p.isValid())
}
------------------------------------------
ИНТЕРФЕЙСЫ
------------------------------------------
	
	
//Интерфейс в Go — это набор сигнатур методов (то есть список методов без реализации).

//Интерфейс геометрической фигуры:

type geometry interface {
    area() float64
    perim() float64
}
//Реализуем интерфейс в типе «прямоугольник». Реализовать интерфейс = реализовать его методы. Действует «утиный» принцип, как в питоне: если у типа есть перечисленные в интерфейсе методы — значит, он реализовал интерфейс. Явно указывать, что rect реализует geometry, не требуется:

type rect struct {
    width, height float64
}

func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
//Аналогично для типа «круг»:

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}
//Если у переменной интерфейсный тип, она поддерживает все методы, заданные на интерфейсе. Благодаря этому функция measure() работает с любой фигурой, реализующей интерфейс geometry:

func measure(g geometry) {
    fmt.Printf("%T: %+v\n", g, g)
    fmt.Println("area:", g.area())
    fmt.Println("perimiter:", g.perim())
}
//Раз типы circle и rect реализуют интерфейс geometry, мы можем передать их экземпляры в функцию measure():

r := rect{width: 3, height: 4}
c := circle{radius: 5}

measure(r)
// main.rect: {width:3 height:4}
// area: 12
// perimiter: 14

measure(c)
// main.circle: {radius:5}
// area: 78.53981633974483
// perimiter: 31.41592653589793


-----------------------------------------------------------


//Встраивание интерфейса
//Иногда при композиции хочется дать доступ к поведению, но скрыть структуру. В этом поможет встраивание интерфейса (interface embedding).

Есть тип «счетчик»:

type counter struct {
    val uint
}
func (c *counter) increment() {
    c.val++
}
func (c *counter) value() uint {
    return c.val
}
//Мы хотим встраивать счетчик в другие типы, но не давать прямой доступ к полю val — чтобы менять значение счетчика можно было только через методы.

//Определим интерфейс счетчика:

type Counter interface {
    increment()
    value() uint
}
//И вместо конкретного типа counter встроим интерфейс Counter, который он реализует:

type usage struct {
    service string
    Counter
}
//В конструкторе будем создавать конкретное значение типа counter, но потребителям об этом знать не обязательно:

func newUsage(service string) *usage {
    return &usage{service, &counter{}}
}
//Поскольку мы встроили интерфейс, прямого доступа к counter.val больше нет. Можно использовать только методы интерфейса:

usage := newUsage("find")
usage.increment()
usage.increment()
usage.increment()
fmt.Printf("%s usage: %d\n", usage.service, usage.value())
// find usage: 3
----------
	
	
package main

import (
	"fmt"
)

// counter представляет целочисленный счетчик
type counter struct {
	val uint
}

// increment увеличивает счетчик на единицу
func (c *counter) increment() {
	c.val++
}

// value возвращает значение счетчика
func (c *counter) value() uint {
	return c.val
}

// Counter представляет счетчик
type Counter interface {
	increment()
	value() uint
}

// usage представляет использование сервиса
type usage struct {
	service string
	Counter
}

// newUsage создает usage
func newUsage(service string) *usage {
	return &usage{service, &counter{}}
}

func main() {
	usage := newUsage("find")
	usage.increment()
	usage.increment()
	usage.increment()
	fmt.Printf("%s usage: %d\n", usage.service, usage.value())
	// find usage: 3
}

?????
	
	Получается так? Структура1 , в своем составе имеет интерфейс1 и при создании экземпляра этой структуры1 мы можем использовать методы дургой структуры2 которая реализует интерфейс 1. 

-----------------------------------------------------------

Пустой интерфейс
	
Если у интерфейса нет методов, его называют пустым (empty):

interface{}
Пустой интерфейс может содержать значение любого типа (ведь у каждого типа есть как минимум 0 методов). Пустые интерфейсы используют, если тип значения заранее не известен. Например, функция из пакета fmt:

func Println(a ...interface{}) (n int, err error)
fmt.Println() умеет печатать что угодно, поэтому принимает значения типа interface{}.

Начиная с Go 1.18 для interface{} ввели псевдоним any. Разницы между ними нет (псевдоним — это буквально тот же самый тип), но многие теперь предпочитают any за его краткость и выразительность.

func repr(val any) string {
    return fmt.Sprintf("%#v", val)
}

func main() {
    var num int = 42
    fmt.Println(repr(num))
    // 42

    var thing interface{} = "shy string"
    fmt.Println(repr(thing))
    // "shy string"
}
-----------------------------------------------
	
Приведение типа
Приведение типа (type assertion) извлекает конкретное значение из переменной интерфейсного типа:

var value any = "hello"
str := value.(string)
fmt.Println(str)
// hello
Если тип конкретного значения отличается от указанного, произойдет ошибка:

flo := value.(float64)
// ошибка
Чтобы проверить тип конкретного значения, используют опциональный флаг, который сигнализирует — правильный тип или нет:

str, ok := value.(string)
fmt.Println(str, ok)
// hello true

flo, ok := value.(float64)
fmt.Println(flo, ok)
// 0 false
Переключатель типа
Приведение типа можно использовать вместе со switch. Такая конструкция называется переключателем типа (type switch):

var value any = "hello"

switch v := value.(type) {
case string:
    fmt.Printf("%#v is a string\n", v)
case float64:
    fmt.Printf("%#v is a float\n", v)
default:
    fmt.Printf("%#v is a mystery\n", v)
}
// "hello" is a string
v внутри сработавшей ветки переключателя имеет конкретный тип вместо any (в примере — string).

	
-----------------------------------------------
	задачка
	
	Определите интерфейс универсального итератора (iterator), который можно использовать в функции выбора максимального элемента (max). Реализуйте интерфейс для итератора по срезу целых чисел.

Подробности — по коду задания. (на маке у меня не работает)

Sample Input:

1 4 5 2 3
Sample Output:

5
	
		
	package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

)

// element represents an element in the sequence.
type element interface{}

// iterator represents an iterator over a sequence.
type iterator interface {
	next() bool
	val() element
}

// intIterator is an iterator over a sequence of integers.
type intIterator struct {
	nums []int
	pos  int
}

func newIntIterator(nums []int) *intIterator {
	return &intIterator{nums: nums}
}

func (it *intIterator) next() bool {
	if it.pos < len(it.nums) {
		it.pos++
		return true
	}
	return false
}

func (it *intIterator) val() element {
	return it.nums[it.pos-1]
}

// weightFunc is a function that returns the weight of an element.
type weightFunc func(el element) int

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
func main() {
	nums := readInput()
	it := newIntIterator(nums)

	weight := func(el element) int {
		return el.(int)
	}
	m := max(it, weight)
	fmt.Println(m)
}

// max возвращает максимальный элемент в последовательности.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func max(it iterator, weight weightFunc) element {
	var maxEl element
	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}

	

-------------------Ошибки-----------------------------------------------
	
	
	Ошибки
В Go нет исключений и блока try-catch, как в питоне или js. Вместо этого функции явно возвращают ошибку отдельным значением. Благодаря этому ошибки невозможно проигнорировать, а разработчики продумывают поведение программы в случае проблем.

Ошибки принято возвращать последним значением с интерфейсным типом error:

func sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, errors.New("expect x >= 0")
    }
    // `nil` в качестве ошибки указывает, что ошибок не было.
    return math.Sqrt(x), nil
}
Проверим работу sqrt() на положительном и отрицательном значениях. Обратите внимание, как мы получаем результат и проверяем ошибку внутри условия if — это стандартная практика в Go.

for _, x := range []float64{49, -49} {
    if res, err := sqrt(x); err != nil {
        fmt.Printf("sqrt(%v) failed: %v\n", x, err)
    } else {
        fmt.Printf("sqrt(%v) = %v\n", x, res)
    }
}
// sqrt(49) = 7
// sqrt(-49) failed: expect x >= 0



--------
	package main

import (
	"fmt"
	"strings"
)

// lookupError представляет ошибку поиска значения в строке
type lookupError struct {
	src    string
	substr string
}

// Error возвращает описание ошибки
func (e lookupError) Error() string {
	return fmt.Sprintf("'%s' not found in '%s'", e.substr, e.src)
}

// indexOf() возвращает индекс вхождения подстроки substr в строку src
func indexOf(src string, substr string) (int, error) {
	idx := strings.Index(src, substr)
	if idx == -1 {
		return -1, lookupError{src, substr}
	}
	return idx, nil
}

func main() {
	src := "go is awesome"
	for _, substr := range []string{"go", "js"} {
		if res, err := indexOf(src, substr); err != nil {
			fmt.Printf("indexOf(%#v, %#v) failed: %v\n", src, substr, err)
		} else {
			fmt.Printf("indexOf(%#v, %#v) = %v\n", src, substr, res)
		}
	}
	// indexOf("go is awesome", "go") = 0
	// indexOf("go is awesome", "js") failed: 'js' not found in 'go is awesome'

	_, err := indexOf(src, "js")
	if err, ok := err.(lookupError); ok {
		fmt.Println("err.src:", err.src)
		fmt.Println("err.substr:", err.substr)
	}
	// err.src: go is awesome
	// err.substr: js
}


---------------задачка
	https://stepik.org/lesson/526881/step/3?auth=login&unit=519600
Напишите логику работы с лицевым счетом в банке. У счета есть баланс (сумма на счете) и овердрафт (на какую сумму можно уйти в минус). К счету последовательно применяются транзакции списания и пополнения. Транзакции изменяют баланс счета, овердрафт не меняется.
Со счета нельзя списать больше, чем (баланс + овердрафт). Зачислить или списать нулевую сумму тоже нельзя.
Например:
начальное состояние счета: баланс = 100, овердрафт = 10
транзакции: 10, -50, 20
результат: баланс = 80, овердрафт = 10
Или:
начальное состояние счета: баланс = 100, овердрафт = 10
транзакции: -100, -10, -10
результат: недостаточно средств на счете
Гарантируется, что сумма на счете, овердрафт и размер транзакции не выйдут за пределы типа int, как по отдельности, так и все вместе. Гарантируется, что овердрафт больше либо равен 0.
	
	package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

)

// errInsufficientFunds сигнализирует,
// что на счете недостаточно денег,
// чтобы выполнить списание
var errInsufficientFunds error = errors.New("insufficient funds")

// errInvalidAmount сигнализирует,
// что указана некорректная сумма транзакции
var errInvalidAmount error = errors.New("invalid transaction amount")

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

// deposit зачисляет деньги на счет.
func (acc *account) deposit(amount int) error {
	acc.balance += amount

	if amount == 0 {
	return errInvalidAmount
	} else {
	return nil
	}
}

// withdraw списывает деньги со счета.
func (acc *account) withdraw(amount int) error {
	acc.balance -= amount
	if acc.balance + acc.overdraft < 0 {
	return errInsufficientFunds
	} else {
	return nil
	}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

type test struct {
	acc   account
	trans []int
}
//{100 10} [10 -50 20]
var tests = map[string]test{
	"{100 10} [10 -50 20]":   {account{100, 10}, []int{10, -50, 20}},
	"{30 0} [-20 -10]":       {account{30, 0}, []int{-20, -10}},
	"{30 0}, [-20 -10 -10]":  {account{30, 0}, []int{-20, -10, -10}},
	"{30 0}, [-100]":         {account{30, 0}, []int{-100}},
	"{0 0}, [10 20 30]":      {account{0, 0}, []int{10, 20, 30}},
	"{0 0}, [10 -10 20 -20]": {account{0, 0}, []int{10, -10, 20, -20}},
	"{20 10}, [-20 -10]":     {account{20, 10}, []int{-20, -10}},
	"{20 10}, [-20 -10 -10]": {account{20, 10}, []int{-20, -10, -10}},
	"{0 100}, [-20 -10]":     {account{0, 100}, []int{-20, -10}},
	"{0 30}, [-20 -10]":      {account{0, 30}, []int{-20, -10}},
	"{0 30}, [-20 -10 -10]":  {account{0, 30}, []int{-20, -10, -10}},
	"{70 30}, [-100 100]":    {account{70, 30}, []int{-100, 100}},
	"{100 10}, [10 0 20]":    {account{100, 10}, []int{10, 0, 20}},
}

func main() {
	var err error
	name, err := readString()
	if err != nil {
		log.Fatal(err)
	}
	testCase, ok := tests[name]
	if !ok {
		log.Fatalf("Test case '%v' not found", name)
	}
	for _, t := range testCase.trans {
		if t >= 0 {
			err = testCase.acc.deposit(t)
		} else {
			err = testCase.acc.withdraw(-t)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	if err == nil {
		fmt.Println(testCase.acc)
	}
}

// readString считывает и возвращает строку из os.Stdin
func readString() (string, error) {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(str), nil
}


----------------defer-----------------

Defer позволяет отложить выполнение кода до момента завершения функции. Обычно его используют, чтобы освободить ресурсы, выделенные внутри функции (открытые файлы, соединения и тому подобное). В питоне в таких случаях применяют контекстные менеджеры, а в js конструкцию try-finally.

Допустим, мы хотим создать файл, записать в него что-то и закрыть. Вот как поможет defer:


func main() {
    f, err := createFile("/tmp/defer.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }

    defer closeFile(f)    // (1)

    if err := writeFile(f); err != nil {
        fmt.Println("Error writing to file:", err)
        return            // (2)
    }

    fmt.Println("Success!")
}
После того как файл открыт, мы с помощью defer указываем, что необходимо вызвать отложенную функцию closeFile() ➊. Она выполнится в самом конце, при завершении функции main(). Причем отложенная функция отработает в любом случае — даже если во время записи в файл произошла ошибка и сработал досрочный return ➋.

Допустим, создание файла пройдет успешно, а при записи случится ошибка:

func createFile(name string) (*os.File, error) {
    fmt.Println("Creating file...")
    // ...
}

func writeFile(f *os.File) error {
    fmt.Println("Writing to file...")
    // эмулируем неминуемую ошибку
    return fmt.Errorf("oh no, it all went wrong!")
}

func closeFile(f *os.File) {
    fmt.Println("Closing file...")
    // ...
}
closeFile() все равно отработает:

Creating file...
Writing to file...
Error writing to file: oh no, it all went wrong!
Closing file...
	
	
-------Panic--------------
	
	
Если во время выполнения программы происходит неисправимая ошибка, срабатывает паника (panic). Это аналог исключения в питоне или js.

Допустим, мы написали функцию, которая возвращает символ строки по индексу, но забыли проверить, что индекс попадает в границы:

func getChar(str string, idx int) byte {
    return str[idx]
}
Если вызвать getChar() с некорректным индексом — сработает паника:

c := getChar("hello", 10)
// panic: runtime error: index out of range [10] with length 5
Панику можно вызвать и вручную с помощью одноименной встроенной функции:

panic("oops")
Так редко делают — в Go принято возвращать ошибку из функции, а не паниковать.	
	
-----------Recover--------------
	
	
Раз есть непредвиденные ошибки (паника), должен быть и способ их поймать. В Go для этого используется встроенная функция recover(). Посмотрим, как она работает.

Мы все так же забыли проверить, что индекс попадает в границы:

func getChar(str string, idx int) byte {
    return str[idx]
}
Но зная свою забывчивость, решили отловить любые непредвиденные ошибки:

func protect(fn func()) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("ERROR:", err)
        } else {
            fmt.Println("Everything went smoothly!")
        }
    }()
    fn()
}
protect() первым делом объявляет анонимную отложенную функцию, которая сработает после того, как будет выполнена fn(). Если срабатывает паника, вызывается отложенная функция. Внутри нее recover() возвращает ошибку, которая вызвала панику. Если паники не было, отложенная функция тоже вызывается, но recover() внутри возвращает nil.

Здесь сработает паника:

protect(func() {
    c := getChar("hello", 10)
    fmt.Println("hello[10] = ", c)
})
// ERROR: runtime error: index out of range [10] with length 5
А здесь функция отработает без ошибок:

protect(func() {
    c := getChar("hello", 4)
    fmt.Println("hello[4] =", c)
})
// hello[4] = 111
// Everything went smoothly!



Возможно, вы заметили, что ручной вызов panic() в сочетании с defer() и recover() можно использовать, чтобы эмулировать конструкцию try-catch. В Go так не принято. Всегда старайтесь явно возвращать ошибки из функции, а на вызывающей стороне проверять их.
	
	
--------------песок--------------
		
package main

import "fmt"

// getChar возвращает символ строки по указанному индексу
func getChar(str string, idx int) byte {
	return str[idx]
}

// protect ловит панику при выполнении fn и печатает ошибку
func protect(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR:", err)
		} else {
			fmt.Println("Everything went smoothly!")
		}
	}()
	fn()
}

func main() {
	// паника
	protect(func() {
		c := getChar("hello", 10)
		fmt.Println("hello[10] = ", c)
	})
	// ERROR: runtime error: index out of range [10] with length 5

	// без паники
	protect(func() {
		c := getChar("hello", 4)
		fmt.Println("hello[4] =", c)
	})
	// hello[4] = 111
	// Everything went smoothly!
}
	
	
--------------------
	
	
Задачка
	
	package main

// не меняйте импорты, они нужны для проверки
import (
	"bufio"
	"errors"
	//"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

)

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

func main() {
	// var acc account
	// var trans []int
	if acc, trans, err := parseInput(); err != nil {
		fmt.Print(err)
	} else {
		
		fmt.Print("-> ")
		fmt.Println(acc, trans)

	}

}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {

	accSrc, transSrc := readInput()
	acc, err := parseAccount(accSrc)

	if err != nil {
		return acc, nil, err
	}
	trans, err2 := parseTransactions(transSrc)
	
	if err2 != nil && err == nil {
		return acc, trans, err2
	}
	return acc, trans, nil

}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc

}

// parseAccount парсит счет из строки))
// в формате balance/overdraft.
func parseAccount(src string) (account, error) {
	parts := strings.Split(src, "/")
	balance, err := strconv.Atoi(parts[0])

	overdraft, err2 := strconv.Atoi(parts[1])

	if err != nil {
		fmt.Printf("-> strconv.Atoi: parsing %#v: invalid syntax", parts[0])
		return account{balance, overdraft}, errors.New("")
	}
	if err2 != nil {
		fmt.Printf("-> strconv.Atoi: parsing %#v: invalid syntax", parts[1])
		return account{balance, overdraft}, errors.New("")
	}

	if overdraft < 0 {
		return account{balance, overdraft}, errors.New("expect overdraft >= 0")
	}
	if balance < -overdraft {
		return account{balance, overdraft}, errors.New("balance cannot exceed overdraft")
	}
	return account{balance, overdraft}, nil
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		t, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("-> strconv.Atoi: parsing %#v: invalid syntax", s)
			return trans, errors.New("")
		}
		trans[idx] = t
	}
	return trans, nil
}

-----------------------------------------
	errors.Is() и errors.As()
	
	
Получив ошибку-матрешку, клиент может проверить, есть ли на каком-то слое интересующая его проблема. Для этого используют функцию errors.Is():

func (l languages) describe(lang string) (string, error) {
    descr, err := getValue(l, lang)
    if err != nil {
        return "", languageErr{lang, err}
    }
    return descr, nil
}

// ...

func main() {
    descr, err := langs.describe("java")
    if errors.Is(err, errNotFound) {
        fmt.Println("this is an errNotFound error")
        // do something about it...
    }
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(descr)
}
this is an errNotFound error
java language error: not found
Неважно, сколько в ошибке слоев. Если на каком-то из них встретилось значение errNotFound — errors.Is() вернет true.

песочница

errors.As()
Раньше мы использовали приведение типа, чтобы получить доступ к ошибке конкретного типа вместо абстрактного error:

descr, err := langs.describe("java")
if langErr, ok := err.(languageErr); ok {
    fmt.Println("Language error:", langErr.lang)
}
Language error: java
Но это работает только для ошибки самого верхнего уровня. До ошибки из середины «матрешки» через приведение типа не добраться. А вот через errors.As() — можно:

descr, err := langs.describe("java")
// обернем еще раз, чтобы languageErr
// оказалась внутрь матрешки
err = fmt.Errorf("wrap once more: %w", err)

var langErr languageErr
if errors.As(err, &langErr) {
    fmt.Println("Language error:", langErr.lang)
}
Language error: java
wrap once more: java language error: not found
errors.As() проверяет каждый слой ошибки, и если видит там искомый тип languageErr — заполняет значение langErr по переданному указателю, и возвращает true. Если искомого типа нет — возвращает false.



Итого по обертыванию:

Простой способ создать новую ошибку на основе существующей — fmt.Errorf() и спецификатор %w
Если нужен собственный тип ошибки, придется добавить в него поле типа error и метод Unwrap()
errors.Is() проверяет конкретную ошибку на каждом слое.
errors.As() заполняет ошибку конкретного типа, если он встречается на одном из слоев.
Закрепим задачкой.
	
	
	песок
		
		
	package main

import (
	"errors"
	"fmt"

)
// Неважно, сколько в ошибке слоев. Если на каком-то из них встретилось значение errNotFound — errors.Is() вернет true.
// errNotFound описывает ошибку поиска ключа в карте
var errNotFound error = errors.New("not found")

// getValue извлекает значение по ключу из карты
func getValue(m map[string]string, key string) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", errNotFound
	}
	return val, nil
}

// languages представляет информацию о языках
type languages map[string]string

// describe возвращает описание языка по названию
func (l languages) describe(lang string) (string, error) {
	descr, err := getValue(l, lang)
	if err != nil {
		return "", fmt.Errorf("error describing %s: %w", lang, err)
	}
	return descr, nil
}

var langs languages = languages{
	"go":     "is awesome",
	"python": "is everywhere",
	"php":    "just is",
}

func main() {
	descr, err := langs.describe("java")
	
	if errors.Is(err, errNotFound) {
		fmt.Println("this is an errNotFound error")
		// do something about it...
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(descr)
}

и 
	
	
package main

import (
	"errors"
	"fmt"
)

// errNotFound описывает ошибку поиска ключа в карте
var errNotFound error = errors.New("not found")

// getValue извлекает значение по ключу из карты
func getValue(m map[string]string, key string) (string, error) {
	val, ok := m[key]
	if !ok {
		return "", errNotFound
	}
	return val, nil
}

// languageErr описывает ошибку получения языка
type languageErr struct {
	lang string
	err  error
}

// Error возвращает строковое представление ошибки
func (le languageErr) Error() string {
	return fmt.Sprintf("%s language error: %v", le.lang, le.err)
}

// Unwrap возвращает внутреннюю ошибку
func (le languageErr) Unwrap() error {
	return le.err
}

// languages представляет информацию о языках
type languages map[string]string

// describe возвращает описание языка по названию
func (l languages) describe(lang string) (string, error) {
	descr, err := getValue(l, lang)
	if err != nil {
		return "", languageErr{lang, err}
	}
	return descr, nil
}

var langs languages = languages{
	"go":     "is awesome",
	"python": "is everywhere",
	"php":    "just is",
}

func main() {
	descr, err := langs.describe("java")
	err = fmt.Errorf("wrap once more: %w", err)
	var langErr languageErr
	if errors.As(err, &langErr) {
		fmt.Println("Language error:", langErr.lang)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(descr)
}

--------------------------------------------
	
	
	package main

import (
	"bufio"
	"fmt"
	//"strconv"
	"os"
	
)

func main() {

	var text string
	text, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	var count int

	

	for _ ,i := range text {
		if i == ' '{
		count ++
		
		}
	}
	if count == 0 {
		fmt.Println(count)
	}else {
	fmt.Println(count+1)
	}
	
	
}

--------------------------------------
	
	
	
	
	
	
	
	