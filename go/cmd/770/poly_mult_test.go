package main

import (
	ds "complicated-prob-solving"
	"fmt"
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
	// t.Skip()
	assert.Exactly(
		t,
		Polynomial{"-10"},
		PolyMult(Polynomial{"5"}, Polynomial{"-2"}),
	)
}

func Test_ShouldGive1Ans_WhenInput2Deg1Exps_2(t *testing.T) {
	// t.Skip()
	assert.Exactly(
		t,
		Polynomial{"-8*a*b"},
		PolyMult(Polynomial{"2"}, Polynomial{"-4*a*b"}),
	)
}

func Test_ShouldDemonstrate(t *testing.T) {
	fmt.Println(ds.Readable(PolyMult(
		Polynomial{"-4*a*a", "2*a", "3"},
		Polynomial{"5*b*b", "2*c", "1"},
	)))
}

func Test_ShouldWork3(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"1*e*e"},
		PolyMult(Polynomial{"1*e", "8"}, Polynomial{"1*e", "-8"}),
	)
}
