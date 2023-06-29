package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type element interface{}

type weightFunc func(element) int

type iterator interface {
	next() bool
	val() element
	
}

type intIterator struct {
	numbers []int
	position int
}

type ints struct {
	
}

func (r intIterator) next() bool {
	for _, v := range r.numbers {
		if v == 0 {
			return	false 
		}
	}
	return true
}
func (r intIterator) val() element {
    return r.numbers[r.position]
}

func newIntIterator(src []int) *intIterator {
	return &intIterator{src, 0}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
func main() {
	nums := readInput()
	fmt.Println("1")
	it := newIntIterator(nums) //есть итератор, созданный newIntIterator
	weight := func(el element) int { //принимает элемент последовательности, возврщ число
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
