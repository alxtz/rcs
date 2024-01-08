package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldBe2456(t *testing.T) {
	assert.ElementsMatch(
		t,
		[]int{2, 4, 5, 6},
		eventualSafeNodes(
			[][]int{
				{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {},
			},
		),
	)
}

func Test_ShouldBe4(t *testing.T) {
	assert.ElementsMatch(
		t,
		[]int{4},
		eventualSafeNodes(
			[][]int{
				{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {},
			},
		),
	)
}

func Test_ShouldBe01234(t *testing.T) {
	assert.ElementsMatch(
		t,
		[]int{0, 1, 2, 3, 4},
		eventualSafeNodes(
			[][]int{
				{}, {0, 2, 3, 4}, {3}, {4}, {},
			},
		),
	)
}
