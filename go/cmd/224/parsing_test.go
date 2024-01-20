package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldA(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"1", "+", "1"},
		parsing("1 + 1"),
	)
}

func Test_ShouldB(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"2", "-", "1", "+", "2"},
		parsing(" 2-1 + 2 "),
	)
}

func Test_ShouldC(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"(", "1", "+", "(", "4", "+", "5", "+", "2", ")", "-", "3", ")", "-", "(", "6", "+", "8", ")"},
		parsing("(1+(4+5+2)-3)-(6+8)"),
	)
}

func Test_ShouldD(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"33"},
		parsing(" 33 "),
	)
}

func Test_ShouldE(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"1", "-", "(", "-", "2", ")"},
		parsing("1-( -2)"),
	)
}

func Test_ShouldF(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"-", "(", "3", "+", "(", "4", "+", "5", ")", ")"},
		parsing("- (3 + (4 + 5))"),
	)
}

func Test_ShouldReturnG(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"(", "(", ")", ")"},
		parsing("(())"),
	)
}
