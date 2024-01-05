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

type BSTIterator struct {
	Stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		Stack: []*TreeNode{root},
	}
}

func (itrr *BSTIterator) Next() int {
	if len(itrr.Stack) == 0 {
		return 0
	}

	for {
		var lastItemOfStack = itrr.Stack[len(itrr.Stack)-1]
		if lastItemOfStack.Left == nil {
			break
		} else {
			itrr.Stack = append(itrr.Stack, lastItemOfStack.Left)
			continue
		}
	}

	var itemToRemove = itrr.Stack[len(itrr.Stack)-1]

	// pop
	itrr.Stack = itrr.Stack[:len(itrr.Stack)-1]

	if len(itrr.Stack) > 0 {
		itrr.Stack[len(itrr.Stack)-1].Left = nil

		if itemToRemove.Right != nil {
			itrr.Stack[len(itrr.Stack)-1].Left = itemToRemove.Right
		}
	}

	if itemToRemove.Right != nil {
		itrr.Stack = append(itrr.Stack, itemToRemove.Right)
	}

	return itemToRemove.Val
}

func (itrr *BSTIterator) HasNext() bool {
	return len(itrr.Stack) != 0
}
