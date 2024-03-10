package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")
}

type MyStack struct {
	Vector []int
}

func Constructor() MyStack {
	return MyStack{Vector: []int{}}
}

func (s *MyStack) Push(x int) {
	s.Vector = append(s.Vector, x)
}

func (s *MyStack) Pop() int {
	if len(s.Vector) == 0 {
		panic("cannot pop from an empty stack")
	}

	toPop := s.Vector[len(s.Vector)-1]

	s.Vector = s.Vector[:len(s.Vector)-1]

	return toPop
}

func (s *MyStack) Top() int {
	if len(s.Vector) == 0 {
		panic("cannot peek from an empty stack")
	}

	return s.Vector[len(s.Vector)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.Vector) == 0
}
