package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleTree = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 6,
					},
				},
			},
		},
	},
	Right: &TreeNode{
		Val:  9,
		Left: &TreeNode{Val: 10},
		Right: &TreeNode{
			Val:   11,
			Right: &TreeNode{Val: 12},
		},
	},
}

func Test_ShouldDoA(t *testing.T) {
	assert.Exactly(
		t,
		10,
		FindClosestLeaf(exampleTree, 2),
	)
}

func Test_ShouldDoB(t *testing.T) {
	assert.Exactly(
		t,
		10,
		FindClosestLeaf(&TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val:   9,
					Right: &TreeNode{Val: 10},
				},
			},
		}, 4),
	)
}

func Test_ShouldDoC(t *testing.T) {
	assert.Exactly(
		t,
		4,
		FindClosestLeaf(&TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{Val: 3},
		}, 2),
	)
}

// func Test_ShouldDfsWork(t *testing.T) {
// 	target := dfs(exampleTree, 2)
// 	assert.Exactly(
// 		t,
// 		2,
// 		target.Val,
// 	)
// 	assert.Exactly(
// 		t,
// 		3,
// 		target.Left.Val,
// 	)
// }
