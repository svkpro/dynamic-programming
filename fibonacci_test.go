package main

import (
	"testing"
)

const n = 100

var lookup = make([]int, 101)

func BenchmarkFibRecursively(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciBottomUp(n)
	}
}

func BenchmarkFibBottomUp(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciBottomUp(n)
	}
}

func BenchmarkFibonacciTopDown(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciTopDown(n, lookup)
	}
}
