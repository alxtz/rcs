package main

import (
	"rcs/pkg/ds"
	"slices"
)

func main() {
	// stack := ds.Stack[string]{}
	// stack.Push("1")
	// stack.Pop()
	// stack.Push("2")
	// fmt.Println("Hello, World", stack.Peek(), stack.Len())
}

var leftSymbols = []string{"(", "[", "{"}

var rightToLeftSymbolMap = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func isValid(s string) bool {
	stack := ds.Stack[string]{}

	for _, charCode := range s {
		char := string(charCode)

		if slices.Contains(leftSymbols, char) {
			stack.Push(char)
			continue
		}

		// pushing a right symbol
		if stack.Len() == 0 {
			return false
		}

		currTop := stack.Peek()
		expectedTop := rightToLeftSymbolMap[char]

		if currTop == expectedTop {
			stack.Pop()
		} else {
			return false
		}
	}

	return stack.Len() == 0
}
