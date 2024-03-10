package main

import (
	"fmt"
	ds "rcs/pkg/ds"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Hello, World")
}

func parsing(s string) []string {
	stack := ds.Stack[string]{}

	for _, charCode := range s {
		char := string(charCode)
		if char != " " {
			if slices.Contains([]string{"+", "-", "*", "/"}, char) {
				stack.Push(char)
				continue
			}

			if stack.Len() > 0 {
				isPeekNum := !slices.Contains([]string{"+", "-", "*", "/"}, stack.Peek())

				if isPeekNum {
					stack.Push(stack.Pop() + char)
					continue
				}
			}

			// else push
			stack.Push(char)
		}
	}

	return stack.Slice
}

func calculate(s string) int {
	input := parsing(s)

	oprStack := ds.Stack[string]{}
	numStack := ds.Stack[int]{}

	// var toPush string

	for i, val := range input {
		toPush := val
		isOpr := slices.Contains([]string{"+", "-", "*", "/"}, val)

		if isOpr {
			oprStack.Push(toPush)
		} else {
			num, _ := strconv.Atoi(toPush)
			numStack.Push(num)
		}

		if pushedNum := !isOpr; pushedNum {
			if oprStack.Len() >= 1 && slices.Contains([]string{"*", "/"}, oprStack.Peek()) {
				// had just push a num, with existing mult/div opr
				rightVal := numStack.Pop()
				leftVal := numStack.Pop()

				oprToUse := oprStack.Pop()

				if oprToUse == "*" {
					numStack.Push(leftVal * rightVal)
					// continue
				} else {
					numStack.Push(div(leftVal, rightVal))
					// continue
				}
			}
		}

		shouldPack := false

		if oprStack.Len() >= 2 && isOpr {
			if slices.Contains([]string{"+", "-"}, toPush) {
				shouldPack = true
			}
		}

		// fmt.Println("packing", i)
		if i == len(input)-1 && numStack.Len() >= 2 {
			shouldPack = true
		}

		if shouldPack {
			// fmt.Println("packing", numStack.Slice)
			// opr := oprStack.Pop()
			opr := oprStack.Slice[0]
			oprStack.Slice = oprStack.Slice[1:len(oprStack.Slice)]
			rightVal := numStack.Pop()
			leftVal := numStack.Pop()

			var computed int
			if opr == "+" {
				computed = leftVal + rightVal
			} else {
				computed = leftVal - rightVal
			}

			numStack.Push(computed)
		}
	}

	// fmt.Println(numStack)

	return numStack.Pop()
}

func div(rawLeftVal int, rightVal int) int {
	remainder := rawLeftVal % rightVal

	leftVal := rawLeftVal - remainder

	return leftVal / rightVal
}
