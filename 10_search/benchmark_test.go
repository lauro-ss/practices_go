package main

import (
	"math"
	"math/rand"
	"testing"
)

func BenchmarkBinarySearch(b *testing.B) {
	const SIZE = math.MaxInt16

	list := make([]int32, SIZE)
	for i := range list {
		list[i] = int32(i + 1)
	}
	for i := 0; i < b.N; i++ {
		binarySearch(list, rand.Int31n(SIZE))
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	const SIZE = math.MaxInt16

	list := make([]int32, SIZE)
	for i := range list {
		list[i] = int32(i + 1)
	}
	for i := 0; i < b.N; i++ {
		linearSearch(list, rand.Int31n(SIZE))
	}
}
