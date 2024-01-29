package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldPolyAddWork1(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"10"},
		PolyAdd(Polynomial{"1"}, Polynomial{"9"}),
	)
}

func Test_ShouldPolyAddWork2(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"-2"},
		PolyAdd(Polynomial{"5"}, Polynomial{"-7"}),
	)
}

func Test_ShouldPolyAddWork3(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"-2*e"},
		PolyAdd(Polynomial{"5*e"}, Polynomial{"-7*e"}),
	)
}

func Test_ShouldPolyAddWork4(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"-2*e", "-2*a"},
		PolyAdd(Polynomial{"-2*e"}, Polynomial{"-2*a"}),
	)
}

func Test_ShouldPolyAddWork5(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"2*e"},
		PolyAdd(Polynomial{"-2*e"}, Polynomial{"4*e"}),
	)
}

func Test_ShouldPolyAddWork6(t *testing.T) {
	assert.Exactly(
		t,
		Polynomial{"1*e*a"},
		PolyAdd(Polynomial{"2*e*a"}, Polynomial{"-1*e*a"}),
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
		PolyAdd(Polynomial{"1*e*a"}, Polynomial{"-1*e*a"}),
	)
}
