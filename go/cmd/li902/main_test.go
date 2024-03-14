package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	assert.Exactly(
		t,
		10,
		KthSmallest(&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 7},
			},
			Right: &TreeNode{
				Val:   50,
				Left:  &TreeNode{Val: 20},
				Right: &TreeNode{Val: 100},
			},
		}, 4),
	)
}
