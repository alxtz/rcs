package main

import (
	"fmt"
)

func main() {
	fmt.Println("list=[4, 9, 11, 13, 18, 25], target=9, ans=??")
	fmt.Println("ans(index) =", BinarySearchTemplate2(
		[]int{4, 9, 11, 13, 18, 25}, 9,
	))
	fmt.Println("-----")
	fmt.Println("list=[4, 9, 11, 13, 18, 25, 31], target=9, ans=??")
	fmt.Println("ans(index) =", BinarySearchTemplate2(
		[]int{4, 9, 11, 13, 18, 25, 31}, 9,
	))
}

func BinarySearchTemplate2(nums []int, target int) (ansIndex int) {
	ansIndex = -1

	if len(nums) == 0 {
		return ansIndex
	}

	headIndex, tailIndex := 0, len(nums)-1

	for headIndex <= tailIndex {
		diff := tailIndex - headIndex
		midIndex := headIndex + diff/2
		mid := nums[midIndex]

		fmt.Println(
			"  -> exec",
			"head:", nums[headIndex],
			"tail:", nums[tailIndex],
			"mid:", mid,
			"search against:", nums[headIndex:tailIndex+1],
		)

		switch true {
		case mid == target:
			ansIndex = midIndex
			return ansIndex
		case target < mid:
			tailIndex = midIndex
		case mid < target:
			headIndex = midIndex + 1
		}
	}

	// if nums[headIndex] == target {
	// 	ansIndex = headIndex
	// }

	return ansIndex
}
