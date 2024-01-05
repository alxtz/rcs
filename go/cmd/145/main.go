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
	fmt.Println("Hello, World")
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var ans = []int{}
	var stack = []*TreeNode{root}

	for {
		if len(stack) == 0 {
			break
		}

		for {
			var tail = stack[len(stack)-1]
			var leftLeg = tail.Left
			var rightLeg = tail.Right

			if leftLeg == nil && rightLeg == nil {
				break
			}

			if leftLeg != nil {
				stack = append(stack, leftLeg)
				continue
			}

			if rightLeg != nil {
				stack = append(stack, rightLeg)
				continue
			}
		}

		var itemToPop = stack[len(stack)-1]

		// pop()
		stack = stack[:len(stack)-1]

		// write ans
		ans = append(ans, itemToPop.Val)

		// set parent leg pointer to nil, if possible
		var hasParent = len(stack) > 0
		if hasParent {
			var parent = stack[len(stack)-1]

			if parent.Left != nil {
				parent.Left = nil
			} else if parent.Right != nil {
				parent.Right = nil
			}
		}
	}

	return ans
}
