package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldExtract_When1(t *testing.T) {
	coe, indeterminate := ExtractFromTerm("-8*a*b")

	assert.Exactly(
		t,
		-8,
		coe,
	)

	assert.Exactly(
		t,
		"a*b",
		*indeterminate,
	)
}

func Test_ShouldExtract_When2(t *testing.T) {
	coe, indeterminate := ExtractFromTerm("-10")

	assert.Exactly(
		t,
		-10,
		coe,
	)

	assert.Nil(
		t,
		indeterminate,
	)
}

func Test_ShouldSort_When1(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"-20*a*a*b*b", "-8*a*a*c", "10*a*b*b", "-4*a*a", "15*b*b", "4*a*c", "2*a", "6*c", "3"},
		SortPoly([]string{"6*c", "-20*a*a*b*b", "-8*a*a*c", "-4*a*a", "10*a*b*b", "4*a*c", "2*a", "15*b*b", "3"}),
	)
}
