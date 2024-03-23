package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func say(done chan<- struct{}, pending <-chan string, id int) {

	for phrase := range pending {
		for _, word := range strings.Fields(phrase) {
			fmt.Printf("Worker #%d says: %s...\n", id, word)
			dur := time.Duration(rand.Intn(100)) * time.Millisecond
			time.Sleep(dur)
		}
	}
	done <- struct{}{}
}

func main() {
	
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}

	pending := make(chan string)

	go func() {
		for _, phrase := range phrases {
			pending <- phrase
		}
		close(pending)
	}()

	done := make(chan struct{})

	go say(done, pending, 1)
	go say(done, pending, 2)

	<-done
	<-done
}
