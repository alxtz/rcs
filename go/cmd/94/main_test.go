package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnCorrectly_WhenGivenThreeNodeTree(t *testing.T) {
	// t.Skip()
	var tree = TreeNode{
		Val:  1,
		Left: nil, Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3}, Right: nil,
		},
	}

	assert.Exactly(
		t,
		[]int{1, 3, 2},
		inorderTraversal(&tree),
	)
}

func Test_ShouldReturnCorrectly_WhenGivenOneNodeTree(t *testing.T) {
	// t.Skip()
	var tree = TreeNode{
		Val: 1,
	}

	assert.Exactly(
		t,
		[]int{1},
		inorderTraversal(&tree),
	)
}

func Test_ShouldReturnCorrectly_WhenGivenBinarySearchTree10To30(t *testing.T) {
	// t.Skip()

	var tree = TreeNode{
		Val: 100,
		Left: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 10},
			Right: &TreeNode{Val: 30},
		},
		Right: &TreeNode{
			Val:   200,
			Left:  &TreeNode{Val: 150},
			Right: &TreeNode{Val: 300},
		},
	}

	assert.Exactly(
		t,
		[]int{10, 20, 30, 100, 150, 200, 300},
		inorderTraversal(&tree),
	)
}

func Test_ShouldAcceptNilRoot(t *testing.T) {

	assert.Exactly(
		t,
		[]int{},
		inorderTraversal(nil),
	)
}
