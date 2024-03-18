package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[][]int{
			{10},
			{20, 30},
			{40, 50, 60, 70},
		},
		levelOrder(&TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 40,
				},
				Right: &TreeNode{
					Val: 50,
				},
			},
			Right: &TreeNode{
				Val: 30,
				Left: &TreeNode{
					Val: 60,
				},
				Right: &TreeNode{
					Val: 70,
				},
			},
		}),
	)
}
