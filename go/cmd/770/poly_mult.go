package main

import (
	"strconv"
	"strings"
)

func PolyMult(leftVal Polynomial, rightVal Polynomial) Polynomial {
	if len(leftVal) == 1 && len(rightVal) == 1 {
		var leftCoEff int
		var rightCoEff int

		leftStr := (leftVal)[0]
		rightStr := (rightVal)[0]

		idxLeft := strings.Index(leftStr, "*")
		idxRight := strings.Index(rightStr, "*")

		remainingLeft := ""
		if idxLeft != -1 {
			leftCoEff, _ = strconv.Atoi(leftStr[:idxLeft])
			remainingLeft = leftStr[idxLeft:]
		} else {
			leftCoEff, _ = strconv.Atoi(leftStr)
		}

		remainingRight := ""
		if idxRight != -1 {
			rightCoEff, _ = strconv.Atoi(rightStr[:idxRight])
			remainingRight = rightStr[idxRight:]
		} else {
			rightCoEff, _ = strconv.Atoi(rightStr)
		}

		finalCoE := strconv.Itoa(leftCoEff * rightCoEff)

		return Polynomial{finalCoE + remainingLeft + remainingRight}
	}

	var bulkPolynomial = []string{}

	for _, iVal := range leftVal {
		for _, jVal := range rightVal {
			var computedTerm = PolyMult(Polynomial{iVal}, Polynomial{jVal})

			bulkPolynomial = append(bulkPolynomial, computedTerm...)
		}
	}

	var packedPoly = bulkPolynomial
	var sortedPoly = packedPoly

	return sortedPoly
}
