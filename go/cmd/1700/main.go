package main

import (
	"fmt"
	"rcs/pkg/ds"
)

func main() {
	fmt.Println("Hello, World")
}

// 1. there's "Circular" or "Square" sandwiches (0, 1)
// 2. count(student) == count(sandwiches)
// 3. sandwiches is from a stack (can only pop)
// 4. when student likes it, pop(sandwich)
// 5. when student dislike, deq(student); enque(student)

// a. how to know the current sandwich stack is exhausted

func countStudents(students []int, sandwiches []int) int {
	studentQueue := ds.Queue[int]{}
	sandwichStack := ds.Stack[int]{}

	studentPref0 := 0
	studentPref1 := 0

	for _, studentVal := range students {
		studentQueue.Enque(studentVal)
		if studentVal == 0 {
			studentPref0 += 1
		} else {
			studentPref1 += 1
		}
	}

	for index := range sandwiches {
		sandwichStack.Push(sandwiches[len(sandwiches)-1-index])
	}

	for {
		if sandwichStack.Empty() {
			break
		}

		// break crit
		currTop := sandwichStack.Peek()
		if currTop == 0 && studentPref0 == 0 {
			break
		}
		if currTop == 1 && studentPref1 == 0 {
			break
		}

		toDeque := studentQueue.Deque()

		if sandwichStack.Peek() == toDeque {
			sandwichStack.Pop()
			switch toDeque {
			case 0:
				{
					studentPref0 -= 1
				}
			case 1:
				{
					studentPref1 -= 1
				}
			}
		} else {
			studentQueue.Enque(toDeque)
		}

	}

	return studentQueue.Len()
}
