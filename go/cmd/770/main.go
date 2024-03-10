package main

import (
	ds "complicated-prob-solving"
	"complicated-prob-solving/cmd/770/add"
	"complicated-prob-solving/cmd/770/minus"
	"complicated-prob-solving/cmd/770/mult"
	"complicated-prob-solving/cmd/770/utils"
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Hello, World")
}

var operatorSequenceMap = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
}

func BasicCalculatorIV(expression string, evalVars []string, evalInts []int) []string {
	indeterminateMap := map[string]bool{}

	for _, indeterminate := range evalVars {
		indeterminateMap[indeterminate] = true
	}

	polyStack := ds.Stack[utils.Polynomial]{Slice: []utils.Polynomial{}}
	oprStack := ds.Stack[string]{Slice: []string{}}

	for _, charCode := range expression {
		char := string(charCode)

		if char == " " {
			continue
		}

		if slices.Contains([]string{"("}, char) {
			oprStack.Push(char)
			continue
		}

		if slices.Contains([]string{"+", "-", "*"}, char) {
			for {
				if oprStack.Len() == 0 {
					break
				}
				lastOpr := oprStack.Peek()

				lastOprSeq, seqExists := operatorSequenceMap[lastOpr]

				if !seqExists {
					break
				}

				currOprSeq := operatorSequenceMap[char]

				if currOprSeq <= lastOprSeq {
					var resultPoly utils.Polynomial

					opr := oprStack.Pop()
					valRight := polyStack.Pop()
					valLeft := polyStack.Pop()

					if opr == "+" {
						resultPoly = add.PolyAdd(
							valLeft,
							valRight,
						)
					} else if opr == "*" {
						resultPoly = mult.PolyMult(
							valLeft,
							valRight,
						)
					} else if opr == "-" {
						resultPoly = minus.PolyMinus(
							valLeft,
							valRight,
						)
					}

					polyStack.Push(resultPoly)
				} else {
					break
				}
			}

			oprStack.Push(char)
			continue
		}

		if slices.Contains([]string{")"}, char) {
			var resultPoly utils.Polynomial

			for {
				if oprStack.Peek() == "(" {
					oprStack.Pop()
					break
				}

				opr := oprStack.Pop()
				valRight := polyStack.Pop()
				valLeft := polyStack.Pop()

				if opr == "+" {
					resultPoly = add.PolyAdd(
						valLeft,
						valRight,
					)
				} else if opr == "*" {
					resultPoly = mult.PolyMult(
						valLeft,
						valRight,
					)
				} else if opr == "-" {
					resultPoly = minus.PolyMinus(
						valLeft,
						valRight,
					)
				}

				polyStack.Push(resultPoly)
			}

			continue
		}

		if isIndeterminate := indeterminateMap[char]; isIndeterminate {
			polyStack.Push(utils.Polynomial{"1*" + char})
		} else {
			polyStack.Push(utils.Polynomial{char})
		}

	}

	var resultPoly utils.Polynomial

	for {
		if oprStack.Len() == 0 {
			break
		}

		opr := oprStack.Pop()
		valRight := polyStack.Pop()
		valLeft := polyStack.Pop()

		if opr == "+" {
			resultPoly = add.PolyAdd(
				valLeft,
				valRight,
			)
		} else if opr == "*" {
			resultPoly = mult.PolyMult(
				valLeft,
				valRight,
			)
		} else if opr == "-" {
			resultPoly = minus.PolyMinus(
				valLeft,
				valRight,
			)
		}

		polyStack.Push(resultPoly)
	}

	return polyStack.Peek()
}
