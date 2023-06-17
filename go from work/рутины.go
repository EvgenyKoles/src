--------------------------------------
	статья
	Go Channels Internals https://habr.com/ru/companies/oleg-bunin/articles/522742/
	
	
	
	type hchan struct // структура, которая содержит основные данные;
	type sudog struct // это супер-сет of g со специальными полями, которые нужны для работы каналов, select и прочего.
	type waitq struct // структура, которая представляет собой linkedlist (связанный список) и хранит в себе два pointers — на голову и хвост
	
	Что умеют каналы?

-У каналов есть размер. Это очевидно и понятно. Вы можете создать канал размером N, 10, 5, 1 и т.д. То есть даже когда вы создаете небуферизированный канал (как в предыдущем примере make(chan string)), вы на самом деле создаете буферизированный, у которого размер 0.
-Передавать данные между горутинами. Это цель существования каналов.
Каналы горутино-безопасны. Мы не паримся локами, синхронизацией, ужасными CAS a loop, не думаем о data races и т.д. Мы просто берем канал, пишем с одной стороны, читаем с другой, и всё замечательно. Даже если мы вызовем close, мы не паримся — зачем нам локи, если они потоко-безопасны?
-FIFO. Канал работает как очередь, то есть First In — First Out. Это означает, что если 5 горутин пытаются записать в канал, они всегда сделают это в определенном порядке.
-Горутины ждут каналы. То есть канал может повесить горутину, которая может подождать результата из канала или подождать перед записью в канал и т.д.
-Каналы могут блокировать/разблокировать горутины. Вытекающий из предыдущего пункт.

Но это всё лирика. Если мы рассмотрим публичные API каналы, то обнаружим, что есть всего 4 публичных действия:
newChan := make(chan int) — создать канал;
newChan< — 1 — записать в канал;
<-newChan — прочитать из канала;
close(newChan) — закрыть канал.
	
----------------------
	
Рассмотрим простой пример:
plate := make(chan string, 3) 
	

Что происходит, когда вызывается этот код? Создается структура, которая содержит в себе около 10 полей, из них — 4 важных. 
Первое — Ring buffer (circle queue, или кольцевая очередь), которое хранит наши данные.

Обычно имплементация кольцевой очереди делается на массиве и на двух указателях (индексах) в этом массиве: один индекс мы постоянно пишем и инкрементируем, 
а другой индекс читаем и инкрементируем. 
И когда мы достигаем конца, то просто перемещаемся в начало. 
Здесь же у нас есть массив и два индекса, по которым пишем и читаем. И есть еще lock. При этом у нас пустая структура — и буфер, и индексы пустые:

https://habrastorage.org/r/w780/webt/9c/qn/r-/9cqnr-l5i8mi25kde_o23jmqvm4.jpeg

	
Про nil буферизированный канал
	Когда мы начнем читать из этого канала, мы начнем выстраиваться в очередь на нём, так как у нас нет ни пишущих горутин, ни буфера. У нас есть только две горутины, которые пытаются прочитать из канала. И они одна за одной выстроятся в очередь: «Ага, я подожду, пока кто-нибудь придет». Если приходит пишущая горутина, она подумает: «О, в списке получения есть что-то — запишу-ка я туда». Иначе — она точно так же встанет в запись и будет ждать. А waitq — указатель на голову и хвост linked листа горутин:
--------
Когда я начал программировать на Go, меня всегда интересовало: почему в примерах, когда нам просто нужен канал, или, допустим, хэш-сет, мы просто пишем пустую структуру?

func main(){
  c := make(chan struct{})
  go func() {c <- struct{}{}}()
  <- c
}


Ответ: потому что все сложные примитивы примерно одинаково в Go работают. 
Когда у нас пустой struct, это по сути special case, элемент size — ноль, тип максимально упрощен. 
При использовании special case не грузится дополнительная информация, мы не таскаем с собой дополнительные размеры, экономим место, и поэтому счастливо живем:
makechan: chan=0xc0000700c0; elemsize=0; elemalg=0x10bef90; dataqsiz=0
	

