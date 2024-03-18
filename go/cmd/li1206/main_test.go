package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]int{-1, 3, -1},
		NextGreaterElement(
			[]int{4, 1, 2},
			[]int{1, 3, 4, 2},
		),
	)
}
