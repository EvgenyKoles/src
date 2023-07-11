package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// element represents an element in the sequence.
type element interface{}

// iterator represents an iterator over a sequence.
type iterator interface {
	next() bool
	val() element
}
// intIterator is an iterator over a sequence of integers.
type intIterator struct {
	nums []int
	pos  int
}
func newIntIterator(nums1 []int) *intIterator {
	return &intIterator{
		nums: nums1}
}
func (it *intIterator) next() bool {
	if it.pos < len(it.nums) {
		it.pos++
		return true
	}
	return false
}
func (it *intIterator) val() element {
	return it.nums[it.pos-1]
}
// weightFunc is a function that returns the weight of an element.
type weightFunc func(el element) int

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
func main() {
	nums := readInput()
	it := newIntIterator(nums)       //есть итератор, созданный newIntIterator
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
