package main

import (
	"testing"
)

const n = 10

var lookup = make([]int, n+1)

func BenchmarkFibRecursively(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciRecursively(n)
	}
}

func BenchmarkFibBottomUp(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciBottomUp(n, lookup)
	}
}

func BenchmarkFibonacciTopDown(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciTopDown(n, lookup)
	}
}
