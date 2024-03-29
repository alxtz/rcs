package main

import "fmt"

func main() {
	ans1 := FixedPoint([]int{0, 1, 3, 5, 7})
	fmt.Println("FixedPoint([]int{0, 1, 3, 5, 7})", ans1 == 0)
}

func FixedPoint(nums []int) (ansIndex int) {
	ansIndex = -1

	// write your code here
	leftPartition := 0
	rightPartition := len(nums)

	for {
		shouldEnd := rightPartition-leftPartition <= 1

		midPartition := leftPartition + (rightPartition-leftPartition)/2
		midItem := nums[midPartition]

		if midPartition < midItem {
			// things is in left half
			rightPartition = midPartition
		} else if midItem < midPartition {
			// things is in right half
			leftPartition = midPartition
		} else {
			ansIndex = midPartition
			rightPartition = midPartition
			// break
		}

		if shouldEnd {
			break
		}
	}

	return ansIndex
}
