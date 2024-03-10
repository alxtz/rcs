package utils

import (
	"slices"
	"strconv"
	"strings"
)

type Polynomial []string

func ExtractFromTerm(term string) (coe int, indeterminate *string) {
	indeterminate = nil
	coe = 0

	idxSep := strings.Index(term, "*")

	if idxSep != -1 {
		coe, _ = strconv.Atoi(term[:idxSep])
		s := term[idxSep+1:]
		indeterminate = &s
	} else {
		coe, _ = strconv.Atoi(term)
		indeterminate = nil
	}

	return coe, indeterminate
}

func SortPoly(poly Polynomial) Polynomial {
	strPoly := ([]string)(poly)
	input := append([]string{}, strPoly...)
	// lexicographical ordered
	slices.SortFunc(input, func(a, b string) int {
		return strings.Compare(a, b)
	})

	slices.SortStableFunc(input, func(a, b string) int {
		aStrDeg := len(strings.Split(a, "*"))
		bStrDeg := len(strings.Split(b, "*"))

		if aStrDeg < bStrDeg {
			return 1
		} else if aStrDeg > bStrDeg {
			return -1
		}

		return 0
	})

	return input
}