--------------------------------------
package main

import (
	//"container/list"
	"fmt"
)

func removeDuplicates(t, digit []string) {
	for i := 1; i < len(t); i++ {
		if t[i] != t[i-1] {
			digit = append(digit, t[i-1])
			}
		if i == len(t)-1 && t[len(t)-1] != t[len(t)-2]{
			digit = append(digit, t[len(t)-1])
		}
	}
	fmt.Println(digit)
	
}

func main() {

	t := []string{"1", "2", "3", "3", "4", "5"}
	digit := make([]string, 0, 3)

	removeDuplicates(t, digit)

}

---------------------------
	закрыт ли канал 
	
	val, ok := <- inputStream
        if !ok {
           //......
        }
        
---------------------------
	
Задачка: 
	Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.
Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки, во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал ;)

Функция должна называться removeDuplicates()

Выводить или вводить ничего не нужно!
	
	
	package main

import (
	"fmt"

)

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	go func() {
		defer close(inputStream)
		for _, r := range "q112334456" {
			inputStream <- string(r)
		}
	}()
	for x := range outputStream {
		fmt.Print(x)
	}
}


func removeDuplicates(in chan string, out chan string) {
    var m string //пустая строка
   
    for i := range in { //проходим по q112334456
        if m != i { // если строка не ровна строке(в канале именно строки)
            out <- i // пишем эту строку в канал вывода
        }
        m = i // или присваемаем ее нашей сроке м
    }
    close(out)
}
	
---------------------------------
		
package main
import (
	"fmt"
)

func main() {

	<-myFunc() // игнорируем полученные данные
	// строчка выше аналогична записи
    // c := myFunc() 
    // <-c

}

func myFunc() <-chan struct{} {

	//Функция myFunc создает канал, запускает горутину и сразу же возвращает созданный канал, завершая свое выполнение.

	done := make(chan struct{})
	go func() {
		fmt.Println("hello")
		close(done)
	}()

	return done
}


---------------------------------
	
	
//Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения. 

//Функция work() ничего не принимает и не возвращает.
	
	
	done := make(chan struct{})

go func(d chan struct{}) {
    work()
    time.Sleep(1 * time.Second)
    close(d)
}(done)
<-done
	
	
	
---------------------------------------------
	Очень ВАЖНОЕ замечание: wg.Add() должно вызываться ДО запуска горутины, а wg.Done() внутри горутины. Новички, часто себе стреляют в ногу таким образом и получают несоответствие количества вызов Add() количеству вызовов Done(), что приводит к runtime panic
	
	package main

import (
	"fmt"
	"sync"
	//"time"
)

func main() {
	wg := new(sync.WaitGroup)// определяем группу

	for i := 0; i < 5; i++ {
		wg.Add(1)      // Увеличиваем счетчик горутин в группе
		go work(i, wg) // Вызываем функцию work в отдельной горутине
	}

	wg.Wait() // ожидаем завершения всех горутин в группе, счетчик горутин станет нулевым. wg.Add(1)++, wg.Done()--. !!!!!!!!!!!!!!!!!!!
	fmt.Println("Горутины завершили выполнение")
}

func work(id int, wg *sync.WaitGroup) {

	defer wg.Done()//сигнализировать, что элемент группы завершил свое выполнение
	fmt.Printf("Горутина %d начала выполнение \n", id)
	//time.Sleep(2 * time.Second)
	fmt.Printf("Горутина %d завершила выполнение \n", id)
}

	
	
	
-----------
	wg := new(sync.WaitGroup)
