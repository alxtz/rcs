package main

import (
	ds "complicated-prob-solving"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Hello, World")
}

func parsing(s string) []string {
	var inputStack = ds.Stack[string]{Slice: []string{}}

	for _, charCode := range s {
		char := string(charCode)
		if char == " " {
			continue
		} else if slices.Contains([]string{"+", "-", "(", ")"}, char) {
			inputStack.Push(char)
		} else {
			lastIsNum := inputStack.Len() > 0 && !slices.Contains([]string{"+", "-", "(", ")"}, inputStack.Peek())

			if lastIsNum {
				inputStack.Push(inputStack.Pop() + char)
			} else {
				inputStack.Push(char)
			}
		}
	}

	return inputStack.Slice
}

func calculate(s string) int {
	var input = parsing(s)

	oprStack := ds.Stack[string]{Slice: []string{}}
	numStack := ds.Stack[int]{Slice: []int{}}

	shouldPack := false

	for _, val := range input {
		if val == " " {
			continue
		}

		fmt.Println(ds.Readable(oprStack.Slice))
		fmt.Println(ds.Readable(numStack.Slice))
		fmt.Println("---")

		if val == "+" || val == "-" {
			oprStack.Push(val)
		} else if val == "(" {
			oprStack.Push(val)
		} else if val == ")" {
			oprStack.Pop()
			shouldPack = true
		} else {
			num, _ := strconv.Atoi(val)
			numStack.Push(num)
			shouldPack = true
		}

		for shouldPack {
			fmt.Println("packing", val)
			if (oprStack.Len() > 0 && oprStack.Peek() == "(") || len(oprStack.Slice) == 0 {
				shouldPack = false
				continue
			}

			if numStack.Len() == 0 {
				shouldPack = false
				continue
			}

			if numStack.Len() >= 2 {
				rightVal := numStack.Pop()
				leftVal := numStack.Pop()
				opr := oprStack.Pop()

				var val int

				if opr == "+" {
					val = leftVal + rightVal
				} else {
					val = leftVal - rightVal
				}

				numStack.Push(val)
				continue
			}

			if numStack.Len() == 1 {
				rightVal := numStack.Pop()
				oprStack.Pop()

				// var val int

				// if opr == "+" {
				// 	val = leftVal + rightVal
				// } else {
				// 	val = leftVal - rightVal
				// }

				numStack.Push(-rightVal)
				shouldPack = false
				continue
			}

		}
	}

	return numStack.Slice[0]
}
