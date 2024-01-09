package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// }

// func ConvertToRPN(expression []string) []string {
// 	var temp = SimpleToRPN(expression)
// 	return strings.Split(temp[0], " ")
// }

// func SimpleToRPN(expression []string) []string {
// 	if len(expression) == 0 {
// 		return []string{}
// 	}

// 	var input = expression

// 	var numStack = Stack{Slice: []string{}}
// 	var oprStack = Stack{Slice: []string{}}

// 	for {
// 		fmt.Println("debug", numStack, oprStack, input)

// 		if len(input) == 0 {
// 			break
// 		}

// 		var char = input[0]

// 		if char == "+" || char == "-" || char == "*" || char == "/" {
// 			var operatorToPush = char

// 			oprStack.Push(operatorToPush)
// 		} else if char == "(" {
// 			var leftPthCount = 0
// 			var rightPthCount = 0

// 			var rightPtsIndex int

// 			for i, val := range input {
// 				if val == "(" {
// 					leftPthCount++
// 					continue
// 				}
// 				if val == ")" {
// 					rightPthCount++

// 					if leftPthCount == rightPthCount {
// 						rightPtsIndex = i
// 						break
// 					}

// 					continue
// 				}
// 			}

// 			var bulk = input[1:rightPtsIndex]
// 			input = input[rightPtsIndex+1:]

// 			var r = strings.Join(SimpleToRPN(bulk), " ")

// 			fmt.Println("r", r)

// 			numStack.Push(r)

// 			continue
// 		} else {
// 			numStack.Push(char)

// 			var lengthOk = len(numStack.Slice) >= 2 && len(oprStack.Slice) >= 1
// 			if lengthOk {
// 				var topOfOprStack = oprStack.Slice[len(oprStack.Slice)-1]
// 				if topOfOprStack == "*" || topOfOprStack == "/" {
// 					var pop1st = numStack.Pop()
// 					var pop2nd = numStack.Pop()
// 					var bulk = pop2nd + " " + pop1st + " " + oprStack.Pop()
// 					numStack.Push(bulk)
// 				}
// 			}
// 		}

// 		input = input[1:]
// 	}

// 	var numQueue = Queue(numStack)
// 	var oprQueue = Queue(oprStack)

// 	for {
// 		var lengthOk = len(numQueue.Slice) >= 2 && len(oprQueue.Slice) >= 1
// 		if lengthOk {
// 			var pop1st = numQueue.Pop()
// 			var pop2nd = numQueue.Pop()
// 			var bulk = pop1st + " " + pop2nd + " " + oprQueue.Pop()
// 			numQueue.Slice = append([]string{bulk}, numQueue.Slice...) // prepend
// 		} else {
// 			break
// 		}
// 	}

// 	return numQueue.Slice
// }

// type Stack struct {
// 	Slice []string
// }

// func (s *Stack) Pop() string {
// 	var itemToPop = s.Slice[len(s.Slice)-1]
// 	s.Slice = s.Slice[:len(s.Slice)-1]

// 	return itemToPop
// }

// func (s *Stack) Push(val string) {
// 	s.Slice = append(s.Slice, val)
// }

// type Queue struct {
// 	Slice []string
// }

// func (q *Queue) Pop() string {
// 	var itemToPop = q.Slice[0]
// 	q.Slice = q.Slice[1:len(q.Slice)]

// 	return itemToPop
// }

// func (q *Queue) Push(val string) {
// 	q.Slice = append(q.Slice, val)
// }
