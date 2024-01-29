package main

import (
	ds "complicated-prob-solving"
	"fmt"
)

func main() {
	fmt.Println("Hello, World")
}

type Polynomial []string

func BasicCalculatorIV(expression string, evalVars []string, evalInts []int) []string {

	numStack := ds.Stack[Polynomial]{Slice: []Polynomial{}}
	oprStack := ds.Stack[string]{Slice: []string{}}

	return nil
}
