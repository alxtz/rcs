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

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}

	stack := ds.Stack[*TreeNode]{}
	stack.Push(root)

	for {
		if stack.Len() == 0 {
			break
		}

		for {
			if stack.Peek().Left != nil {
				stack.Push(stack.Peek().Left)
			} else {
				break
			}
		}

		leftMostLeaf := stack.Pop()
		ans = append(ans, leftMostLeaf.Val)

		if leftMostLeaf.Right != nil {
			if stack.Len() > 0 {
				stack.Peek().Left = leftMostLeaf.Right
			}
			stack.Push(leftMostLeaf.Right)
		} else {
			if stack.Len() > 0 {
				stack.Peek().Left = nil
			}
		}
	}

	return ans
}
