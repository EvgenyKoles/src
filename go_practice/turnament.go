


// футбольный турнир
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// result представляет результат матча
type result byte

// возможные результаты матча
const (
	win  result = 'W'
	draw result = 'D'
	loss result = 'L'
)

// team представляет команду
type team byte

// match представляет матч
// например, строке BAW соответствует
// first=B, second=A, result=W
type match struct {
	first  team
	second team
	result result
}

// rating представляет турнирный рейтинг команд -
// количество набранных очков по каждой команде
type rating map[team]int

// tournament представляет турнир
type tournament []match //содержит слайс матчей

// calcRating считает и возвращает рейтинг турнира
func (trn *tournament) calcRating() rating {

	m := rating{}
	for _, value := range *trn {
		switch value.result {

		case win :
			m[value.first] += 3
			m[value.first] += 0
		
		case loss:
			m[value.first] += 0
			m[value.second] += 3
		
		case draw:
			m[value.first] += 1
			m[value.second] += 1
		}
	}
	
	return m
	
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// код, который парсит результаты игр, уже реализован
// код, который печатает рейтинг, уже реализован
// ваша задача - реализовать недостающие структуры и методы выше
func main() {
	src := readString()
	trn := parseTournament(src)
	fmt.Println(trn)
	rt := trn.calcRating()
	rt.print()
}

// readString считывает и возвращает строку из os.Stdin
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return str
}

// parseTournament парсит турнир из исходной строки
func parseTournament(s string) tournament {
	descriptions := strings.Split(s, " ") // делим большую строку
	trn := tournament{}                   //создаем турнир
	for _, descr := range descriptions {  // пробегаемся по слайсу разделенной строки
		m := parseMatch(descr) // парсим каждые три буквы
		trn.addMatch(m)        // добавляем матч к турниру
	}
	return trn
}

// parseMatch парсит матч из фрагмента исходной строки
func parseMatch(s string) match {
	return match{
		first:  team(s[0]),
		second: team(s[1]),
		result: result(s[2]),
	}
}

// addMatch добавляет матч к турниру
func (t *tournament) addMatch(m match) {
	*t = append(*t, m)
}

// print печатает результаты турнира
// в формате Aw Bx Cy Dz
func (r *rating) print() {
	var parts []string
	for team, score := range *r {
		part := fmt.Sprintf("%c%d", team, score)
		parts = append(parts, part)
	}
	sort.Strings(parts)
	fmt.Println(strings.Join(parts, " "))
}


//-------------------------------------