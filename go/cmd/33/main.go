package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		"search([]int{4,5,6,7,0,1,2}, 0) == 4",
		search([]int{4, 5, 6, 7, 0, 1, 2}, 0) == 4,
	)
	fmt.Println(
		"search([]int{4,5,6,7,0,1,2}, 3) == -1",
		search([]int{4, 5, 6, 7, 0, 1, 2}, 3) == -1,
	)
	fmt.Println(
		"search([]int{4,5,6,7,0,1}, 3) == -1",
		search([]int{4, 5, 6, 7, 0, 1}, 3) == -1,
	)
	fmt.Println(
		"search([]int{4,5,1}, 3) == -1",
		search([]int{4, 5, 1}, 3) == -1,
	)
	fmt.Println(
		"search([]int{1, 3}, 3) == 1",
		search([]int{1, 3}, 3) == 1,
	)
	fmt.Println(
		"search([]int{3, 1}, 3) == 0",
		search([]int{3, 1}, 3) == 0,
	)
	fmt.Println(
		"search([]int{1, 2, 3, 4, 5}, 5) == ?",
		search([]int{1, 2, 3, 4, 5}, 5),
	)
	fmt.Println(
		"search([]int{8, 9, 2, 3, 4}, 9) == 9",
		search([]int{8, 9, 2, 3, 4}, 9) == 1,
	)
}

func search(nums []int, target int) (ansIndex int) {
	offset := -1

	leftPartition := 0
	rightPartition := len(nums)

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	if !(nums[0] < nums[len(nums)-1]) {
		for {
			midPartition := leftPartition + (rightPartition-leftPartition)/2

			rightHead, rightTail := nums[midPartition], nums[rightPartition-1]
			pivotInRight := rightHead > rightTail

			if nums[midPartition-1] > nums[midPartition] {
				offset = midPartition - 1
				break
			}

			if pivotInRight {
				leftPartition = midPartition
			} else {
				rightPartition = midPartition
			}
		}

	}

	var newNums []int

	if offset == -1 {
		newNums = nums
	} else {
		newNums = append(nums[offset+1:], nums[:offset+1]...)
	}

	leftPartition = 0
	rightPartition = len(nums)
	offsetAns := -1

	for {
		shouldEnd := (rightPartition - leftPartition) <= 1

		midPartition := leftPartition + (rightPartition-leftPartition)/2
		midItem := newNums[midPartition]

		if midItem == target {
			offsetAns = midPartition
			break
		} else if target < midItem {
			rightPartition = midPartition
		} else if midItem < target {
			leftPartition = midPartition
		}

		if shouldEnd {
			break
		}
	}

	if offsetAns == -1 {
		return -1
	} else if offset == -1 {
		return offsetAns
	} else {
		return (offsetAns + offset + 1) % len(nums)
	}
}
