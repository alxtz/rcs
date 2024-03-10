package main

import (
	"rcs/cmd/770/mult"
	"rcs/cmd/770/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDemonstrate(t *testing.T) {
	result := utils.SortPoly([]string(mult.PolyMult(
		[]string{"-4*a*a", "2*a", "3"},
		[]string{"5*b*b", "2*c", "1"},
	)))
	assert.Exactly(
		t,
		[]string{
			"-20*a*a*b*b",
			"-8*a*a*c",
			"10*a*b*b",
			"-4*a*a",
			"15*b*b",
			"4*a*c",
			"2*a",
			"6*c",
			"3",
		},
		result,
	)
}

func Test_ShouldCalcA(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"-1*a", "1*e", "13"},
		BasicCalculatorIV(
			"e + 8 - a + 5",
			[]string{"a", "e"},
			[]int{},
		),
	)

	assert.Exactly(
		t,
		[]string{"6*e*e"},
		BasicCalculatorIV(
			"e * e * 2 * 3",
			[]string{"e"},
			[]int{},
		),
	)

	assert.Exactly(
		t,
		[]string{"-1*e"},
		BasicCalculatorIV(
			"e * 2 - e * 3",
			[]string{"e"},
			[]int{},
		),
	)
}

func Test_ShouldCalcB(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"1*e*e", "-64"},
		BasicCalculatorIV(
			"(e + 8) * (e - 8)",
			[]string{"e"},
			[]int{},
		),
	)

	assert.Exactly(
		t,
		[]string{
			"2*e*e*e*e",
			"-24*e*e*e",
			"8*e*e",
			"384*e",
			"512",
		},
		BasicCalculatorIV(
			"2 * ((e + 2) * (e - 8)) * ((e + 2) * (e - 8))",
			[]string{"e"},
			[]int{},
		),
	)
}

// func Test_ShouldB(t *testing.T) {
// 	assert.Exactly(
// 		t,
// 		[]string{"-1*a", "14"},
// 		basicCalculatorIV(
// 			"e - 8 + temperature - pressure",
// 			[]string{"e", "temperature"},
// 			[]int{1, 12},
// 		),
// 	)
// }

// func Test_ShouldD(t *testing.T) {
// 	assert.Exactly(
// 		t,
// 		[]string{},
// 		basicCalculatorIV(
// 			"e * ((e + 2) * (e - 8)) * ((e + 2) * (e - 8))",
// 			[]string{"e", "temperature"},
// 			[]int{1, 12},
// 		),
// 	)
// }

// func Test_ShouldD(t *testing.T) {
// 	assert.Exactly(
// 		t,
// 		[]string{},
// 		basicCalculatorIV(
// 			"(e * 8) * (e - 8)",
// 			[]string{"e", "temperature"},
// 			[]int{1, 12},
// 		),
// 	)
// }
