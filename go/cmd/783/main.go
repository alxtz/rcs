package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	var numList = inorderTraversal(root)
	// fmt.Println(numList)

	var minDiff = -1

	for index, val := range numList {
		if index == 0 {
			continue
		}

		var currDiff = val - numList[index-1]

		if minDiff == -1 {
			minDiff = currDiff
			continue
		}

		if currDiff < minDiff {
			minDiff = currDiff
		}
	}

	return minDiff
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var ansList = []int{}
	var stack = []*TreeNode{root}

	for {
		if len(stack) == 0 {
			break
		}

		for {
			if lastItem := stack[len(stack)-1]; lastItem.Left != nil {
				stack = append(stack, lastItem.Left)
				continue
			}
			break
		}

		var itemToPop = stack[len(stack)-1]

		// pop()
		stack = stack[:len(stack)-1]

		ansList = append(ansList, itemToPop.Val)

		if len(stack) > 0 {
			newLastItem := stack[len(stack)-1]
			newLastItem.Left = nil
			if itemToPop.Right != nil {
				newLastItem.Left = itemToPop.Right
			}
		}

		if itemToPop.Right != nil {
			stack = append(stack, itemToPop.Right)
		}
	}

	return ansList
}
