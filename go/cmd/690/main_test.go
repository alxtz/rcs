package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldDoA_WhenGivenB(t *testing.T) {
	assert.Exactly(
		t,
		11,
		getImportance([]*Employee{
			{Id: 1, Importance: 5, Subordinates: []int{2, 3}},
			{Id: 2, Importance: 3, Subordinates: []int{}},
			{Id: 3, Importance: 3, Subordinates: []int{}},
		}, 1),
	)
}
