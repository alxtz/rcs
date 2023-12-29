package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var input = []Pair{
		{"a", 3},
		{"b", 1},
		{"a", 3},
		{"c", 2},
		{"d", 1},
		{"e", 2},
		{"f", 1},
		{"g", 2},
		{"x", 3},
		{"y", 1},
		{"x", 3},
	}

	var charSet = make(map[string]bool)

	for _, val := range input {
		if _, ok := charSet[val.Char]; !ok {
			charSet[val.Char] = true
		}
	}

	for key := range charSet {
		var result = toValueSlice(input, key)
		j, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(j))
	}

	// var result = toValueSlice(input)
	// j, _ := json.MarshalIndent(result, "", "  ")
	// fmt.Println(string(j))

	// fmt.Println("x" != "a")
	// fmt.Println("a" != "a")
}

type Container struct {
	Char  string
	Count int
}

type Cost struct {
	AccumCost int
	MaxCost   int
}

type LastOccurEndMap map[string]int

// type DelCostMap map[string]Cost

type DelCostMap map[string]*Cost

func getLengthOfOptimalCompression(s string, k int) int {
	// var tempAns string

	var globalMem *Container = nil
	var lastOccurEndMap LastOccurEndMap = make(LastOccurEndMap)
	var delCostMap DelCostMap = make(DelCostMap)

	// for _, charCode := range s {
	// 	var currentChar = string(charCode)
	// 	// fmt.Println("char", currentChar)

	// 	if globalMem != nil {
	// 		if globalMem.Char != currentChar {
	// 			tempAns += (globalMem.Char + fmt.Sprint(globalMem.Count))
	// 			globalMem = &Container{
	// 				Char:  currentChar,
	// 				Count: 1,
	// 			}
	// 		} else {
	// 			globalMem.Count += 1
	// 		}
	// 	}

	// 	if globalMem == nil {
	// 		globalMem = &Container{
	// 			Char:  currentChar,
	// 			Count: 1,
	// 		}
	// 	}

	// 	// fmt.Println(*globalMem)
	// }

	// tempAns += (globalMem.Char + fmt.Sprint(globalMem.Count))

	// fmt.Println("tempAns", tempAns)

	for index, charCode := range s {
		var newFoundChar = string(charCode)

		if globalMem == nil {
			globalMem = &Container{
				Char:  ".",
				Count: 0,
			}
			continue
		}

		if _, ok := delCostMap[newFoundChar]; !ok {
			delCostMap[newFoundChar] = &Cost{
				AccumCost: 0,
				MaxCost:   0,
			}
		}

		if globalMem.Char != newFoundChar {
			// tempAns += (globalMem.Char + fmt.Sprint(globalMem.Count))
			globalMem = &Container{
				Char:  newFoundChar,
				Count: 1,
			}

			delCostMap[globalMem.Char].AccumCost = 0

			var costObj = delCostMap[newFoundChar]
			// costObj.AccumCost += 1
			if costObj.AccumCost > costObj.MaxCost {
				costObj.MaxCost = costObj.AccumCost
			}

			if _, ok := lastOccurEndMap[globalMem.Char]; !ok {
				lastOccurEndMap[globalMem.Char] = index - 1
			} else {
				costObj.AccumCost = index - lastOccurEndMap[globalMem.Char]
				if costObj.AccumCost > costObj.MaxCost {
					costObj.MaxCost = costObj.AccumCost
				}
			}
		} else {
			// 重複出現 char 了
			globalMem.Count += 1
			// var costObj = delCostMap[newFoundChar]
			// // costObj.AccumCost += 1
			// if costObj.AccumCost > costObj.MaxCost {
			// 	costObj.MaxCost = costObj.AccumCost
			// }
		}

		lastOccurEndMap[newFoundChar] = index
	}

	fmt.Println("lastOccurEndMap", lastOccurEndMap)
	// fmt.Println("delCostMap", delCostMap)

	for k, v := range delCostMap {
		fmt.Println(k, v.MaxCost)
	}

	return 0
}
