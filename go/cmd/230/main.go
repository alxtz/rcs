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
	fmt.Println("Hello, World", kthSmallest(nil, 0))
}

func kthSmallest(root *TreeNode, k int) int {
	var result = inorderTraversal(root, k)
	return result[len(result)-1]
}

func inorderTraversal(root *TreeNode, limit int) []int {
	if root == nil {
		return []int{}
	}

	var ans = []int{}
	var stack = []*TreeNode{root}

	for {
		if len(ans) == limit {
			break
		}

		if len(stack) == 0 {
			break
		}

		for {
			if lastItem := stack[len(stack)-1]; lastItem.Left != nil {
				stack = append(stack, lastItem.Left)
				continue
			} else {
				break
			}
		}

		var itemToPop = stack[len(stack)-1]
		// pop
		stack = stack[:len(stack)-1]

		ans = append(ans, itemToPop.Val)

		if len(stack) > 0 {
			var tail = stack[len(stack)-1]
			tail.Left = nil

			if itemToPop.Right != nil {
				tail.Left = itemToPop.Right
			}
		}

		if itemToPop.Right != nil {
			stack = append(stack, itemToPop.Right)
		}
	}

	return ans
}
