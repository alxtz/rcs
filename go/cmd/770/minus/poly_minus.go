package minus

import (
	"complicated-prob-solving/cmd/770/utils"
	"strconv"
	"strings"
)

func PolyMinus(leftVal utils.Polynomial, rightVal utils.Polynomial) utils.Polynomial {
	/*
		zeroDegree int
		{
			"a*a": "-1"
			"b*b": "-2"
			"b*c*dd": "100"
			"b*eee": "8"
		}
	*/
	var defaultDegree = 0
	var degreeMap = map[string]int{}

	for _, val := range leftVal {
		polyTerm := val

		var degree *string = nil
		var coE int

		idxSep := strings.Index(polyTerm, "*")

		if idxSep != -1 {
			coE, _ = strconv.Atoi(polyTerm[:idxSep])
			s := polyTerm[idxSep+1:]
			degree = &s
		} else {
			coE, _ = strconv.Atoi(polyTerm)
			degree = nil
		}

		if degree == nil {
			defaultDegree += coE
		} else {
			existingVal, hasExisting := degreeMap[*degree]
			if hasExisting {
				degreeMap[*degree] = existingVal + coE
			} else {
				degreeMap[*degree] = coE
			}
		}
	}

	for _, val := range rightVal {
		polyTerm := val

		var degree *string = nil
		var coE int

		idxSep := strings.Index(polyTerm, "*")

		if idxSep != -1 {
			coE, _ = strconv.Atoi(polyTerm[:idxSep])
			s := polyTerm[idxSep+1:]
			degree = &s
		} else {
			coE, _ = strconv.Atoi(polyTerm)
			degree = nil
		}

		if degree == nil {
			defaultDegree -= coE
		} else {
			existingVal, hasExisting := degreeMap[*degree]
			if hasExisting {
				degreeMap[*degree] = existingVal - coE
			} else {
				degreeMap[*degree] = -coE
			}
		}
	}

	var ans utils.Polynomial

	for degreeKey, coeVal := range degreeMap {
		if coeVal == 0 {
			continue
		}
		coeStr := strconv.Itoa(coeVal)
		ans = append(ans, coeStr+"*"+degreeKey)
	}

	if defaultDegree != 0 {
		ans = append(ans, strconv.Itoa(defaultDegree))
	}

	// fmt.Println(leftVal, rightVal, ans)

	return ans
}
