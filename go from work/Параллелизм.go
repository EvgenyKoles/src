Параллелизм, многопоточность

channel <- num    // отправляет в канал channel значение num
num = <- channel // получение из канала channel в переменную num
----------------------------
Здесь есть две концепции и обе они совершенно разные, первая синхронное и асинхронное программирование и вторая – однопоточные и многопоточные приложения.

1- Синхронная программная модель – это программная модель, когда потоку назначается одна задача и начинается выполнение. Когда завершено выполнение задачи тогда появляется возможность заняться другой задачей. В этой модели невозможно останавливать выполнение задачи чтобы в промежутке выполнить другую задачу. 

Однопоточность – если мы имеем несколько задач, которые надлежит выполнить, и текущая система предоставляет один поток, который может работать со всеми задачами, то он берет поочередно одну за другой и процесс выглядит так:
thread1 - task1->task2->tsak3->taks4 --->

Многопоточность – в этом сценарии, мы использовали много потоков, которые могут брать задачи и приступать к работе с ними. У нас есть пулы потоков (новые потоки также создаются, основываясь на потребности и доступности ресурсов) и множество задач. Итак, поток может работать вот так:
thread1 - task1---->
thread2 - task2---->
thread3 - task3---->
thread4 - task4---->
Это идеальный сценарий, но в обычных условиях мы используем большее количество задач чем количество доступных потоков, таким образом освободившийся поток получает другое задание

2 - Асинхронная модель программирования – в отличии от синхронной программной модели, здесь поток однажды начав выполнение задачи может приостановить выполнение сохранив текущее состояние и между тем начать выполнение другой задачи.

однопоточный
thread1 -> T1 - T2- T1-T3-T4-T1-T3 --->
Здесь мы можем видеть, что один поток отвечает за выполнение всех задач и задачи чередуются друг за другом.

многопоточный
Если наша система способно иметь много потоков тогда все потоки могут работать в асинхронной модели как показано ниже:

thread1 --> T1-T2-T1-T3-T4-T1-T3 --->
thread2 --> T4-T5-T6-T7-T48 --->
thread3 --> T5-T8-T6-T3-T47-T3-T8 --->
thread4 --> T8-T7-T8-T7-T6-T8-T3 --->

Одна задача может начать обрабатываться на одном потоке, и заканчивает на другом

------Конкурентность (Concurrency) и параллелизм (Parallelism)------
	
Конкурентность предполагает работу приложения с двумя и более задачами одновременно, когда происходит создание нескольких процессов, выполняющихся независимо друг от друга. Когда мы начинаем говорить о многопоточной разработке, нужно ввести такие понятия, как Concurrency и Parallelism. В мире Go есть выражение «Concurrency is not Parallelism». Суть в том, что Concurrency — это о дизайне, то есть о том, как мы проектируем нашу программу. Parallelism — это просто способ выполнения нашего кода. 

https://ucarecdn.com/6dd28b1b-bc63-4bbb-973f-845e74cee366/
https://ucarecdn.com/81faff0e-13f8-4f2f-95e3-0d03a60fddaf/


---Горутины---
	
	
Горутина не имеет полного соответствия с потоками. Но похожи
	
Горутины (goroutines) представляют параллельные операции
 
Преимущества горутин:
Они легковесны.
Легко и без проблем масштабируют.
Они — практически потоки.
Требуют меньше памяти (2KB).
Предоставляют дополнительную память горутинам во время выполнения.

	Запуск горутины очень прост, достаточно прописать перед функцией ключевое слово "go", например:

package main

import "fmt"

func main() {
	go myFunc()
}

func myFunc() {
	fmt.Println("hello")
}
//ничего не выведет
		
функция main продолжает выполняться и завершается, не дожидавшись завершения выполнения всех прочих горутин, соответственно myFunc просто не успевает завершить выполнение. Go предусматривает несколько способов синхронизации выполнения горутин и мы рассмотри их в этом уроке.
	
если внесем паузу
	
	
package main

import (
	"fmt"
	"time"
)

func main() {
	go myFunc()
	time.Sleep(1 * time.Second) // Пауза в 1 секунду
}

func myFunc() {
	fmt.Println("hello")
}

// Вывод: hello


---анонимные горутины---
	
	//Самовызывающаяся анонимная горутина
go func() {
 fmt.Println("Привет, я анонимная горутина")
}()
Это позволяет писать код, который используется в одном месте и сразу вызвать, не заботясь об объявлении функции


----
Функция main()
Удивительно, но функция main() вызывает свою собственную горутину.

Замечание

По умолчанию используется кол-во ядер вашего процессора. Но мы можем изменить количество используемых ядер простой строчкой кода. Приложению будет дана команда перейти на 4 ядра:
runtime.GOMAXPROCS(4)



-----каналы------
	
	

