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

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	stack := ds.Stack[*TreeNode]{}
	stack.Push(root)
	ans = append(ans, root.Val)

	for {
		if stack.Len() == 0 {
			break
		}

		for {
			topOfStack := stack.Peek()

			if topOfStack.Left != nil {
				stack.Push(topOfStack.Left)
				ans = append(ans, topOfStack.Left.Val)
				continue
			}

			if topOfStack.Right != nil {
				stack.Push(topOfStack.Right)
				ans = append(ans, topOfStack.Right.Val)
				continue
			}

			break /* leaf Node{nil, nil} */
		}

		popped := stack.Pop()

		if !stack.Empty() {
			parentOfPopped := stack.Peek()
			if parentOfPopped.Left == popped {
				parentOfPopped.Left = nil
			}
			if parentOfPopped.Right == popped {
				parentOfPopped.Right = nil
			}
		}
	}

	return ans
}
