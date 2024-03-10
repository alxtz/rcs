package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA(t *testing.T) {
	assert.Exactly(
		t,
		1,
		ParenthesesScore("()"),
	)
}

func Test_ShouldDoB(t *testing.T) {
	assert.Exactly(
		t,
		2,
		ParenthesesScore("(())"),
	)
}

func Test_ShouldDoC(t *testing.T) {
	assert.Exactly(
		t,
		2,
		ParenthesesScore("()()"),
	)
}

func Test_ShouldDoD(t *testing.T) {
	assert.Exactly(
		t,
		6,
		ParenthesesScore("(()(()))"),
	)
}

func Test_ShouldDoE(t *testing.T) {
	assert.Exactly(
		t,
		3,
		ParenthesesScore("()()()"),
	)
}
