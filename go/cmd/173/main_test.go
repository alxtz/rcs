package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	// t.Skip()

	var tree = TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val:   15,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20},
		},
	}

	var itrr = Constructor(&tree)

	assert.Exactly(
		t,
		3,
		itrr.Next(),
	)

	assert.Exactly(
		t,
		7,
		itrr.Next(),
	)

	assert.Exactly(
		t,
		true,
		itrr.HasNext(),
	)

	assert.Exactly(
		t,
		9,
		itrr.Next(),
	)

	assert.Exactly(
		t,
		15,
		itrr.Next(),
	)

	assert.Exactly(
		t,
		20,
		itrr.Next(),
	)

	assert.Exactly(
		t,
		false,
		itrr.HasNext(),
	)
}