Канал — это объект связи, с помощью которого горутины обмениваются данными. 
Технически это конвейер (или труба), откуда можно считывать или помещать данные. 
То есть одна горутина может отправить данные в канал, а другая 
— считать помещенные в этот канал данные.
	
создали канал

package main
import "fmt"

func main() {
    var c chan int
    fmt.Println(c)
} 

инициализировать канал
можете передать или получить данные из канала

package main
import "fmt"

func main() {
    c := make(chan int)

    fmt.Printf("type of `c` is %T\n", c)
    fmt.Printf("value of `c` is %v\n", c)
}


вывод - 
type of `c` is chan int
value of `c` is 0xc0420160c0


c <- data 	- передаем данные в канал
data := <- c   	двнные из канала c, который имеет тип int, 
могут быть записаны в переменную data
Go определит тип данных, передаваемый каналу c, 
и предоставит data корректный тип данных.
	
	
-------канал на практике---------

package main

import "fmt"

func greet(c chan string) { // функция greet,  принимает канал С как аргумент. 
    //В этой функции мы считываем данные из канала c и выводим в консоль.
    fmt.Println("Hello " + <-c + "!") 
}

func main() {
    fmt.Println("main() started") //
    c := make(chan string) // создаем канал c с типом даных string

    go greet(c) //Помещаем канал с в функцию greet и запускаем функцию 
    //как горутину, используя ключевое слово go.
//Теперь у нас имеется две горутины main и greet,
// main по-прежнему остается активной.

    c <- "John" // Помещаем данные в канал с. 
    //В этот момент main блокируется до тех пор, пока другая горутина (greet) не считает данные из канала c.
    // Планировщик Go выполняет функцию greet 
    fmt.Println("main() stopped") // После чего main снова становится активной и выводит в консоль "main() stopped"
}

вывод 
main() started
Hello John!
main() stopped
	
	
	//добавил свое горутину. не понял почему первая выполняется buy buy		
package main

import "fmt"
func greet(c chan string) { 
    fmt.Println("Hello " + <-c + "!") 
}
func goodbuy(a chan string){
	fmt.Println("Buy,Buy " + <-a )
}
func main() {
    fmt.Println("main() started")
    c := make(chan string) 
	abc := make(chan string)
    go greet(c) 
	go goodbuy(abc)
	abc <- "Evgeny"
    c <- "John"  
	fmt.Println("main() stopped") 
}


main() started
Buy,Buy Evgeny
Hello John!
main() stopped
//---------------------------------------------
	
	
//хороший пример демонитрации как работает блокировка
//вывод 
a08406776@iMac-4 new_folder % go run hello.go
main() started
0 true
1 true
4 true
9 true
16 true
25 true
36 true
49 true
64 true
81 true
0 false <-- loop broke!
main() stopped


----
package main

import "fmt"

func squares(c chan int) {
    for i := 0; i <= 9; i++ {
        c <- i * i
    }

    close(c) // close channel
}

func main() {
    fmt.Println("main() started")
    c := make(chan int)

    go squares(c) // start goroutine

    // periodic block/unblock of main goroutine until chanel closes
    for {
        val, ok := <-c
        if ok == false {
            fmt.Println(val, ok, "<-- loop broke!")
            break // exit break loop
        } else {
            fmt.Println(val, ok)
        }
    }

    fmt.Println("main() stopped")
}

с рэндж

package main

import "fmt"

func squares(c chan int) {
    for i := 0; i <= 9; i++ {
        c <- i * i
    }

    close(c) // close channel
}

func main() {
    fmt.Println("main() started")
    c := make(chan int)

    go squares(c) // start goroutine

    // periodic block/unblock of main goroutine until chanel closes
    for val := range c {
        fmt.Println(val)
    }

    fmt.Println("main() stopped")
}
//Если вы не закроете канал для цикла for с использованием range, 
//то программа будет завершена аварийно из-за dealock во время выполнения.
//------------------------------------

буферезированный канал

package main

import "fmt"

func squares(c chan int) {
    for i := 0; i <= 3; i++ {
        num := <-c
        fmt.Println(num * num)
    }
}

func main() {
    fmt.Println("main() started")
    c := make(chan int, 3)

    go squares(c)

    c <- 1
    c <- 2
    c <- 3
    c <- 4 // blocks here

    fmt.Println("main() stopped")
}


//--------------------------------
Длина и емкость канала

Подобно срезам, буферизированный канал имеет длину и емкость. Длина канала — это количество значений в очереди (не считанных) в буфере канала, емкость — это размер самого буфера канала. Для того, чтобы вычислить длину, мы используем функцию len, а, используя функцию cap, получаем размер буфера.

package main

import "fmt"

func main() {
    c := make(chan int, 3)
    c <- 1
    c <- 2

    fmt.Printf("Length of channel c is %v and capacity of channel c is %v", len(c), cap(c))
    fmt.Println()
}


