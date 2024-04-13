package main

import (
	"fmt"
)

func main() {
	fmt.Println(
		"nextGreatestLetter([]byte{'c','f','j'}, 'a') == 'c'",
		nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a') == 'c',
	)
	fmt.Println(
		"nextGreatestLetter([]byte{'b', 'b', 'y', 'y', 'z', 'z'}, 'x') == 'b',",
		nextGreatestLetter([]byte{'b', 'b', 'y', 'y', 'z', 'z'}, 'x') == 'b',
	)
	fmt.Println(
		"nextGreatestLetter([]byte{'b', 'b', 'y',  'z', 'z'}, 'x') == 'b',",
		nextGreatestLetter([]byte{'b', 'b', 'y', 'z', 'z'}, 'x') == 'b',
	)
	fmt.Println(
		"nextGreatestLetter([]byte{'a', 'b'}, 'x') == 'b',",
		nextGreatestLetter([]byte{'a', 'b'}, 'x') == 'b',
	)
	fmt.Println(
		"nextGreatestLetter([]byte{'y', 'z'}, 'x') == 'y',",
		nextGreatestLetter([]byte{'y', 'z'}, 'x') == 'y',
	)
	fmt.Println(
		"nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c') == 'f',",
		nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c') == 'f',
	)
}

// given ordered char list
// there's at least 2 uq chars

func nextGreatestLetter(letters []byte, target byte) (ansByte byte) {
	leftPartition := 0
	rightPartition := len(letters)

	for {
		// should stop
		shouldStop := (rightPartition - leftPartition) <= 1

		midPartition := leftPartition + (rightPartition-leftPartition)/2
		midItem := letters[midPartition]

		if midItem == target {
			rightPartition = midPartition
		} else if target < midItem {
			rightPartition = midPartition
		} else if midItem < target {
			leftPartition = midPartition
		}

		if shouldStop {
			break
		}
	}

	fmt.Println("l,r", leftPartition, rightPartition)

	// if ansIndex >= 0 {
	// 	// return letters[ansIndex-1]
	// } else {
	return letters[leftPartition]
	// }
}
