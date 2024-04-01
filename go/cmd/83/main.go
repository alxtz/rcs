package main

import (
	"fmt"
)

func main() {
	result := deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: nil}}}}})
	// deleteDuplicates(&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: nil}}}}})
	next := result
	for next != nil {
		fmt.Println("next", next.Val)
		next = next.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	currAnchor := head
	nextAnchorCandidate := head

	for {
		nextAnchorCandidate = nextAnchorCandidate.Next

		if nextAnchorCandidate == nil {
			currAnchor.Next = nil
			break
		}

		if nextAnchorCandidate.Val != currAnchor.Val {
			currAnchor.Next = nextAnchorCandidate
			currAnchor = currAnchor.Next
			continue
		}
	}

	return head
}
