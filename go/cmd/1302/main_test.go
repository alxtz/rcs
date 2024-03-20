package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	assert.Exactly(
		t,
		15,
		deepestLeavesSum(&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 7},
				},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 6,
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		}),
	)
}
