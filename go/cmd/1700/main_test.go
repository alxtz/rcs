package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA(t *testing.T) {
	assert.Exactly(
		t,
		0,
		countStudents(
			[]int{1, 1, 0, 0},
			[]int{0, 1, 0, 1},
		),
	)
}

func Test_ShouldDoB(t *testing.T) {
	assert.Exactly(
		t,
		3,
		countStudents(
			[]int{1, 1, 1, 0, 0, 1},
			[]int{1, 0, 0, 0, 1, 1},
		),
	)
}
