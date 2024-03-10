package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenA(t *testing.T) {
	assert.True(t, isValid("()"))

	assert.True(t, isValid("()[]{}"))

	assert.False(t, isValid("(]"))

	assert.False(t, isValid("}"))
	assert.False(t, isValid("{"))

	assert.False(t, isValid("([)]"))
}
