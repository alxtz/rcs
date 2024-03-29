package main

import "fmt"

func main() {
	ans1 := BinarySearchPartitioning([]int{-1, 0, 3, 5, 9, 12}, 9)
	fmt.Println("searching [-1,0,3,5,9,12], 9", ans1 == 4)

	ans2 := BinarySearchPartitioning([]int{-1, 0, 3, 5, 9, 12}, 2)
	fmt.Println("searching [-1,0,3,5,9,12], 9", ans2 == -1)

	ans3 := BinarySearchPartitioning([]int{9}, 9)
	fmt.Println("searching [9], 9", ans3 == 0)

	ans4 := BinarySearchPartitioning([]int{0, 9}, 9)
	fmt.Println("searching [9], 9", ans4 == 1)

	ans5 := BinarySearchPartitioning([]int{0, 9}, 0)
	fmt.Println("searching [9], 9", ans5 == 0)

	ans6 := BinarySearchPartitioning([]int{2, 9}, 1)
	fmt.Println("searching [9], 9", ans6 == -1)
}

func BinarySearchPartitioning(nums []int, target int) (ansIndex int) {
	ansIndex = -1

	leftPartition := 0
	rightPartition := len(nums)

	for {
		shouldEnd := (rightPartition - leftPartition) <= 1

		//
		midPartition := leftPartition + (rightPartition-leftPartition)/2
		midItem := nums[midPartition]

		if target < midItem {
			rightPartition = midPartition
		} else if midItem < target {
			leftPartition = midPartition
		} else {
			ansIndex = midPartition
			break
		}
		//

		if shouldEnd {
			break
		}
	}

	return ansIndex
}

func Search(nums []int, target int) int {
	return BinarySearchPartitioning(nums, target)
}
