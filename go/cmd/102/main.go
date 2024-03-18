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

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return ans
	}

	queue := ds.Queue[*TreeNode]{}
	queue.Enque(root)

	for {
		if queue.Empty() {
			break
		}

		nextLevelQueue := ds.Queue[*TreeNode]{}
		levelAns := []int{}

		for !queue.Empty() {
			removed := queue.Deque()
			fmt.Println("removed", removed.Val)
			levelAns = append(levelAns, removed.Val)
			if removed.Left != nil {
				nextLevelQueue.Enque(removed.Left)
			}
			if removed.Right != nil {
				nextLevelQueue.Enque(removed.Right)
			}
		}

		queue = nextLevelQueue

		ans = append(ans, levelAns)
	}

	return ans
}
