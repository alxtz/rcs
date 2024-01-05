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
		[]int{3, 2, 1},
		postorderTraversal(&tree),
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
		postorderTraversal(&tree),
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
		[]int{10, 30, 20, 150, 300, 200, 100},
		postorderTraversal(&tree),
	)
}

func Test_ShouldAcceptNilRoot(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]int{},
		postorderTraversal(nil),
	)
}
