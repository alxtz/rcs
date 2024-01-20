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

	toFinalCompute := ds.Stack[string]{Slice: []string{}}

	for _, val := range input {
		if val == ")" {
			collect := []string{}
			for {
				toPop := toFinalCompute.Pop()
				if toPop != "(" {
					collect = append([]string{toPop}, collect...)
				} else {
					break
				}
			}

			toFinalCompute.Push(fmt.Sprint(symbolsToInt(collect)))
		} else {
			toFinalCompute.Push(val)
		}
	}

	return symbolsToInt(toFinalCompute.Slice)
}

func symbolsToInt(symbols []string) int {
	var accu = 0
	var minusOn = false

	for _, val := range symbols {
		if val == "-" {
			minusOn = true
			continue
		}

		if val == "+" {
			continue
		}

		num, _ := strconv.Atoi(val)

		if minusOn {
			accu -= num
			minusOn = false
		} else {
			accu += num
		}
	}

	return accu
}
