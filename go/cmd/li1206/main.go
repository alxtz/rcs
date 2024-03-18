package main

import (
	"fmt"
	"rcs/pkg/ds"
)

func main() {
	fmt.Println("Hello, World")
}

func NextGreaterElement(nums1 []int, num2List []int) []int {
	lookMap := make(map[int]int)
	stack := ds.Stack[int]{}

	for {
		if len(num2List) == 0 {
			break
		}

		tail := num2List[len(num2List)-1]
		num2List = num2List[:len(num2List)-1] // del last ele

		// about to push tail
		for {
			if stack.Empty() {
				break
			}
			fmt.Println("stack.for", stack.Slice)

			if currTop := stack.Peek(); currTop < tail {
				stack.Pop()
			} else {
				break
			}
		}

		if stack.Empty() {
			lookMap[tail] = -1
		} else {
			lookMap[tail] = stack.Peek()
		}

		stack.Push(tail)
	}

	ans := []int{}

	for _, num := range nums1 {
		val := lookMap[num]
		ans = append(ans, val)
	}

	return ans
}
