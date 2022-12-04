package main

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	// Bind на порт ОС
	listener, _ := net.Listen("tcp", ":5000")// открываем сокет, сюда сваливаются соединения, 
	//мы должны взять их из очереди
	

	for {
		// ждём пока не придёт клиент
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Can not connect!!")
			conn.Close()
			continue
		}
		fmt.Println("Connected")//успешно соединилисj

		// создаём Reader для чтения информации из сокета
		// далее читаем
		bufReader := bufio.NewReader(conn) 
		fmt.Println("Start reading")

		//выносим чтение в одтельный поток. при запуске второго телнета не принимает сообщения
		go func (c net.Conn) {
			defer conn.Close()

		for {
			// побайтово читаем
			rbyte, err := bufReader.ReadByte()

			if err != nil {
				fmt.Println("Can not read!", err) // если не смогли прочитать
				break
				}
			
			fmt.Print(string(rbyte))
			}
		} (conn) // передаем коннект в функцию
	}
}