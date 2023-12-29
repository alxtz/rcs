package main

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()

	// var result_a = subsets_1([]int{1, 2, 3})
	// fmt.Println(len(result_a) == 8)

	// var result_b = subsets_1([]int{1, 2})
	// fmt.Println(len(result_b) == 4)

	var result_c = subsets_1([]int{1, 2, 3, 4})
	fmt.Println(len(result_c) == 16)

	// subsets_1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	// fmt.Println(len(result_c) == 16)
}

func subsets_1(nums []int) [][]int {
	var mainMem = [][]int{}

	var knownMap = make(map[string][]int)

	findPosib(nums, &mainMem, &knownMap)

	mainMem = append(mainMem, nums)

	return mainMem
}

func findPosib(nums []int, mem *[][]int, knownMap *map[string][]int) {
	fmt.Println("run")
	for _, element := range nums {
		var tempNums = slices.Clone(nums)

		newNums := slices.DeleteFunc(tempNums, func(i int) bool {
			return i == element
		})

		key, _ := json.Marshal(newNums)

		if (*knownMap)[string(key)] == nil {
			*mem = append(*mem, newNums)
			(*knownMap)[string(key)] = newNums

			findPosib(newNums, mem, knownMap)
		}
	}
}
