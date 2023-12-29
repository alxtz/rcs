package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World 0")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var traversalResults = []int{}

	var stack []*TreeNode = []*TreeNode{root}

	for {
		if len(stack) == 0 {
			break
		}

		for {
			var lastOfStack = stack[len(stack)-1]

			if lastOfStack.Left != nil {
				stack = append(stack, lastOfStack.Left)
			} else {
				break
			}
		}

		var prevOfStack = stack[len(stack)-1]

		traversalResults = append(traversalResults, prevOfStack.Val)
		stack = stack[:len(stack)-1] // pop

		if len(stack) != 0 {
			var latestLastOfStack = stack[len(stack)-1]
			latestLastOfStack.Left = nil
		}

		if prevOfStack.Right != nil {
			stack = append(stack, prevOfStack.Right)
		}
	}

	return traversalResults
}
