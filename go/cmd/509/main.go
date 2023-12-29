package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World 509")

	// F(0) = 0, F(1) = 1
	// F(n) = F(n - 1) + F(n - 2), for n > 1.

	fmt.Println(fib(0) == 0)
	fmt.Println(fib(1) == 1)
	fmt.Println(fib(2) == 1)
	fmt.Println(fib(3) == 2)
	fmt.Println(fib(4) == 3)
	fmt.Println(fib(5) == 5)
	fmt.Println(fib(6) == 8)
	fmt.Println(fib(7) == 13)
}

func fib(n int) int {
	if n < 0 {
		panic("n < 0")
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
