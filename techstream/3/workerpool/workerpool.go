// Более сложный пример, с использованием пула обработчиков для типовых задач
package main

import (
	"fmt"
	"sync"
)

// Task - описание интрефейса работы
type Task interface {
	Execute()
}

// Pool - структура, нам потребуется Мутекс, для гарантий атомарности изменений самого объекта

// публичная структура 
type Pool struct {
	mu    sync.Mutex
	size  int
	tasks chan Task // Канал входящих задач
	kill  chan struct{} // Канал отмены, для завершения работы
	wg    sync.WaitGroup // WaitGroup для контроля завершнеия работ
}

// Скроем внутреннее усройство за 
// конструктором, пользователь может влиять только на размер(size)пула
func NewPool(size int) *Pool {
	pool := &Pool{
		// Канал задач - буферизированный, 
		// чтобы основная программа не блокировалась при постановке задач
		tasks: make(chan Task, 128),
		// Канал kill для убийства "лишних воркеров"
		kill: make(chan struct{}),
	}
	// Вызовем метод resize, чтобы установить соответствующий размер пула
	pool.Resize(size)
	return pool
}
// Жизненный цикл воркера (умрет если новые задачи не поступают)
func (p *Pool) worker() {
	defer p.wg.Done() // счетчик декрементим
	for {
		select {
		// Если есть задача, то ее нужно обработать
		case task, ok := <-p.tasks:
			if !ok {
				return
			}
			task.Execute()
			// Если пришел сигнал умирать, выходим
		case <-p.kill:
			return
		}
	}
}

func (p *Pool) Resize(n int) {
	// Захватывам лок, чтобы 
	// избежать одновременного изменения состояния
	p.mu.Lock() // что бы заблокировать парралельные го рутины в этом месте
	defer p.mu.Unlock()
	for p.size < n {
		p.size++
		p.wg.Add(1) //  добавляем в группу ожидания еще одного воркера
		go p.worker() // cтартуем воркер
	}
	// дикремент счетчика
	//
	for p.size > n { 
		p.size-- // уменьшаем счетчик размера
		p.kill <- struct{}{} // посылаем в канал килл пустую структуру(минимальный контейнер)
	}
}

func (p *Pool) Close() { // закроем канал задач и в 45 строке не ОК
	close(p.tasks)
}

func (p *Pool) Wait() { // ждем завершение
	p.wg.Wait()
}

func (p *Pool) Exec(task Task) { //добавляем задачу на исполнение
	p.tasks <- task
}

type ExampleTask string

func (e ExampleTask) Execute() {
	fmt.Println("executing:", string(e))
}

func main() {
	pool := NewPool(5)

	pool.Exec(ExampleTask("foo"))
	pool.Exec(ExampleTask("bar"))

	pool.Resize(3)

	pool.Resize(6)

	for i := 0; i < 20; i++ {
		pool.Exec(ExampleTask(fmt.Sprintf("additional_%d", i+1)))
	}

	pool.Exec(ExampleTask("is"))

	pool.Close()//for commit 4

	pool.Wait()
}

