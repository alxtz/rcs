package main

import (
	"fmt"
	"rcs/pkg/ds"
)

func main() {
	fmt.Println("Hello, World")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}

	currDepth := 1
	levelSumMap := make(map[int]int)

	queue := ds.Queue[*TreeNode]{}

	queue.Enque(root)

	for {
		newLevelQueue := ds.Queue[*TreeNode]{}

		for !queue.Empty() {
			removed := queue.Deque()

			if removed.Left != nil {
				newLevelQueue.Enque(removed.Left)
			}

			if removed.Right != nil {
				newLevelQueue.Enque(removed.Right)
			}

			// removed is a leaf
			levelSumMap[currDepth] += removed.Val
		}

		if newLevelQueue.Empty() {
			break
		}

		queue = newLevelQueue
		currDepth += 1
	}

	maxDepth := currDepth

	return levelSumMap[maxDepth]
}
