package main

import (
	"fmt"
	"rcs/pkg/ds"
)

func main() {
	fmt.Println("Hello, World")
}

func AsteroidCollision(asteroids []int) []int {
	stack := ds.Stack[int]{}

	for _, ast := range asteroids {
		isRight := RightCheck(ast)

		if isRight {
			stack.Push(ast)
			continue
		}

		// is left ast "<-", -x

		for {
			if stack.Len() == 0 {
				stack.Push(ast)
				break
			}

			topStackItem := stack.Peek()

			if RightCheck(topStackItem) {
				collideResult := topStackItem + ast

				if collideResult > 0 {
					// ast(left) to push is smaller, got destroyed
					break
				} else if collideResult < 0 {
					stack.Pop()
					continue
				} else {
					// collideResult == 0
					stack.Pop()
					break
				}
			} else {
				stack.Push(ast) // expect ast to all flying left
				break
			}
		}
	}

	return stack.Slice
}

func RightCheck(ast int) bool {
	return ast > 0
}
