// Один из механизмов синхронизации - каналы
// Каналы, это объект через который можно обеспечить взаимодействие нескольких горутин
// В принимающей (или возвращающей) канал функции, можно указать 
// направление работы с каналом
// Только для чтения - "<-chan" или только для записи "chan<-"
package main

import "fmt"

var c chan int

func main() {

	// Создаем канал в котором будет строки
	c := make(chan string)
	// стартуем пишущую горутину

	go greet(c)
	go sorry(c)

	for i := 0; i < 5; i++ {
		// Читаем 5 раз пару строк из канала
		fmt.Println(<-c, ",", <-c)
	}

	stuff := make(chan int, 7) // буферезированный канал. можно записать 7 значений. 
	// и только на 8м 
	// значении если никто не читал программа будет заблокированна в close
	// тут только читаем. записываем в process ?
	for i := 0; i < 19; i = i + 3 {
		stuff <- i
		fmt.Println(i)
	}
	close(stuff)

	fmt.Println("Res", process(stuff))

}

func greet(c chan<- string) {
	// Запускаем бесконечный цикл 
	for {
		// и пишем в канал пару строк	
		// Подпрограмма будет заблокирована до того, 
		// как кто-то захочет прочитать из канала
		c <- fmt.Sprintf("Владыка")
		c <- fmt.Sprintf("Штурмовик")
	}
}

func sorry (c chan<- string) {
	for {
		c <-fmt.Sprint("Dich") 
	}
}
//пробежатся по каналу, читаем из канала в input и пробегаемся - range input  
func process(input <-chan int) (res int) {
	for r := range input {
		res += r
	}
	return
}


//   dataChannel := make(chan string, 3)
//   dataChannel <- "Some Sample Data"
//   dataChannel <- "Some Other Sample Data"
//   dataChannel <- "Buffered Channel"
//   fmt.Println(<-dataChannel)
//   fmt.Println(<-dataChannel)
//   fmt.Println(<-dataChannel)
				