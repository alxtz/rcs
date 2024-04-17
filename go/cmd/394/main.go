package main

import (
	"fmt"
	"strconv"
)

func main() {
	// fmt.Println("Hello, World")

	// decodeString("")

	// candidate, _ := recur("[3[a2[c]]d]")
	// fmt.Println("candidate", candidate)

	candidateY1, _ := recur("1[3[a2[c]]d]")
	fmt.Println("candidateY1", candidateY1)

	candidateY2, _ := recur("1[c]")
	fmt.Println("candidateY2", candidateY2 == "c")

	candidateY3, _ := recur("c")
	fmt.Println("candidateY3", candidateY3 == "c")
}

// "3[a2[c]]d" -> extract "3["
// "a2[c]]d" -> extract "a2["
// "c]]d" -> base case, returns multCandidate = "c", remaining = "]d"

// func decodeString(s string) string {
// 	return "1["
// }

func recur(s string) (multCandidate string, remaining string) {
	fmt.Println("<stack> got", s)
	inputStr := s

	outputStr := ""

	numStr := ""

	for {
		// fmt.Println("loop", inputStr, outputStr, numStr)

		if len(inputStr) == 0 {
			break
		}
		extracted := inputStr[0]
		inputStr = inputStr[1:]

		if extracted == ']' {
			return outputStr, inputStr
		} else if 'a' <= extracted && extracted <= 'z' {
			outputStr += string(extracted)
			continue
		} else if '0' <= extracted && extracted <= '9' {
			numStr += string(extracted)
		} else if extracted == '[' {
			fmt.Println("about to give", inputStr)
			multCandidate, remaining = recur(inputStr)
			fmt.Println("--got child", multCandidate, "remaining", inputStr)
			multFactor, _ := strconv.Atoi(numStr)
			outputStr += dup(multCandidate, multFactor)
			numStr = "" // cleanup
			inputStr = remaining
			continue
		} else {
			break // should not happen
		}
	}

	return outputStr, ""

	// return "", "" // should not happen
}

func dup(candidate string, factor int) (ans string) {
	for i := 1; i <= factor; i++ {
		ans += candidate
	}

	return ans
}
