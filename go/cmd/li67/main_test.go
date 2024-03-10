package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldA(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]int{4, 10, 12, 15, 18, 22, 24},
		InorderTraversal(&TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val:   10,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 12},
			},
			Right: &TreeNode{
				Val:   22,
				Left:  &TreeNode{Val: 18},
				Right: &TreeNode{Val: 24},
			},
		}),
	)
}
