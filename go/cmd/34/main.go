package main

import (
	"fmt"
)

func main() {
	ans1 := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println("searchRange([]int{5,7,7,8,8,10}, 8) == [3, 4]", ans1[0] == 3 && ans1[1] == 4)

	ans2 := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	fmt.Println("searchRange([]int{5,7,7,8,8,10}, 6) == [-1, -1]", ans2[0] == -1 && ans2[1] == -1)
}

func searchRange(nums []int, target int) (ans []int) {
	ans = []int{-1, -1}

	if len(nums) == 0 {
		return ans
	}

	rangeStartIndex := -1
	rangeEndIndex := -1

	// 2nd binary search, find rangeStartIndex
	leftPartition := 0
	rightPartition := len(nums)

	for {
		shouldEnd := (rightPartition - leftPartition) <= 1

		midPartition := leftPartition + (rightPartition-leftPartition)/2
		midItem := nums[midPartition]

		if target == midItem {
			rangeStartIndex = midPartition
			rightPartition = midPartition
		} else if target < midItem {
			rightPartition = midPartition
		} else if midItem < target {
			leftPartition = midPartition
		}

		if shouldEnd {
			break
		}
	}

	leftPartition = 0
	rightPartition = len(nums)

	for {
		shouldEnd := (rightPartition - leftPartition) <= 1

		midPartition := leftPartition + (rightPartition-leftPartition)/2
		midItem := nums[midPartition]

		if target == midItem {
			rangeEndIndex = midPartition
			leftPartition = midPartition
		} else if target < midItem {
			rightPartition = midPartition
		} else if midItem < target {
			leftPartition = midPartition
		}

		if shouldEnd {
			break
		}
	}

	ans[0] = rangeStartIndex
	ans[1] = rangeEndIndex

	return ans
}
