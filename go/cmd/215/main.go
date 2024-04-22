package main

import (
	"fmt"
)

func main() {
	fmt.Println("QuickSelect_3way([]int{40, 41, 17, -22, 25, 55, -18, 35, 10, 25, 33, 19, 44, 51, 24}, 4)",
		QuickSelect_3way([]int{40, 41, 17, -22, 25, 55, -18, 35, 10, 25, 33, 19, 44, 51, 24}, 4),
	)
	// fmt.Println("QuickSelect_3way([]int{-1, -1}, 2)",
	// 	QuickSelect_3way([]int{-1, -1}, 2),
	// )
	fmt.Println("QuickSelect_3way([]int{-1, -1}, 2)",
		QuickSelect_3way([]int{-1, -1}, 2),
	)
}

func QuickSelect_3way(input []int, kth int) int {
	var recurFunc func(recurInput []int, recurKth int) []int

	recurFunc = func(recurInput []int, recurKth int) []int {
		pivotValue := recurInput[len(recurInput)-1]
		listToPartition := recurInput

		lowPartition := 0
		midPartition := 0

		highPartition := len(listToPartition)

		for midPartition != highPartition {

			selectedUnknown := listToPartition[midPartition]

			if selectedUnknown < pivotValue {
				listToPartition[lowPartition], listToPartition[midPartition] = listToPartition[midPartition], listToPartition[lowPartition]

				lowPartition += 1
				midPartition += 1
			} else if pivotValue < selectedUnknown {
				listToPartition[midPartition], listToPartition[highPartition-1] = listToPartition[highPartition-1], listToPartition[midPartition]
				highPartition -= 1
			} else if pivotValue == selectedUnknown {
				midPartition += 1
			}
		}

		toPassLeft, toPassRight := recurInput[:lowPartition], recurInput[highPartition:]

		higherThanMidCount := len(listToPartition) - highPartition
		higherThanLowCount := len(listToPartition) - lowPartition

		if recurKth <= higherThanMidCount {
			return recurFunc(toPassRight, recurKth)
		} else if recurKth > higherThanLowCount {
			nth := recurKth - (len(listToPartition) - lowPartition)
			return recurFunc(toPassLeft, nth)
		} else {
			return []int{pivotValue}
		}
	}

	ans := recurFunc(input, kth)

	return ans[0]
}
