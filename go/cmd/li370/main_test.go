package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldPass0(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{"3", "4", "-", "5", "+"},
		ConvertToRPN([]string{"3", "-", "4", "+", "5"}),
	)
}

func Test_ShouldPass1(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{"A", "B", "C", "*", "+"},
		ConvertToRPN([]string{"A", "+", "B", "*", "C"}),
	)
}

func Test_ShouldPass2(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{"5", "6", "-", "7", "*"},
		ConvertToRPN([]string{"(", "5", "-", "6", ")", "*", "7"}),
	)
}

func Test_ShouldPass3(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{"A", "B", "+", "C", "D", "-", "*"},
		ConvertToRPN([]string{
			"(", "A", "+", "B", ")",
			"*",
			"(", "C", "-", "D", ")",
		}),
	)
}

func Test_ShouldPass4(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{},
		ConvertToRPN([]string{
			"(", "(", "(", ")", ")", ")",
		}),
	)
}

func Test_ShouldPass5(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{
			"2", "6", "*", "23", "7", "+", "1", "2", "+", "/", "-",
		},
		ConvertToRPN([]string{
			"2", "*", "6", "-", "(", "23", "+", "7", ")", "/", "(", "1", "+", "2", ")",
		}),
	)
}

func Test_ShouldPass6(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{
			"2", "6", "*", "23", "7", "+", "1", "2", "+", "/", "-",
		},
		ConvertToRPN([]string{
			"2", "*", "6", "-", "(", "23", "+", "7", ")", "/", "(", "1", "+", "2", ")",
		}),
	)
}
