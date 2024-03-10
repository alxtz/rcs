package main

import ds "rcs/pkg/ds"

func main() {
}

type ExpressionTreeNode struct {
	Symbol string
	Left   *ExpressionTreeNode
	Right  *ExpressionTreeNode
}

func Build(infixExpr []string) *ExpressionTreeNode {
	var rpn = ConvertToRPN(infixExpr)
	return PostfixToBet(rpn)
}

func PostfixToBet(postfixExpr []string) *ExpressionTreeNode {
	if len(postfixExpr) == 0 {
		return nil
	}

	var input = postfixExpr

	var stack ds.Stack[*ExpressionTreeNode]

	for len(input) > 0 {
		var ele = input[0]
		input = input[1:]

		if _, isOpr := precMap[ele]; isOpr {
			val1 := stack.Pop()
			val2 := stack.Pop()

			stack.Push(&ExpressionTreeNode{
				Symbol: ele,
				Left:   val2,
				Right:  val1,
			})
		} else {
			stack.Push(&ExpressionTreeNode{
				Symbol: ele,
			})
		}
	}

	return stack.Slice[0]
}
