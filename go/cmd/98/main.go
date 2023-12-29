package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println("Hello, World 0")
}

// var count = 0

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var traversalMem = []int{}

	var stack = []*TreeNode{root}

	for {

		if len(stack) == 0 {
			break
		}

		if lastStackItem := stack[len(stack)-1]; lastStackItem.Left != nil {
			stack = append(stack, lastStackItem.Left)
			continue
		}

		var itemToPop = stack[len(stack)-1]

		stack = stack[:len(stack)-1]

		if itemToPop.Right != nil {
			if len(stack) > 0 {
				stack[len(stack)-1].Left = itemToPop.Right
			}
			stack = append(stack, itemToPop.Right)
		} else {
			if len(stack) > 0 {
				stack[len(stack)-1].Left = nil
			}
		}

		if len(traversalMem) == 0 {
			traversalMem = append(traversalMem, itemToPop.Val)
			continue
		}

		if itemToPop.Val <= traversalMem[len(traversalMem)-1] {
			return false
		} else {
			traversalMem = append(traversalMem, itemToPop.Val)
		}
	}

	return true
}
