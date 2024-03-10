package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldA(t *testing.T) {
	assert.Exactly(
		t,
		[]int{5, 10},
		AsteroidCollision([]int{5, 10, -5}),
	)
}

func Test_ShouldB(t *testing.T) {
	assert.Exactly(
		t,
		[]int{10},
		AsteroidCollision([]int{10, 2, -5}),
	)
}

func Test_ShouldC(t *testing.T) {
	assert.Exactly(
		t,
		[]int{1, 2, 3},
		AsteroidCollision([]int{1, 2, 3}),
	)
}

func Test_ShouldD(t *testing.T) {
	assert.Exactly(
		t,
		[]int{-1, -2, -3},
		AsteroidCollision([]int{-1, -2, -3}),
	)
}

func Test_ShouldE(t *testing.T) {
	assert.Exactly(
		t,
		[]int{-1, -3, 4},
		AsteroidCollision([]int{-1, 2, -3, 4}),
	)
}

func Test_ShouldF(t *testing.T) {
	assert.Exactly(
		t,
		[]int{-10},
		AsteroidCollision([]int{1, 2, 3, -10}),
	)
}

func Test_ShouldG(t *testing.T) {
	assert.Exactly(
		t,
		[]int{},
		AsteroidCollision([]int{10, 1, 2, 3, -10}),
	)
}
