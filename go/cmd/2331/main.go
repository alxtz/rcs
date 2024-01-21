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

func evaluateTree(root *TreeNode) bool {
	val := root.Val

	if val == 0 {
		return false
	}

	if val == 1 {
		return true
	}

	if val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}

	if val == 3 {
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}

	return false
}
