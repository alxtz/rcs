package main

import "fmt"

func main() {
	result := reverseList(
		&ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	)

	fmt.Println("result", result)
	fmt.Println("result.N", result.Next)
	fmt.Println("result.N.N", result.Next.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newLinkedListHead *ListNode

	for {
		toRemove := head
		nextHead := head.Next
		toRemove.Next = nil

		if newLinkedListHead == nil {
			newLinkedListHead = toRemove
		} else {
			toRemove.Next = newLinkedListHead
			newLinkedListHead = toRemove
		}

		if nextHead == nil {
			break
		} else {
			head = nextHead
		}
	}

	return newLinkedListHead
}
