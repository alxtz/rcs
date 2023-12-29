// // func subsets(nums []int) [][]int {
// // 	return [][]int{[]int{1}}
// // 	return [][]int{{1}, {2, 3}, {4, 5, 6}}
// // }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"slices"
// )

// func main() {
// 	// 	Given an integer array nums of unique elements, return all possible
// 	// subsets
// 	//  (the power set).

// 	// The solution set must not contain duplicate subsets. Return the solution in any order.

// 	// var result_a = subsets_1([]int{1, 2, 3})
// 	// fmt.Println(len(result_a) == 8)

// 	var result_b = subsets_1([]int{1, 2})
// 	fmt.Println(len(result_b) == 4)

// 	var result_c = subsets_1([]int{1, 2, 3, 4})
// 	fmt.Println(len(result_c) == 16)
// }

// type Node struct {
// 	Val      []int  `json:"val"`
// 	Children []Node `json:"children"`
// }

// func subsets_1(nums []int) [][]int {
// 	var rootNode = Node{
// 		Val:      nums,
// 		Children: genChildren(nums),
// 	}

// 	// fmt.Println(rootNode)

// 	rN, err := json.MarshalIndent(rootNode, "", "  ")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(rN))

// 	return [][]int{{1}, {2, 3}, {4, 5, 6}}
// }

// func genChildren(nums []int) []Node {
// 	var nodes []Node

// 	for _, element := range nums {
// 		var temp = slices.Clone(nums)

// 		newVal := slices.DeleteFunc(temp, func(i int) bool {
// 			return i == element
// 		})

// 		newChildNode := Node{
// 			Val: newVal,
// 			// Children: []Node{},
// 			Children: genChildren(newVal),
// 		}

// 		nodes = append(nodes, newChildNode)
// 	}

// 	return nodes
// }