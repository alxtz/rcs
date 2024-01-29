package mult

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldGive1Ans_WhenInput2Deg1Exps_0(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"-20*a*b*c*d"},
		PolyMult(Polynomial{"-4*a*b"}, Polynomial{"5*c*d"}),
	)
}

func Test_ShouldGive1Ans_WhenInput2Deg1Exps_1(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"-10"},
		PolyMult(Polynomial{"5"}, Polynomial{"-2"}),
	)
}

func Test_ShouldGive1Ans_WhenInput2Deg1Exps_2(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"-8*a*b"},
		PolyMult(Polynomial{"2"}, Polynomial{"-4*a*b"}),
	)
}

func Test_ShouldWork3(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"1*e*e", "-64"},
		PolyMult(Polynomial{"1*e", "8"}, Polynomial{"1*e", "-8"}),
	)
}
