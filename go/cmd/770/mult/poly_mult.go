package mult

import (
	"complicated-prob-solving/cmd/770/utils"
	"fmt"
	"strconv"
	"strings"
)

type Polynomial []string

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

	fmt.Println("bulkPoly", bulkPolynomial)

	// TODO: implementing packing
	termsMap := map[string]int{}

	defaultCoe := 0

	for _, termStr := range bulkPolynomial {
		coe, indeterminate := utils.ExtractFromTerm(termStr)

		if indeterminate == nil {
			defaultCoe += coe
			continue
		}

		termsMap[*indeterminate] += coe
	}

	fmt.Println("map", termsMap)

	packedPoly := []string{}
	for kIndeterminate, vCoe := range termsMap {
		if vCoe != 0 {
			packedPoly = append(packedPoly, strconv.Itoa(vCoe)+"*"+kIndeterminate)
		}
	}

	if defaultCoe != 0 {
		packedPoly = append(packedPoly, strconv.Itoa(defaultCoe))
	}

	var sortedPoly = packedPoly

	return sortedPoly
}
