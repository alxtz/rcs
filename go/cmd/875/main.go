package main

import (
	"fmt"
)

func main() {
	fmt.Println("minEatingSpeed([3,6,7,11], 8) = 4", minEatingSpeed(
		[]int{3, 6, 7, 11}, 8,
	) == 4)

	fmt.Println("minEatingSpeed([30,11,23,4,20], 5) = 30", minEatingSpeed(
		[]int{30, 11, 23, 4, 20}, 5,
	) == 30)
}

func minEatingSpeed(piles []int, givenHours int) (minEatSpeed int) {
	maxPileVal := -1

	for _, singlePileVal := range piles {
		if singlePileVal > maxPileVal {
			maxPileVal = singlePileVal
		}
	}

	minEatSpeed = maxPileVal

	leftAnchor := 1
	rightAnchor := maxPileVal

	for {

		lastLoop := (rightAnchor - leftAnchor) <= 1

		stimulatedMid := leftAnchor + (rightAnchor-leftAnchor)/2

		stimulatedFinishHr := 0

		for _, singlePileVal := range piles {
			base, remainder := singlePileVal/stimulatedMid, singlePileVal%stimulatedMid
			if remainder > 0 {
				base += 1
			}
			stimulatedFinishHr += base
		}

		if stimulatedFinishHr <= givenHours {
			minEatSpeed = stimulatedMid
			rightAnchor = stimulatedMid
		} else if givenHours < stimulatedFinishHr {
			leftAnchor = stimulatedMid
		}

		if lastLoop {
			break
		}
	}

	return minEatSpeed
}
