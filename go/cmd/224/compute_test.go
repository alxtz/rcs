package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldComputeA(t *testing.T) {
	assert.Exactly(
		t,
		2,
		symbolsToInt([]string{"1", "+", "1"}),
	)
}

func Test_ShouldComputeB(t *testing.T) {
	assert.Exactly(
		t,
		3,
		symbolsToInt([]string{"2", "-", "1", "+", "2"}),
	)
}

func Test_ShouldComputeC(t *testing.T) {
	assert.Exactly(
		t,
		-5,
		symbolsToInt([]string{"1", "+", "11", "-", "3", "-", "14"}),
	)
}

func Test_ShouldComputeE(t *testing.T) {
	assert.Exactly(
		t,
		-3,
		symbolsToInt([]string{"-", "1", "-", "2"}),
	)
}

func Test_ShouldComputeF(t *testing.T) {
	assert.Exactly(
		t,
		6,
		symbolsToInt([]string{"-", "3", "+", "9"}),
	)
}

func Test_ShouldComputeG(t *testing.T) {
	assert.Exactly(
		t,
		9,
		symbolsToInt([]string{"1", "+", "11", "-", "3"}),
	)
}
