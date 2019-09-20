package main

import (
	"fmt"
)

func main() {

	n := 3
	lookup := make([]int, n+1)

	fmt.Println(FibonacciRecursively(n))
	fmt.Println(FibonacciBottomUp(n))
	fmt.Println(FibonacciTopDown(n, lookup))
}

func FibonacciRecursively(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursively(n-1) + FibonacciRecursively(n-2)
}

func FibonacciBottomUp(n int) int {
	fs := make([]int, n+2)

	fs[1] = 1

	for i := 2; i <= n; i++ {
		fs[i] = fs[i-1] + fs[i-2]
	}
	return fs[n]
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

