package main

import (
	"testing"
)

const n = 10

var lookupBottomUp = make([]int, n+1)
var lookupTopDown = make([]int, n+1)

func BenchmarkFibRecursively(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciRecursively(n)
	}
}

func BenchmarkFibBottomUp(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciBottomUp(n, lookupBottomUp)
	}
}

func BenchmarkFibonacciTopDown(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FibonacciTopDown(n, lookupTopDown)
	}
}
