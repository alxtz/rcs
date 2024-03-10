package mult

import (
	"complicated-prob-solving/cmd/770/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldGive1Ans_WhenInput2Deg1Exps_0(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{"-20*a*b*c*d"},
		PolyMult(utils.Polynomial{"-4*a*b"}, utils.Polynomial{"5*c*d"}),
	)
}

func Test_ShouldGive1Ans_WhenInput2Deg1Exps_1(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{"-10"},
		PolyMult(utils.Polynomial{"5"}, utils.Polynomial{"-2"}),
	)
}

func Test_ShouldGive1Ans_WhenInput2Deg1Exps_2(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{"-8*a*b"},
		PolyMult(utils.Polynomial{"2"}, utils.Polynomial{"-4*a*b"}),
	)
}

func Test_ShouldWork3(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{"1*e*e", "-64"},
		PolyMult(utils.Polynomial{"1*e", "8"}, utils.Polynomial{"1*e", "-8"}),
	)
}

func Test_ShouldWork4(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{
			"2*e*e*e*e",
			"-24*e*e*e",
			"8*e*e",
			"384*e",
			"512",
		},
		PolyMult(
			PolyMult(
				PolyMult(
					[]string{"2", "1*e"},
					[]string{"1*e", "-8"},
				),
				PolyMult(
					[]string{"1*e", "2"},
					[]string{"-8", "1*e"},
				),
			),
			[]string{"2"},
		),
	)
}
