package main

import "fmt"

func main() {
	result := reverseBetween(
		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}}}}},
		// 3, 5,
		// 1, 3,
		3, 7,
	)

	fmt.Println(result)
	fmt.Println(result.Next)
	fmt.Println(result.Next.Next)
	fmt.Println(result.Next.Next.Next)
	fmt.Println(result.Next.Next.Next.Next)
	fmt.Println(result.Next.Next.Next.Next.Next)
	fmt.Println(result.Next.Next.Next.Next.Next.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, givenLeft int, givenRight int) *ListNode {
	if head == nil {
		return nil
	}

	if givenLeft == givenRight {
		return head
	}

	// end of neck
	neck := &ListNode{Val: -1000, Next: head}
	count := 0

	for count < givenLeft-1 {
		neck = neck.Next
		count++
	}

	// end of belly
	futureBelly := neck.Next

	var newBellyHead *ListNode
	currBellyHead := neck.Next

	for count < givenRight {
		removed := currBellyHead
		currBellyHead = currBellyHead.Next

		removed.Next = newBellyHead
		newBellyHead = removed

		count++
	}

	neck.Next = newBellyHead
	futureBelly.Next = currBellyHead

	if neck.Val == -1000 {
		return neck.Next
	} else {
		return head
	}
}
