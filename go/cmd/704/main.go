package main

import (
	"fmt"
)

func main() {
	fmt.Println("list=[4, 9, 11, 13, 18, 25], target=9, ans=??")
	fmt.Println("ans(index) =", BinarySearchTemplate1(
		[]int{4, 9, 11, 13, 18, 25}, 9,
	))
	fmt.Println("-----")
	fmt.Println("list=[4, 9, 11, 13, 18, 25, 31], target=9, ans=??")
	fmt.Println("ans(index) =", BinarySearchTemplate1(
		[]int{4, 9, 11, 13, 18, 25, 31}, 9,
	))
}

func BinarySearchTemplate1(nums []int, target int) (ansIndex int) {
	ansIndex = -1

	if len(nums) == 0 {
		return ansIndex
	}

	headIndex := 0
	tailIndex := len(nums) - 1

	for headIndex <= tailIndex {
		diff := tailIndex - headIndex
		midIndex := headIndex + diff/2 // selects a new mid, also note this floors
		mid := nums[midIndex]

		// fmt.Println("  -> exec", "head:", nums[headIndex], "tail:", nums[tailIndex], "mid:", mid)

		if mid == target {
			ansIndex = midIndex
			break
		}

		if target < mid {
			tailIndex = midIndex - 1
			// set current tail to mid
			// Next Search "head ~ mid"
		}

		if mid < target {
			headIndex = midIndex + 1
			// set head to be current mid
			// Next Search "mid ~ tail"
		}
	}

	return ansIndex
}

func Search(nums []int, target int) int {
	return BinarySearchTemplate1(nums, target)
}
