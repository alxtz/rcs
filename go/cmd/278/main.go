package main

import "fmt"

var globalBad = -1

func main() {
	// v = 2
	globalBad = 2
	fmt.Println("bad (l=2)", firstBadVersion(2))
	fmt.Println("bad (l=4)", firstBadVersion(4))
	fmt.Println("bad (l=5)", firstBadVersion(5))

	// v = 1
	globalBad = 1
	fmt.Println("bad (l=1)", firstBadVersion(1))
	fmt.Println("bad (l=2)", firstBadVersion(2))
	fmt.Println("bad (l=4)", firstBadVersion(4))
	fmt.Println("bad (l=5)", firstBadVersion(5))

	// v = 4
	globalBad = 4
	fmt.Println("bad (l=5)", firstBadVersion(5))
}

func isBadVersion(version int) bool {
	return version >= globalBad
	// return version >= 1
	// return version >= 2
}

func firstBadVersion(n int) (minBadVersion int) {
	minBadVersion = -1

	leftPartition := 0
	rightPartition := n

	for {
		shouldEndInBreak := (rightPartition - leftPartition) <= 1

		func() {
			midPartition := leftPartition + (rightPartition-leftPartition)/2
			midItemVersion := midPartition + 1

			isMidVersionBad := isBadVersion(midItemVersion)

			if isMidVersionBad {
				minBadVersion = midItemVersion
			}

			// l = 1, 2, 3
			if isMidVersionBad {
				// B
				// B B
				// ? B B
				rightPartition = midPartition

			} else {
				// G
				// G G
				// ? G G
				leftPartition = midPartition
			}
		}()

		if shouldEndInBreak {
			break
		}
	}

	return minBadVersion
}