Вывод программы:
Length of channel c is 2 and capacity of channel c is 3

еще пример с буферизированным каналом:

package main

import (
    "fmt"
    "runtime"
)

func squares(c chan int) {
    for i := 0; i < 4; i++ {
        num := <-c
        fmt.Println(num * num)
    }
}

func main() {
    fmt.Println("main() started")
    c := make(chan int, 3)
    go squares(c)

    fmt.Println("active goroutines", runtime.NumGoroutine())
    c <- 1
    c <- 2
    c <- 3
    c <- 4 // blocks here

    fmt.Println("active goroutines", runtime.NumGoroutine())

    go squares(c)

    fmt.Println("active goroutines", runtime.NumGoroutine())

    c <- 5
    c <- 6
    c <- 7
    c <- 8 // blocks here

    fmt.Println("active goroutines", runtime.NumGoroutine())
    fmt.Println("main() stopped")
}



//-----
Используя буферизованный канал и цикл for range, мы можем читать с закрытых каналов. Поскольку у закрытых каналов данные все еще живут в буфере, их можно считать:

package main

import "fmt"

func main() {
    c := make(chan int, 3)
    c <- 1
    c <- 2
    c <- 3
    close(c)

    // iteration terminates after receiving 3 values
    for elem := range c {
        fmt.Println(elem)
    }
}


//--------------------------------

две горутины с выводом

package main

import "fmt"

func square(c chan int) {
    fmt.Println("[square] reading")
    num := <-c //пишем данные в переменую
    c <- num * num
}

func cube(c chan int) {
    fmt.Println("[cube] reading")
    num := <-c  //пишем данные в переменую
    c <- num * num * num
}

func main() {
    fmt.Println("[main] main() started")

    squareChan := make(chan int) //создали 2 канала с инт
    cubeChan := make(chan int)

    go square(squareChan) //запускаем square и cube горутины
    go cube(cubeChan)

    testNum := 3 // Так как контроль по-прежнему внутри main testNum получает значение 3.
    fmt.Println("[main] sent testNum to squareChan")
//Затем мы отправляем данные в канал squareChan и cubeChan. 
//Горутина main будет заблокирована, пока данные из каналов не будут считаны. 
//Как только значение будет считано, горутина снова станет активной.
    squareChan <- testNum

    fmt.Println("[main] resuming")
    fmt.Println("[main] sent testNum to cubeChan")

    cubeChan <- testNum

    fmt.Println("[main] resuming")
    fmt.Println("[main] reading from channels")

    squareVal, cubeVal := <-squareChan, <-cubeChan
    sum := squareVal + cubeVal
    //Когда операция записи канала завершена, начинает выполняться main,
    //после чего мы рассчитываем сумму и выводим ее.
    fmt.Println("[main] sum of square and cube of", testNum, "is", sum)
    fmt.Println("[main] main() stopped")
}

вывод 

a08406776@iMac-4 new_folder % go run hello.go
[main] main() started
[main] sent testNum to squareChan
[cube] reading
[square] reading
[main] resuming
[main] sent testNum to cubeChan
[main] resuming
[main] reading from channels
[main] sum of square and cube of 3  is 36
[main] main() stopped
a08406776@iMac-4 new_folder % 


//-------------------------------------------
Однонаправленные каналы

До сих пор мы видели каналы, которые могут передавать и принимать данные. 
Но мы также можем создать канал, который будет однонаправленным. 
Например, канал, который сможет только считывать данные, 
и канал который сможет только записывать их.

Однонаправленный канал также создается с использованием make, 
но с дополнительным стрелочным синтаксисом.

roc := make(<-chan int)
soc := make(chan<- int)

Где roc канал для чтения, а soc канал для записи.
 Следует заметить, что каналы также имеют разный тип.

package main

import "fmt"

func main() {
    roc := make(<-chan int)
    soc := make(chan<- int)

    fmt.Printf("Data type of roc is `%T`\n", roc)
    fmt.Printf("Data type of soc is `%T\n", soc)
}

-----
Go предоставляет простой синтаксис для преобразования двунаправленного канала 
в однонаправленный канал.

import "fmt"

func greet(roc <-chan string) {
    fmt.Println("Hello " + <-roc + "!")
}

func main() {
    fmt.Println("main() started")
    c := make(chan string)

    go greet(c)

    c <- "John"
    fmt.Println("main() stopped")
}



//----------------------------------------------------

задачи

Напишите функцию которая принимает канал и число N типа int. 
Необходимо вернуть значение N+1 в канал. 
Функция должна называться task().

Внимание! Пакет и функция main уже объявлены, выводить и считывать ничего не нужно!

package main

import (
	"fmt"
	//"time"
)

func main() {
	m := make (chan int, 3)
	myfunc(m,8)
}

func myfunc (c chan int, number int) {
	c <- number+1
	fmt.Print(<-c)	
}

