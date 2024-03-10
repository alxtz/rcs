package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateStackWithEmptyVec(t *testing.T) {
	stack := Constructor()
	assert.Exactly(
		t,
		0,
		len(stack.Vector),
	)
}

func Test_ShouldHaveLen1AndVal_WhenPush(t *testing.T) {
	stack := Constructor()
	stack.Push(10)
	assert.Exactly(
		t,
		1,
		len(stack.Vector),
	)
	assert.Exactly(
		t,
		10,
		stack.Vector[0],
	)
}

func Test_ShouldHaveLen1AndVal_WhenPop(t *testing.T) {
	stack := Constructor()
	stack.Push(11)
	stack.Push(22)
	popped := stack.Pop()

	assert.Exactly(t, popped, 22)
	assert.Exactly(t, 1, len(stack.Vector))
	assert.Exactly(t, 11, stack.Vector[0])
}

func Test_ShouldHaveLen1AndVal_WhenTop(t *testing.T) {
	stack := Constructor()
	stack.Push(10)
	assert.Exactly(t, 10, stack.Top())
}

func Test_ShouldEmptyWork(t *testing.T) {
	stack := Constructor()
	expectDefaultEmpty := stack.Empty()
	stack.Push(11)
	expectNonEmpty := stack.Empty()
	stack.Pop()
	expectEmptyAfterPop := stack.Empty()

	assert.True(t, expectDefaultEmpty)
	assert.False(t, expectNonEmpty)
	assert.True(t, expectEmptyAfterPop)
}