for i := 0; i < 5; i++ {
    // у wg теперь есть свой счетчик, мы его должны вручную увеличивать при создании каждой задачи
    wg.Add(1) // Увеличиваем счетчик горутин в группе
    // внутри нашей задачи нам нужно будет по выполнению задачи счетчик уменьшить
    // методом wg.Done(), который под капотом просто уменьшает счетчик на 1
    // поэтому передаем wg в ту функцию-горутину, которую нужно будет подождать
    // (work в нашем примере)
    go work(i, wg) // Вызываем функцию work в отдельной горутине
}
// у wg есть метод Wait(), который тормознет текущий процесс (main в данном примере)
// до тех пор, пока его, wg, счетчик не станет нулевым - но мы то знаем, что это 
// произойдет как раз когда все задачи будут выполнены ;)
wg.Wait() // ожидаем завершения всех горутин в группе fmt.Println("Горутины завершили выполнение") }
	
	
----------------------------------------------------------------
		
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельных горутинах вызвать функцию work() 10 раз и дождаться результатов выполнения вызванных функций.
 

Функция work() ничего не принимает и не возвращает. Пакет "sync" уже импортирован.
	
	package main

import (
	"fmt"
	"sync"
	//"time"

)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1) 
		
		go func() {
			work()
			//time.Sleep(1 * time.Second)
			wg.Done()
		}()
     }

	wg.Wait() 
	fmt.Println("Горутины завершили выполнение")
}

func work() {
fmt.Print("+")
}

	
	
-------------мьютекс------------
	
package main

import (
	"fmt"
	"sync"
)

func main() {
	var x int
	wg := new(sync.WaitGroup)

	for i := 0; i < 1000; i++ {
		// Запускаем 1000 экземпляров горутины, увеличивающей счетчик на 1
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			x++
		}(wg)
	}

	wg.Wait()

	// По идее значение счетчика должно быть 1000, но крайне вероятно, что этого не произойдет
	fmt.Println(x)
}


--------- вывод не будет 1000, т.к. Представим себе что первая горутина получает значение переменной x, а вторая горутина одновременно с этим выполняют такую же операцию. Тогда обе горутины считают, что x = 0, затем производятся расчеты и обе горутины присваивают x значение 1 (0 + 1), в результате работа одной из горутин напрасна.
	

Чтобы не потерять результаты вычислений мы можем использовать тип sync.Mutex:

package main

import (
	"fmt"
	"sync"
)

func main() {
	var x int
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for i := 0; i < 1000; i++ {
		// Запускаем 1000 экземпляров горутины, увеличивающей счетчик на 1
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			x++
			mu.Unlock()
		}(wg, mu)
	}

	wg.Wait()
	fmt.Println(x)
}

если одна горутина взяла Lock, то другие горутины этой операции выполнить уже не могут и вынуждены ожидать, пока взявшая Lock горутина не выполнит Unlock.
	
	

	
	
	
----------------Пакет time и параллелизм---------------

Мы уже прошли пакет time, но есть функции которые связаны с параллелизмом.
Перед тем как перейти к рассмотрению типов Timer и Ticker рассмотрим упрощенные их аналоги:

// func Sleep(d Duration)
// программа засыпает на заданное время
time.Sleep(time.Second * 2) // спим ровно 2 секунды

// func After(d Duration) <-chan Time
// создает канал, который через заданное время вернет значение
timer := time.After(time.Second)
<-timer // значение будет получено из канала ровно через 1 секунду

// func Tick(d Duration) <-chan Time
// создает канал, который будет посылать сигналы постоянно через заданный промежуток времени
ticker := time.Tick(time.Second)
count := 0

for {
	<-ticker
	fmt.Println("очередной тик")
	count++
	if count == 3 {
		break
	}
}

// очередной тик
// очередной тик
// очередной тик

--------------------------------
	Таймеры и тикеры
	
type Timer
Timer по своей сути очень похож на результат работы After, но позволяет остановить таймер или изменить время его выполнения:

t := time.NewTimer(time.Second) // создаем новый таймер, который сработает через 1 секунду
go func() {
	<-t.C // C - канал, который должен вернуть значение через заданное время
}()
t.Stop() // но мы можем остановить таймер и раньше установленного времени

t.Reset(time.Second * 2) // пока таймер не сработал, мы можем сбросить его, установив новый срок выполнения
<-t.C
type Ticker
Ticker же работает как функция Tick, но может быть остановлен:

func NewTicker(d Duration) *Ticker // создаем новый Ticker
func (t *Ticker) Stop() // останавливаем Ticker


