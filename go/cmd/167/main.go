package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")

	fmt.Println(
		twoSum([]int{2, 7, 11, 15}, 9),
	)
	fmt.Println(
		twoSum([]int{2, 3, 4}, 6),
	)
	fmt.Println(
		twoSum([]int{-1, 0}, -1),
	)
}

// 1D ordered int array

func twoSum(numbers []int, target int) []int {
	for i, val := range numbers {
		toFind := target - val

		leftPos := i
		rightPos := len(numbers)

		for {
			shouldEnd := (rightPos - leftPos) <= 1
			midPost := leftPos + (rightPos-leftPos)/2
			midItem := numbers[midPost]

			if midItem == toFind {
				return []int{i + 1, midPost + 1}
			} else if toFind < midItem {
				rightPos = midPost
			} else if midItem < toFind {
				leftPos = midPost
			}

			if shouldEnd {
				break
			}
		}
	}

	return []int{}
}
