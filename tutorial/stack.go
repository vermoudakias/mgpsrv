package main

import (
	"fmt"
)

type stackInt struct {
	slice []int
}

func (s *stackInt) init() {
	array := make([]int, 0, 20)
	s.slice = array[:]
}

func (s *stackInt) push(v int) {
	s.slice = append(s.slice, v)
}

func (s *stackInt) pop() int {
	e := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return e
}

func stackIntSample() {
	var stack stackInt
	stack.init()
	fmt.Printf("%s\n", stack)
	stack.push(34)
	stack.push(42)
	fmt.Printf("%s\n", stack)
	fmt.Printf("%d, %d\n", stack.pop(), stack.pop())
	fmt.Printf("%s\n", stack)
}

