package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")

	fmt.Println("mySqrt(1)", mySqrt(1))
	fmt.Println("mySqrt(2)", mySqrt(2))
	fmt.Println("mySqrt(4)", mySqrt(4))
	fmt.Println("mySqrt(6)", mySqrt(6))
	fmt.Println("mySqrt(9)", mySqrt(9))
	fmt.Println("mySqrt(11)", mySqrt(11))
	fmt.Println("mySqrt(16)", mySqrt(16))
}

func mySqrt(x int) (nearestAns int) {
	if x == 0 {
		return 0
	}

	leftAnchor := 1
	rightAnchor := x

	for {
		shouldStop := (rightAnchor - leftAnchor) <= 1

		mid := leftAnchor + (rightAnchor-leftAnchor)/2

		midSquare := mid * mid

		if midSquare == x {
			nearestAns = mid
			return nearestAns
		} else if x < midSquare {
			rightAnchor = mid
		} else if midSquare < x {
			leftAnchor = mid
		}

		if shouldStop {
			break
		}
	}

	nearestAns = leftAnchor

	return nearestAns
}