------------
	
	//В примере таймер использован для передачи сигнала о необходимости завершить работу, а тикер выдавал «билет» на выполнение работы. Такой подход обычно используется для балансировки нагрузки.
		
	package main

import (
	"fmt"
	"time"

)

func main() {
	<-work()
	/*
	 * тик-так
	 * тик-так
	 * тик-так
	 * тик-так
	 */
}

func work() <-chan struct{} {
	done := make(chan struct{}) // канал для синхронизации горутин

	go func() {
		defer close(done) // синхронизирующий канал будет закрыт, когда функция завершит свою работу

		stop := time.NewTimer(time.Second)

		tick := time.NewTicker(time.Millisecond * 200)
		defer tick.Stop() // освободим ресурсы, при завершении работы функции

		for {
			select {
			case <-stop.C:
				// stop - Timer, который через 1 секунду даст сигнал завершить работу
				return
			case <-tick.C:
				// tick - Ticker, посылающий сигнал выполнить работу каждый 200 миллисекунд
				fmt.Println("тик-так")
			}
		}
	}()
	return done
}

какой канала к моменту case готов выполнять тот и стартует
	
---------------select в канале горутин----------------------	
	
	//Оператор select выглядит как оператор switch. Каждый case внутри select содержит канал получения или отправки. select ждет завершения одного case, а затем запускает его и связанный с ним оператор case. Как будто select смотрит на оба канала сразу и действует, когда видит, что что-то случается с любым из них.
	тут все спят
	
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
  	select { // Оператор select
    case gopherID := <-c: // Ждет, когда проснется гофер
        fmt.Println("gopher ", gopherID, " has finished sleeping")
    case <-timeout: // Ждет окончания времени
        fmt.Println("my patience ran out")
        return // Сдается и возвращается
    }

----
package main

import (
	"fmt"
	"math/rand"
	"time"

)

func init() {
	rand.Seed(time.Now().Unix()) // Настраиваем рандом так, чтобы он был разным для каждого вызова
}

func sleepyGopher(id int, c chan int) {
	duration := time.Duration(rand.Intn(400)) * time.Millisecond
	fmt.Printf("gopher %d sleep for %v\n", id, duration)
	time.Sleep(duration)// что бы гоферы не просыпались сразу а подождали всех
	c <- id
}

func main() {
	timeout := time.After(2 * time.Second)

	c := make(chan int, 5)

	/**
	Горутины для гоферов нужно создать заранее. Если делать это в for вместе с select, то select будет блокировать дальнейшее исполнение и создание cледующей горутины
	**/

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		select { // Оператор select
		case gopherID := <-c: // Ждет, когда проснется гофер
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout: // Ждет окончания времени
			fmt.Println("my patience ran out")

			return // Сдается и возвращается
		}
	}
}


------------------------
	
package main

import (
	"fmt"
	"time"

)

Функция main начинается с пятью горутинами sleepyGopher. Все они спят по три секунды, а затем выводят на экран одно и то же.

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher()
	}
	time.Sleep(4 * time.Second)
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}

-------
	
package main

import (
	"fmt"
	"time"

)

// Мы можем выяснить кто выполнится первым, передав аргумент каждой горутине. Передача аргумента горутине похожа на передачу аргумента любой другой функции: значение копируется и передается как параметр.
//При запуске следующего листинга вы увидите, что хотя мы запустили все горутины по очереди от нуля до девяти, они все завершились в разное время

func main() {
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(4 * time.Second)
}

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...", id)
}
	
	
-------еще пример------
	
https://golangify.com/wp-content/uploads/2020/03/gorutiny-go.png
	
func main() {
    c := make(chan int) // Делает канал для связи
    
    for i := 0; i < 5; i++ {
        go sleepyGopher(i, c)
    }
    for i := 0; i < 5; i++ {
        gopherID := <-c // Получает значение от канала
        fmt.Println("gopher ", gopherID, " has finished sleeping")
    }
}
 
func sleepyGopher(id int, c chan int) { // Объявляет канал как аргумент
    time.Sleep(3 * time.Second)
    fmt.Println("... ", id, " snore ...")
    c <- id // Отправляет значение обратно к main
}





