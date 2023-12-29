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

// 猜想，只在 stack 長度變長時，才印出東西？
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var traversalResults = []int{root.Val}

	var stack = []*TreeNode{
		root,
	}

	for {
		if len(stack) == 0 {
			break
		}

		var lastItemOfStack = stack[len(stack)-1]

		if lastItemOfStack.Left != nil {
			stack = append(stack, lastItemOfStack.Left)

			traversalResults = append(traversalResults, (stack[len(stack)-1]).Val)
			continue
		}

		// pop()
		stack = stack[:len(stack)-1]

		var justPoppedItem = lastItemOfStack
		var poppedItemHasChild = lastItemOfStack.Right != nil

		if poppedItemHasChild {
			if len(stack) > 0 {
				// the prev parent of the left-most leaf node, would grab the right items of prev left-most leaf node
				(stack[len(stack)-1]).Left = justPoppedItem.Right
			}
			stack = append(stack, justPoppedItem.Right)
			traversalResults = append(traversalResults, (stack[len(stack)-1]).Val)
		} else {
			if len(stack) > 0 {
				(stack[len(stack)-1]).Left = nil
			}
		}
	}

	return traversalResults
}
