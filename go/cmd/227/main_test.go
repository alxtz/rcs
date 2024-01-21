package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldCalc7(t *testing.T) {
	assert.Exactly(
		t,
		7,
		calculate("3+2*2"),
	)
}

func Test_ShouldCalc1(t *testing.T) {
	assert.Exactly(
		t,
		1,
		calculate(" 3/2 "),
	)
}

func Test_ShouldCalc5(t *testing.T) {
	assert.Exactly(
		t,
		5,
		calculate(" 3+5 / 2 "),
	)
}

func Test_ShouldCalc66(t *testing.T) {
	assert.Exactly(
		t,
		66,
		calculate("30* 2+2*3"),
	)
}

func Test_ShouldCalc34(t *testing.T) {
	assert.Exactly(
		t,
		34,
		calculate("30* 2+2*3 - 4*5 - 3*4"),
	)
}

func Test_ShouldDiv(t *testing.T) {
	assert.Exactly(
		t,
		1,
		div(3, 2),
	)
}
