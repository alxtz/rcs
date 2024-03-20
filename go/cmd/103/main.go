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

type DoubleEndedQueue struct{ ds.Queue[*TreeNode] }

func (mq *DoubleEndedQueue) PushBack(val *TreeNode) {
	mq.Slice = append(mq.Slice, val)
}

func (mq *DoubleEndedQueue) PushFront(val *TreeNode) {
	mq.Slice = append([]*TreeNode{val}, mq.Slice...)
}

func (mq *DoubleEndedQueue) PopBack() (removed *TreeNode) {
	if !mq.Empty() {
		removed = mq.Slice[len(mq.Slice)-1]
		mq.Slice = mq.Slice[:len(mq.Slice)-1]
	}

	return removed
}

func (mq *DoubleEndedQueue) PopFront() (removed *TreeNode) {
	if !mq.Empty() {
		removed = mq.Slice[0]
		mq.Slice = mq.Slice[1:]
	}

	return removed
}

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return ans
	}

	fifoFlag := true
	queue := DoubleEndedQueue{}
	queue.Enque(root)

	for {
		newLevelQueue := DoubleEndedQueue{}

		ans = append(ans, []int{})

		for !queue.Empty() {
			var removed *TreeNode

			fmt.Println("for.for.queue", queue)

			removed = queue.PopFront()

			if fifoFlag {
				ans[len(ans)-1] = append(ans[len(ans)-1], removed.Val)
			} else {
				ans[len(ans)-1] = append([]int{removed.Val}, ans[len(ans)-1]...)
			}

			if removed.Left != nil {
				newLevelQueue.PushBack(removed.Left)
			}

			if removed.Right != nil {
				newLevelQueue.PushBack(removed.Right)
			}
		}

		if newLevelQueue.Empty() {
			break
		}

		queue = newLevelQueue
		fifoFlag = !fifoFlag
	}

	return ans
}
