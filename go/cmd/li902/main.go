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

func KthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	stack := ds.Stack[*TreeNode]{}
	traversalResult := []int{}
	stack.Push(root)

	for {
		if stack.Empty() {
			break
		}

		for {
			topOfStack := stack.Peek()

			if topOfStack.Left != nil {
				stack.Push(topOfStack.Left)
			} else {
				break
			}
		}

		popped := stack.Pop()
		traversalResult = append(traversalResult, popped.Val)

		if len(traversalResult) == k {
			break
		}

		if !stack.Empty() {
			stack.Peek().Left = nil

			if popped.Right != nil {
				stack.Peek().Left = popped.Right
			}
		}

		if popped.Right != nil {
			stack.Push(popped.Right)
		}
	}

	return traversalResult[len(traversalResult)-1]
}
