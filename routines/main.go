package main

import (
	"fmt"
	"strings"
	//"sync"
	"unicode"
)

// nextFunc возвращает следующее слово из генератора
type nextFunc func() string

// counter хранит количество цифр в каждом слове.
// ключ карты - слово, а значение - количество цифр в слове.
type counter map[string]int

// pair хранит слово и количество цифр в нем
type pair struct {
	word  string
	count int
}

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWords(next nextFunc) counter {

	pending := make(chan string) //читатель
	counted := make(chan pair)   //счетовод

	
	//pair := pair{}

	//var wg sync.WaitGroup

	go func() {
     
		for {
		word := next()
		
           	if word == "" {
			close(pending) // как только кончаются слова на след горутину
			return
		}
		pending	<-word 
	}

	}()

	go func() {
		for word := range pending {
			count := countDigits(word)
			counted <- pair{word, count}
		}
		close(counted) // закончились слова переходим на main горутину
	}()

	
	stats := counter{}
	for p := range counted {
		stats[p.word] = p.count
	}
	return stats
}

// countDigits возвращает количество цифр в строке
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats печатает слова и количество цифр в каждом
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordGenerator возвращает генератор, который выдает слова из фразы
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {

	phrase := "one two three four"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}