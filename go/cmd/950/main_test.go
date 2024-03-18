package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]int{2, 13, 3, 11, 5, 17, 7},
		deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}),
	)
}
