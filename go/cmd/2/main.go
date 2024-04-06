package main

import (
	"fmt"
)

func main() {
	ans1 := addTwoNumbers(
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
	)
	fmt.Println(
		"addTwoNumbers(2->4->3, 5->6->4)",
		ans1, ans1.Next, ans1.Next.Next,
	)
	ans2 := addTwoNumbers(
		&ListNode{9, &ListNode{9, &ListNode{9, nil}}},
		&ListNode{9, &ListNode{9, nil}},
	)
	fmt.Println(
		"addTwoNumbers(9->9->9, 9->9)",
		ans2, ans2.Next, ans2.Next.Next, ans2.Next.Next.Next,
	)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := ListNode{Val: -1, Next: nil}
	ansTailPtr := &dummyHead

	inheritedInc := 0

	for {
		leftVal := 0
		rightVal := 0

		if l1 == nil && l2 == nil && inheritedInc == 0 {
			break
		}

		if l1 != nil {
			leftVal = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			rightVal = l2.Val
			l2 = l2.Next
		}

		levelSum := leftVal + rightVal + inheritedInc

		if inc := levelSum / 10; inc >= 1 {
			inheritedInc = inc
		} else {
			inheritedInc = 0
		}

		remainder := levelSum % 10

		newNode := ListNode{Val: remainder, Next: nil}

		ansTailPtr.Next = &newNode
		ansTailPtr = ansTailPtr.Next
	}

	return dummyHead.Next
}
