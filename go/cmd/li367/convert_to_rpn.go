package main

import (
	ds "rcs/pkg/ds"
)

var precMap = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func ConvertToRPN(infixExpr []string) (postfixExpr []string) {
	var input = infixExpr

	var stack = ds.Stack[string]{Slice: []string{}}
	var queue = ds.Queue[string]{Slice: []string{}}

	for {
		if len(input) == 0 {
			for len(stack.Slice) > 0 {
				queue.Enque(stack.Pop())
			}

			break
		}

		var ele = input[0]
		input = input[1:]

		if _, ok := precMap[ele]; ok {
			oprToPush := ele

			for len(stack.Slice) > 0 {
				stackTopOpr := stack.Slice[len(stack.Slice)-1]

				topOprPrec := precMap[stackTopOpr]
				currOprPrec := precMap[oprToPush]

				if currOprPrec <= topOprPrec {
					queue.Enque(stack.Pop())
					continue
				} else {
					break
				}
			}

			stack.Push(oprToPush) // push the opr into stack
		} else if ele == "(" {
			stack.Push(ele)
			continue
		} else if ele == ")" {
			for len(stack.Slice) > 0 {
				toPop := stack.Slice[len(stack.Slice)-1]

				if toPop == "(" {
					stack.Pop()
					break
				} else {
					queue.Enque(stack.Pop())
				}
			}
		} else {
			// ele is number
			queue.Enque(ele)
		}
	}

	postfixExpr = queue.Slice
	return postfixExpr
}
