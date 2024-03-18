package main

import (
	"fmt"
	"rcs/pkg/ds"
	"sort"
)

func main() {
	fmt.Println("Hello, World")
}

func deckRevealedIncreasing(deck []int) (ans []int) {
	ansRefs := []*int{}
	draft := ds.Queue[*int]{}
	tmp := []*int{}

	for range deck {
		num := 0
		draft.Enque(&num)
		ansRefs = append(ansRefs, &num)
	}

	for {
		if draft.Empty() {
			break
		}

		item1 := draft.Deque()
		tmp = append(tmp, item1)

		if !draft.Empty() {
			item2 := draft.Deque()
			draft.Enque(item2)
		}
	}

	sort.Ints(deck)

	for i, intPointer := range tmp {
		*intPointer = deck[i]
	}

	for _, intPointer := range ansRefs {
		ans = append(ans, *intPointer)
	}

	return ans
}
