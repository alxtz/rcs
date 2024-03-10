package add

import (
	"rcs/cmd/770/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldPolyAddWork1(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{"10"},
		PolyAdd(utils.Polynomial{"1"}, utils.Polynomial{"9"}),
	)
}

func Test_ShouldPolyAddWork2(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{"-2"},
		PolyAdd(utils.Polynomial{"5"}, utils.Polynomial{"-7"}),
	)
}

func Test_ShouldPolyAddWork3(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{"-2*e"},
		PolyAdd(utils.Polynomial{"5*e"}, utils.Polynomial{"-7*e"}),
	)
}

func Test_ShouldPolyAddWork4(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{"-2*a", "-2*e"},
		utils.SortPoly(PolyAdd(utils.Polynomial{"-2*e"}, utils.Polynomial{"-2*a"})),
	)
}

func Test_ShouldPolyAddWork5(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{"2*e"},
		PolyAdd(utils.Polynomial{"-2*e"}, utils.Polynomial{"4*e"}),
	)
}

func Test_ShouldPolyAddWork6(t *testing.T) {
	assert.Exactly(
		t,
		utils.Polynomial{"1*e*a"},
		PolyAdd(utils.Polynomial{"2*e*a"}, utils.Polynomial{"-1*e*a"}),
	)
}

// func Test_ShouldPolyAddWork6(t *testing.T) {
// 	assert.Exactly(
// 		t,
// 		Polynomial{"2e*a", "-1*a*e"},
// 		PolyAdd(Polynomial{"-2e"}, Polynomial{"4e"}),
// 	)
// }

func Test_ShouldPolyAddWorkWhenDel(t *testing.T) {
	assert.Nil(
		t,
		PolyAdd(utils.Polynomial{"1*e*a"}, utils.Polynomial{"-1*e*a"}),
	)
}
