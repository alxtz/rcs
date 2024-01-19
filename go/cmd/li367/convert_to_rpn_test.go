package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldPassRpn0(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{"3", "4", "-", "5", "+"},
		ConvertToRPN([]string{"3", "-", "4", "+", "5"}),
	)
}

func Test_ShouldPassRpn1(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{"A", "B", "C", "*", "+"},
		ConvertToRPN([]string{"A", "+", "B", "*", "C"}),
	)
}

func Test_ShouldPassRpn2(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{"5", "6", "-", "7", "*"},
		ConvertToRPN([]string{"(", "5", "-", "6", ")", "*", "7"}),
	)
}

func Test_ShouldPassRpn3(t *testing.T) {
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

func Test_ShouldPassRpn4(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		[]string{},
		ConvertToRPN([]string{
			"(", "(", "(", ")", ")", ")",
		}),
	)
}

func Test_ShouldPassRpn5(t *testing.T) {
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

func Test_ShouldPassRpn6(t *testing.T) {
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
