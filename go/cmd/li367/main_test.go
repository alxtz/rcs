package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	assert.Exactly(
		t,
		[]int{1, 2, 3},
		[]int{1, 2, 3},
	)
}
