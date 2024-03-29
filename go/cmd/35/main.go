package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		"searchInsert([]int{1,3,5,6}, 5) == 2",
		searchInsert([]int{1, 3, 5, 6}, 5) == 2,
	)
	fmt.Println(
		"searchInsert([]int{1,3,5,6}, 2) == 1",
		searchInsert([]int{1, 3, 5, 6}, 2) == 1,
	)
	fmt.Println(
		"searchInsert([]int{1,3,5,6}, 7) == 4",
		searchInsert([]int{1, 3, 5, 6}, 7) == 4,
	)
	fmt.Println(
		"searchInsert([]int{1}, 7) == 1",
		searchInsert([]int{1}, 7) == 1,
	)
	fmt.Println(
		"searchInsert([]int{7}, 1) == 0",
		searchInsert([]int{7}, 1) == 0,
	)
}

func searchInsert(nums []int, target int) (ansIndex int) {
	leftPartition := 0
	rightPartition := len(nums)

	for {
		shouldEnd := (rightPartition - leftPartition) <= 1

		midPartition := leftPartition + (rightPartition-leftPartition)/2
		midItem := nums[midPartition]

		if midItem == target {
			return midPartition
		} else if target < midItem {
			rightPartition = midPartition
		} else if midItem < target {
			leftPartition = midPartition
		}

		if shouldEnd {
			break
		}
	}

	return rightPartition
}
