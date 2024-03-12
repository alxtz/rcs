package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	assert.Exactly(
		t,
		[]int{4, 2, 1, 3, 6, 5, 7},
		PreorderTraversal(&TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   6,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 7},
			},
		}),
	)
}
