package main

import (
	ds "complicated-prob-solving"
)

func main() {
	// fmt.Println("Hello, World")
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

	var treeNodeStack = ds.Stack[*ExpressionTreeNode]{Slice: []*ExpressionTreeNode{}}

	var stack = ds.Stack[string]{Slice: []string{}}
	var queue = ds.Queue[string]{Slice: []string{}}

	for {
		if len(queue.Slice) >= 2 && len(stack.Slice) >= 1 {
			opr := stack.Pop()
			val1 := queue.Deque()
			val2 := queue.Deque()

			newBetNode := ExpressionTreeNode{
				Symbol: opr,
				Left:   &ExpressionTreeNode{Symbol: val1},
				Right:  &ExpressionTreeNode{Symbol: val2},
			}

			treeNodeStack.Push(&newBetNode)
			continue
		}

		if len(treeNodeStack.Slice) >= 2 && len(stack.Slice) >= 1 {

			opr := stack.Pop()
			val1 := treeNodeStack.Pop()
			val2 := treeNodeStack.Pop()

			newBetNode := ExpressionTreeNode{
				Symbol: opr,
				Left:   val2,
				Right:  val1,
			}

			treeNodeStack.Push(&newBetNode)

			continue
		}

		if len(queue.Slice) >= 1 && len(treeNodeStack.Slice) >= 1 && len(stack.Slice) >= 1 {
			opr := stack.Pop()
			val1 := queue.Deque()
			val2 := treeNodeStack.Pop()

			newBetNode := ExpressionTreeNode{
				Symbol: opr,
				Left:   &ExpressionTreeNode{Symbol: val1},
				Right:  val2,
			}

			treeNodeStack.Push(&newBetNode)
			continue
		}

		if len(input) == 0 && len(stack.Slice) == 0 && len(queue.Slice) == 0 && len(treeNodeStack.Slice) == 1 {
			break
		}

		if len(input) == 0 {
			continue
		}

		var ele = input[0]
		input = input[1:]

		if _, ok := precMap[ele]; ok {
			stack.Push(ele)
		} else {
			queue.Enque(ele)
		}

	}

	if len(treeNodeStack.Slice) == 1 {
		return treeNodeStack.Slice[0]
	} else {
		return nil
	}
}
