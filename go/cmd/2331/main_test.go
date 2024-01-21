package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA(t *testing.T) {
	assert.Exactly(
		t,
		true,
		evaluateTree(&TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 0},
				Right: &TreeNode{Val: 1},
			},
		}),
	)
}
