package main

import (
	"fmt"
	"rcs/pkg/ds"
)

func main() {
	fmt.Println("Hello, World")
}

// 1. has a customer queue (with a slice, 0 ~ n)
// 2. also each customer would like to buy 1 ~ 100 tickets (tickets[i])
// 3. given a position k customer, return how much time it takes to satisfy him

// a. buy ticket takes 1s, once brought (1), queueing is required

type Customer = *int

func timeRequiredToBuy(tickets []int, k int) int {
	spentSecs := 0
	var customer Customer

	customerQueue := ds.Queue[Customer]{}

	for index, ticketPref := range tickets {
		ref := ticketPref
		customerQueue.Enque(&ref)
		if index == k {
			customer = &ref
		}
	}

	for {
		// break, when customer is satisfied
		if *customer == 0 {
			break
		}

		toDeque := customerQueue.Deque()
		*toDeque -= 1
		spentSecs += 1

		// still not satisfied
		if *toDeque > 0 {
			customerQueue.Enque(toDeque)
		}
	}

	return spentSecs
}