--------------------------------------------
	

//В этом примере i была общей для всех создаваемых горутин, и пока внутри горутины мы дойдем до создания элемента map i уже станет равной 10, и получается что 10 горутин создают один элемент map,

package main

import (
    "fmt"
    "sync"
    "time"
)

const N = 10

func main() {
    m := make(map[int]int)

    wg := &sync.WaitGroup{}
    mu := &sync.Mutex{}
    wg.Add(N)
    for i := 0; i < N; i++ {
    fmt.Println("до анонимной",i)
    time.Sleep(1 * time.Second)
        go func() {
            defer wg.Done()
            mu.Lock()
            m[i] = i
    
            fmt.Println("В анонимной",i)
            fmt.Println(m)
            mu.Unlock()
        }()
    }
    wg.Wait()
    fmt.Println(len(m))
    }

//В текущем примере i уже индивидуальна для каждой горутины, и поэтому мы создаем уникальный элемент map. Для понимания, что и в этом примере сначала отработает главная горутина с циклом for, не прерывая его, и только потом запустятся созданные в цикле горутины, работающие с мапой
	
	
package main

import (
	"fmt"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(m))
}



------------------------------------------

из ютуба
	
	package main

import (
	"fmt"
	"time"
	// "math/rand"
	// "time"

)


func main() {

	ch := make(chan string)

	go say("Hello",ch)	

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println(<-ch)//чтение из канала блочит поток. Мэин ждет пока тут данные появятся
}

func say (world string, ch chan string){
	time.Sleep(10* time.Second) //вот тут ждет
	//fmt.Println(world)
	ch <-world // передали данные в канал в общую рутину

}
	
	
-------
	
package main

import (
	"fmt"
	"time"
	//"time"
	// "math/rand"

)

func main() {

	ch := make(chan int)
	go sayHello(ch)
	

	for i := range ch{ // надо закрыть канал что бы не было дедлока
	fmt.Print(i)
	}
}

func say(world string) {
	fmt.Println(world)

}

func sayHello(exit chan int) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		exit <- i
	}
	close(exit) // вот тут закрыли

}
	
	
---select

package main

import (
	"fmt"
	"time"
	//"time"

)

func main() {

	data := make(chan int)
	exit := make(chan int)
	

	go func ()  {
		for i := 0; i < 10; i++ {
			fmt.Print(<-data)
		}
		exit <-0 // индикация об окончании
	}()
	selectOne(data,exit)
}


func selectOne(data, exit chan int) {
	x :=0
	for { //бесконечный цикл
		select {
		case data <-x: // если кто то читает данные из дата
			x += 1
		case <-exit: 
			fmt.Println("exit")
			return // выходим из цикла
		default:
			fmt.Println("waiting")
			time.Sleep(50* time.Millisecond)
		}
		
	}
}	
	
небольшая подсказка по использованию конструкции select - case:
чтение канала происходит в строке с оператором case. то есть:
 case <-ch // информация из канала ch "считывается" оператором case и канал становится пустым.

поэтому, если мы хотим считать информацию в переменную, мы должны объявить ее в строке case:
case x:=<-ch //таким образом вы сможете дальше пользоваться информацией из канала ch 
----------------------------------------------

задачка
	
Вам необходимо написать функцию calculator следующего вида:

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int
Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.

в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал вы должны отправить квадрат аргумента.
в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный) канал вы должны отправить результат умножения аргумента на 3.
в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.
Функция calculator должна быть неблокирующей, сразу возвращая управление. Ваша функция получит всего одно значение в один из каналов - получили значение, обработали его, завершили работу.

После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого не сделаете, то превысите предельное время работы.
	
	
	
package main

import (
	"fmt"
	//"time"

)
func main() {                                   
	ch1, ch2 := make(chan int), make(chan int)  
	stop := make(chan struct{})                 
	r := calculator(ch1, ch2, stop)      // сначала функция потом даные в канал       
	//ch1 <- 3
	ch2 <- 4
	//close(stop)                                 
	fmt.Println(<-r)                            
												
 }  
 
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		select {
		case x := <-firstChan:
			out <- x * x
		case x := <-secondChan:
			out <- x * 3
		case <-stopChan:
			return
		}
	}()

	return out
}


