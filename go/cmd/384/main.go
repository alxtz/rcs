package main

import (
	"fmt"
	"math/rand"
)

func main() {
	solution := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(
		"solution.Shuffle()",
		solution.Shuffle(),
	)
	fmt.Println(
		"solution.Reset()",
		solution.Reset(),
	)
	fmt.Println(
		"solution.Shuffle()",
		solution.Shuffle(),
	)
}

type Solution struct {
	originalNums []int
	shuffledNums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		append([]int{}, nums...),
		append([]int{}, nums...),
	}
}

func (sol *Solution) Reset() []int {
	sol.shuffledNums = append([]int{}, sol.originalNums...)

	return sol.shuffledNums
}

func (sol *Solution) Shuffle() []int {
	// sol.shuffledNums = append([]int{}, sol.originalNums...)
	anchor := len(sol.shuffledNums) - 1

	for anchor >= 0 {
		// index to swap with, within range of [0 ~ anchor]
		toSwapWith := int(rand.Int31() % int32(anchor+1))

		if toSwapWith != anchor {
			sol.shuffledNums[toSwapWith], sol.shuffledNums[anchor] =
				sol.shuffledNums[anchor], sol.shuffledNums[toSwapWith]
		}

		anchor -= 1
	}

	return sol.shuffledNums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
