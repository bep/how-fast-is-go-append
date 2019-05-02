package main

import (
	"testing"
)

// https://dev.to/uilicious/javascript-array-push-is-945x-faster-than-array-concat-1oki
func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		slice := createSlice(10)
		b.StartTimer()
		var result []int
		for i := 0; i < 10000; i++ {
			result = append(result, slice...)
		}
		if len(result) != 100000 {
			b.Fatal()
		}
	}
}

func createSlice(size int) []int {
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i + 1
	}
	return s
}
