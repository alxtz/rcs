package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	assert.Exactly(
		t,
		[]int{1, 3, 2, 5, 7, 6, 4},
		PostorderTraversal(&TreeNode{
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
