package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")
}

type MyQueue struct {
	Slice []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.Slice = append(q.Slice, x)
}

func (q *MyQueue) Pop() int {
	toPop := q.Slice[0]
	q.Slice = q.Slice[1:]
	return toPop
}

func (q *MyQueue) Peek() int {
	return q.Slice[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.Slice) == 0
}
