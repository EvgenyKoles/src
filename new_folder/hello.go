package main

import (
	"fmt"
	//"math/rand"
	"sync"
	"time"
)

const N = 20

func main() {

	fn := func(x int) int {
		//time.Sleep(2 * time.Second)
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

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	var (
		slice1 = make([]int, n)
		slice2 = make([]int, n)
		wg     sync.WaitGroup
		mu     sync.Mutex
	)
	wg.Add(n * 4)

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
			mu.Lock()
			slice1[i] = fn(v)
			mu.Unlock()
			wg.Done()
		}
	}()
	go func() {
		for i, v := range slice2 {
			mu.Lock()
			slice2[i] = fn(v)
			mu.Unlock()
			wg.Done()
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
