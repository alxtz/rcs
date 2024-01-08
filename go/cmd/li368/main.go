package main

import (
	"encoding/json"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Hello, World")
}

func EvaluateExpression(expression []string) int {
	return 1
}

type ExpressionTreeNode struct {
	Symbol []string
	Left   *ExpressionTreeNode
	Right  *ExpressionTreeNode
}

func Build(expression []string) *ExpressionTreeNode {
	// var expStr = strings.Join(expression, "")

	var rootNode *ExpressionTreeNode = nil

	// fmt.Println("exp", expStr)

	rootNode = TreeBuild(expression)

	j, _ := json.MarshalIndent(rootNode, "", "  ")
	fmt.Println(string(j))

	return rootNode
}

// 關於左右括號，勢必要用 stack 來解
// 關於先乘除後加減，可以把他畫約成 continuous string

func TreeBuild(expression []string) *ExpressionTreeNode {
	if len(expression) == 1 {
		return &ExpressionTreeNode{
			Left:   nil,
			Symbol: expression,
			Right:  nil,
		}
	}

	// fmt.Println("exp", expression, len(expression))

	// var head = []string{expression[0]}
	// var symbol = []string{expression[1]}
	// var tail = expression[2:]

	fmt.Println("using", expression)

	var remainingExpression = slices.Clone(expression)

	var head []string
	var tail []string
	var singularSymbol string

	for {
		if len(remainingExpression) == 1 || len(remainingExpression) == 0 {
			head = remainingExpression
			break
		}

		var first = remainingExpression[0]
		var second = remainingExpression[1]

		if isNum(first) {
			head = append(head, first)
			remainingExpression = remainingExpression[1:]
			continue
		}

		if first == "*" || first == "/" {
			head = append(head, first)
			head = append(head, second)
			remainingExpression = remainingExpression[2:]
			continue
		}

		if first == "+" || first == "-" {
			break
		}
	}

	// if len(head) > 1 {
	// 	var accu = 1
	// 	for len(head) > 0 {
	// 		var first = remainingExpression[0]
	// 		var second = remainingExpression[1]
	// 		var third = remainingExpression[2]
	// 	}
	// }

	if len(remainingExpression) == 0 {
		return &ExpressionTreeNode{
			Symbol: head,
		}
	}

	singularSymbol = remainingExpression[0]
	tail = remainingExpression[1:]

	return &ExpressionTreeNode{
		Left:   TreeBuild(head),
		Symbol: []string{singularSymbol},
		Right:  TreeBuild(tail),
	}
}

func isNum(s string) bool {
	return !slices.Contains([]string{"+", "-", "*", "/"}, s)
}
