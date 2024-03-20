package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	assert.Exactly(
		t,
		[][]int{
			{3},
			{20, 9},
			{15, 7},
		},
		zigzagLevelOrder(&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}),
	)
}

func Test_ShouldDoB(t *testing.T) {
	assert.Exactly(
		t,
		[][]int{
			{1},
			{3, 2},
			{4, 5},
		},
		zigzagLevelOrder(&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 5},
			},
		}),
	)
}
