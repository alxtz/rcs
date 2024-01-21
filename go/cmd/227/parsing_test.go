package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldParseA(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"133", "+", "52", "/", "2"},
		parsing(" 133+52 / 2 "),
	)
}

func Test_ShouldParseB(t *testing.T) {
	assert.Exactly(
		t,
		[]string{"3", "*", "2", "+", "23", "*", "3"},
		parsing("3* 2+23*3"),
	)
}
