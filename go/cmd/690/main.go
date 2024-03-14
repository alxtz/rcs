package main

import (
	"fmt"
	"rcs/pkg/ds"
)

func main() {
	fmt.Println("Hello, World")

	// getImportance([]*Employee{
	// 	&Employee{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
	// 	&Employee{Id: 2, Importance: 3, Subordinates: []int{}},
	// 	&Employee{Id: 3, Importance: 3, Subordinates: []int{}},
	// }, 1)
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	stack := ds.Stack[*Employee]{}
	employeeMap := map[int]*Employee{}

	for _, employee := range employees {
		employeeMap[employee.Id] = employee
	}

	stack.Push(employeeMap[id])

	sum := 0

	for {
		if stack.Empty() {
			break
		}

		for {
			topOfStack := stack.Peek()
			if len(topOfStack.Subordinates) != 0 {
				id := topOfStack.Subordinates[0]
				stack.Push(employeeMap[id])
			} else {
				break
			}
		}

		popped := stack.Pop()
		sum += popped.Importance

		if !stack.Empty() {
			parentOfPopped := stack.Peek()
			parentOfPopped.Subordinates = parentOfPopped.Subordinates[1:]
		}
	}

	return sum
}
