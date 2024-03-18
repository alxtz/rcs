package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA(t *testing.T) {
	assert.Exactly(
		t,
		3,
		GetBubbleTea([][]byte{
			{'X', 'X', 'X', 'X', 'X', 'X'},
			{'X', '*', 'O', 'O', 'O', 'X'},
			{'X', 'O', 'O', '#', 'O', 'X'},
			{'X', 'X', 'X', 'X', 'X', 'X'},
		}),
	)
}

func Test_ShouldDoB(t *testing.T) {
	assert.Exactly(
		t,
		-1,
		GetBubbleTea([][]byte{
			{'X', 'X', 'X', 'X', 'X'},
			{'X', '*', 'X', 'O', 'X'},
			{'X', 'O', 'X', '#', 'X'},
			{'X', 'X', 'X', 'X', 'X'},
		}),
	)
}
