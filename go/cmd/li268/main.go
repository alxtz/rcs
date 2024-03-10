package main

import (
	"fmt"
	"rcs/pkg/ds"
	"strconv"
)

func main() {
	fmt.Println("Hello, World")
}

/**
 * @param s: a string with "(" and ")"
 * @return: return the score of the string
 */
func ParenthesesScore(s string) int {
	// stack item is either "(", ")", or "number"
	stack := ds.Stack[string]{}

	for _, charCode := range s {
		char := string(charCode)

		if char == "(" {
			stack.Push(char)
			continue
		}

		// char == ")"
		// unpacking process
		var numToPush int = -1

		for {
			if stack.Peek() == "(" {
				if numToPush == -1 {
					numToPush = 1
				} else {
					numToPush = numToPush * 2
				}

				stack.Pop()
				break
			}

			// stack.peek is "number"

			if numToPush == -1 {
				numToPush, _ = strconv.Atoi(stack.Pop())
				continue
			} else {
				newFoundNum, _ := strconv.Atoi(stack.Pop())
				numToPush += newFoundNum
				continue
			}

		}

		if stack.Len() > 0 {
			topStack := stack.Peek()

			// packs nums, if there's any
			if topStack != "(" && topStack != ")" {
				// top of stack is a number
				toPop, _ := strconv.Atoi(stack.Pop())
				numToPush += toPop
			}
		}

		stack.Push(strconv.Itoa(numToPush))
	}

	result, _ := strconv.Atoi(stack.Peek())
	return result
}
