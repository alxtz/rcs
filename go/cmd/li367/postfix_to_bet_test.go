package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	ds "complicated-prob-solving"
)

func Test_ShouldPassBet0(t *testing.T) {
	// t.Skip()
	// assert.Exactly(
	// 	t,
	// 	[]string{"3", "4", "-", "5", "+"},
	// 	ConvertToRPN([]string{"3", "-", "4", "+", "5"}),
	// )

	// fmt.Println(ds.Readable(PostfixToBet(nil)))
	fmt.Println(ds.Readable(PostfixToBet(
		[]string{"3", "4", "-", "5", "+"},
	)))

	fmt.Println(ds.Readable(PostfixToBet(
		[]string{"A", "B", "+", "C", "D", "-", "*"},
	)))

	fmt.Println(ds.Readable(Build(
		[]string{"2", "*", "6", "-", "(", "23", "+", "7", ")", "/", "(", "1", "+", "2", ")"},
	)))

	fmt.Println(ds.Readable(Build(
		[]string{"(", "(", "(", ")", ")", ")"},
	)))

	fmt.Println(ds.Readable(Build(
		nil,
	)))
}

func Test_ShouldPassBet1(t *testing.T) {
	// t.Skip()
	// assert.Exactly(
	// 	t,
	// 	[]string{"A", "B", "C", "*", "+"},
	// 	ConvertToRPN([]string{"A", "+", "B", "*", "C"}),
	// )
}

func Test_ShouldPassBet2(t *testing.T) {
	// t.Skip()
	// assert.Exactly(
	// 	t,
	// 	[]string{"5", "6", "-", "7", "*"},
	// 	ConvertToRPN([]string{"(", "5", "-", "6", ")", "*", "7"}),
	// )
}

func Test_ShouldPassBet3(t *testing.T) {
	// t.Skip()
	// assert.Exactly(
	// 	t,
	// 	[]string{"A", "B", "+", "C", "D", "-", "*"},
	// 	ConvertToRPN([]string{
	// 		"(", "A", "+", "B", ")",
	// 		"*",
	// 		"(", "C", "-", "D", ")",
	// 	}),
	// )
}

func Test_ShouldPassBet4(t *testing.T) {
	// t.Skip()
	// assert.Exactly(
	// 	t,
	// 	[]string{},
	// 	ConvertToRPN([]string{
	// 		"(", "(", "(", ")", ")", ")",
	// 	}),
	// )
}

func Test_ShouldPassBet5(t *testing.T) {
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

func Test_ShouldPassBet6(t *testing.T) {
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