-----------------------------------------------------------------
	

Вам необходимо написать функцию calculator следующего вида:

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.

Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу. Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.

Функция calculator должна быть неблокирующей, сразу возвращая управление.

Выходной канал должен быть закрыт после выполнения всех оговоренных условий, если вы этого не сделаете, то превысите предельное время работы.


package main

import (
	"fmt"
	//"time"

)

func main(){

   arguments := make(chan int)
   done := make(chan struct{})
   result := calculator(arguments, done)
   for i := 0; i < 10; i++ {
      arguments <- 1
   }
   close(done)
   fmt.Println(<-result)

 }
 

 func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {

	out := make(chan int)
	
	go func() {

		defer close(out)
		var sum int
		for {
		select {
			case num := <-arguments:
				sum += num
			case <-done:
				out<-sum
				return
		}
	}
	}()

	return out
}

-------------------------

Не мог пройти тест пока не уяснил элементарную (для гофера) ситуацию:  
for i:=0; i<n; i++{ 
go func() {...
	}() 

} 
- не работает как того ждешь при попытке использовать i в горутине. Надо передавать i как параметр! Например
	
for i:=0; i<n; i++{ 
go func(i int) {...
		}(i) 
}. 

Так же и для for range. Может для вас, коллеги, это очевидно, но всякий случай...




----------------------
	задачка
Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).
Описание ее работы:
n раз сделать следующее
прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
вычислить f(x1) + f(x2)
записать полученное значение в out
Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.
Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.
Формат ввода:
количество итераций передается через аргумент n.
целые числа подаются через аргументы-каналы in1 и in2.
функция для обработки чисел перед сложением передается через аргумент fn.
Формат вывода:
канал для вывода результатов передается через аргумент out.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

)

const N = 20

func main() {

	fn := func(x int) int {
		time.Sleep(time.Duration(rand.Int31n(N)) * time.Second)
		return x * 2
	}
	in1 := make(chan int, N)
	in2 := make(chan int, N)
	out := make(chan int, N)

	start := time.Now()
	merge2Channels(fn, in1, in2, out, N+1)
	for i := 0; i < N+1; i++ {
		in1 <- i
		in2 <- i
	}

	orderFail := false
	EvenFail := false
	for i, prev := 0, 0; i < N; i++ {
		c := <-out
		if c%2 != 0 {
			EvenFail = true
		}
		if prev >= c && i != 0 {
			orderFail = true
		}
		prev = c
		fmt.Println(c)
	}
	if orderFail {
		fmt.Println("порядок нарушен")
	}
	if EvenFail {
		fmt.Println("Есть не четные")
	}
	duration := time.Since(start)
	if duration.Seconds() > N {
		fmt.Println("Время превышено")
	}
	fmt.Println("Время выполнения: ", duration)
}

// Merge2Channels below

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var (
		slice1 = make([]int, n)
		slice2 = make([]int, n)
		wg sync.WaitGroup
		mu sync.Mutex
	)
	wg.Add(n * 2)

	go func() {
		for i := 0; i < n; i++ {
			mu.Lock()
			slice1[i] = <-in1
			mu.Unlock()
			wg.Done()
		}
	}()
	go func() {
		for i := 0; i < n; i++ {
			mu.Lock()
			slice2[i] = <-in2
			mu.Unlock()
			wg.Done()
		}
	}()
	go func() {
		for i, v := range slice1 {
			slice1[i] = fn(v)
		}
	}()
	go func() {
		for i, v := range slice2 {
			slice2[i] = fn(v)
		}
	}()
	go func() {
		defer close(out)
		wg.Wait()
		for i := 0; i < n; i++ {
			out <- slice1[i] + slice2[i]
		}
	}()
}

Но тут, рутины где вызываются функции их никто не ждет
	
----------------------------	 



