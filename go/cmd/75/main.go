package main

import (
	"fmt"
)

func main() {

	input1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(input1)
	fmt.Println("sortColors([]int{2, 0, 2, 1, 1, 0}):", input1)
}

const lowVal = 0
const midVal = 1
const highVal = 2

func sortColors(nums []int) {
	lowPartition := 0
	midPartition := 0
	highPartition := len(nums)

	for {
		if midPartition == highPartition {
			// end of algo, return (modified input)
			return
		}

		// reveals item right of mid partition
		revealed := nums[midPartition]

		switch revealed {
		case lowVal:
			nums[lowPartition], nums[midPartition] = nums[midPartition], nums[lowPartition]
			midPartition += 1
			lowPartition += 1
		case midVal:
			midPartition += 1
		case highVal:
			nums[midPartition], nums[highPartition-1] = nums[highPartition-1], nums[midPartition]
			highPartition -= 1
		}
	}
}
