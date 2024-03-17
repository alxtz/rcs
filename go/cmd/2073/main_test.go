package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA(t *testing.T) {
	assert.Exactly(
		t,
		6,
		timeRequiredToBuy([]int{2, 3, 2}, 2),
	)
}

func Test_ShouldDoB(t *testing.T) {
	assert.Exactly(
		t,
		8,
		timeRequiredToBuy([]int{5, 1, 1, 1}, 0),
	)
}
