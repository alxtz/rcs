package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldCalc2(t *testing.T) {
	assert.Exactly(
		t,
		2,
		calculate("1 + 1"),
	)
}

func Test_ShouldCalc3(t *testing.T) {
	assert.Exactly(
		t,
		3,
		calculate(" 2-1 + 2 "),
	)
}

func Test_ShouldCalc23(t *testing.T) {
	assert.Exactly(
		t,
		23,
		calculate("(1+(4+5+2)-3)+(6+8)"),
	)
}

func Test_ShouldCalcN5(t *testing.T) {
	assert.Exactly(
		t,
		-5,
		calculate("(1+(4+5+2)-3)-(6+8)"),
	)
}

func Test_ShouldCalc33(t *testing.T) {
	assert.Exactly(
		t,
		33,
		calculate(" 33 "),
	)
}

func Test_ShouldCalcP3(t *testing.T) {
	assert.Exactly(
		t,
		3,
		calculate("1-( -2)"),
	)
}

func Test_ShouldCalcM12(t *testing.T) {
	assert.Exactly(
		t,
		-12,
		calculate("- (3 + (4 + 5))"),
	)
}

func Test_ShouldCalc0(t *testing.T) {
	assert.Exactly(
		t,
		0,
		calculate("(())"),
	)
}
