package main

import (
	"fmt"
	"rcs/pkg/ds"
)

func main() {
	fmt.Println("Hello, World")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	stack := ds.Stack[*TreeNode]{}
	stack.Push(root)

	for {
		if stack.Empty() {
			break
		}

		for {
			topOfStack := stack.Peek()
			if topOfStack.Left != nil {
				stack.Push(topOfStack.Left)
				continue
			}

			if topOfStack.Right != nil {
				stack.Push(topOfStack.Right)
				continue
			}

			break // ends up in leaf node
		}

		popped := stack.Pop()
		ans = append(ans, popped.Val)

		if !stack.Empty() {
			parentOfPopped := stack.Peek()

			switch popped {
			case parentOfPopped.Left:
				parentOfPopped.Left = nil
			case parentOfPopped.Right:
				parentOfPopped.Right = nil
			}
		}
	}

	return ans
}
