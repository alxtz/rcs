package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldBe2(t *testing.T) {
	t.Skip()
	assert.Exactly(
		t,
		2,
		EvaluateExpression([]string{"2", "*", "6", "-", "(", "23", "+", "7", ")", "/", "(", "1", "+", "2", ")"}),
	)
}

func Test_ShouldBe7(t *testing.T) {
	t.Skip()
	assert.Exactly(
		t,
		7,
		EvaluateExpression([]string{"4", "-", "(", "2", "-", "3", ")", "*", "2", "+", "5", "/", "5"}),
	)
}

func Test_ShouldBuildNodeTree1(t *testing.T) {
	t.Skip()
	fmt.Println("result", Build([]string{"2", "-", "3"}))
}

func Test_ShouldBuildNodeTree2(t *testing.T) {
	fmt.Println("result", Build([]string{"2", "+", "5", "/", "5", "-", "10"}))
}
