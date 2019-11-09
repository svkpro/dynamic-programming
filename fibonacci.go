package main

import (
	"fmt"
)

func main() {

	n := 10
	lookup := make([]int, n+1)

	fmt.Println(FibonacciRecursively(n))
	fmt.Println(FibonacciBottomUp(n, lookup))
	fmt.Println(FibonacciTopDown(n, lookup))
}

func FibonacciRecursively(n int) int {
	if n <= 1 {
		return n
	}

	return FibonacciRecursively(n-1) + FibonacciRecursively(n-2)
}

func FibonacciBottomUp(n int, lookup []int) int {

	lookup[1] = 1

	for i := 2; i <= n; i++ {
		lookup[i] = lookup[i-1] + lookup[i-2]
	}

	return lookup[n]
}

func FibonacciTopDown(n int, lookup []int) int {

	if lookup[n] == 0 {
		if n <= 1 {
			lookup[n] = n;
		} else {
			lookup[n] = FibonacciTopDown(n-1, lookup) + FibonacciTopDown(n-2, lookup);
		}
	}

	return lookup[n];
}
