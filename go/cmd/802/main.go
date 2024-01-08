package main

import (
	"fmt"
	"slices"
	"sort"
	// mapset "github.com/deckarep/golang-set/v2"
)

func main() {
	fmt.Println("Hello, World")
}

type DirectedGraph struct {
	Val          int
	OutNeighbors []*DirectedGraph
	// OutNeighbors      mapset.Set[*DirectedGraph]
	InNeighborIndexes []int
	// X                 mapset.Set[int]
}

func eventualSafeNodes(graph [][]int) []int {
	var isSafeMap = make(map[int]bool)

	var allGraphs []*DirectedGraph

	for i := range graph {
		isSafeMap[i] = true
		allGraphs = append(allGraphs, &DirectedGraph{
			Val: i,
		})
	}

	// fmt.Println("is", isSafeMap)

	for graphIndex, outs := range graph {
		var currGraph = allGraphs[graphIndex]
		for _, outIndex := range outs {
			currGraph.OutNeighbors = append(currGraph.OutNeighbors, allGraphs[outIndex])
			// currGraph.OutNeighbors.Add(allGraphs[outIndex])
		}
	}

	for graphIndex, outs := range graph {
		// var currGraph = allGraphs[graphIndex]
		for _, outIndex := range outs {
			allGraphs[outIndex].InNeighborIndexes = append(allGraphs[outIndex].InNeighborIndexes, graphIndex)
		}
	}

	for i := range allGraphs {
		// fmt.Println("run-")
		var srtPoint = allGraphs[i]

		var visitedMap = make(map[int]bool)

		var visitingStack = []*DirectedGraph{srtPoint}

		for {
			// fmt.Println("run--")
			if len(visitingStack) == 0 {
				break
			}

			var tail = visitingStack[len(visitingStack)-1]

			// fmt.Println("traveled", tail.Val)

			visitedMap[tail.Val] = true

			if len(tail.OutNeighbors) > 0 {
				var toAdd = tail.OutNeighbors[len(tail.OutNeighbors)-1]

				if _, visited := visitedMap[toAdd.Val]; visited {
					// fmt.Println("broke promise", toAdd.Val)
					isSafeMap[i] = false
					break
				}

				visitingStack = append(visitingStack, toAdd)

				continue
			}

			if len(tail.OutNeighbors) == 0 {
				var valPopped = tail.Val
				// pop itself from the stack
				visitingStack = visitingStack[:len(visitingStack)-1]

				// removing all references pointing to the popped item
				var allRefs = allGraphs[tail.Val].InNeighborIndexes
				for _, key := range allRefs {
					allGraphs[key].OutNeighbors = slices.DeleteFunc(allGraphs[key].OutNeighbors, func(obj *DirectedGraph) bool {
						return obj.Val == valPopped
					})
				}

				// // pop (itself) reference from the parent stack
				// if len(visitingStack) > 0 {
				// 	var newTail = visitingStack[len(visitingStack)-1]
				// 	newTail.OutNeighbors = newTail.OutNeighbors[:len(newTail.OutNeighbors)-1]
				// }
				continue
			}

		}

		// fmt.Println("-----")
	}

	// fmt.Println("safe candidates", isSafeMap)

	// for i := range allGraphs {
	// 	fmt.Println("concreteGraph", allGraphs[i])
	// }

	var ans []int

	for key, val := range isSafeMap {
		if val {
			ans = append(ans, key)
		}
	}

	sort.Ints(ans)

	return ans
}
