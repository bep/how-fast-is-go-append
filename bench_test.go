package main

import (
	"fmt"
	"testing"
)

// https://dev.to/uilicious/javascript-array-push-is-945x-faster-than-array-concat-1oki
func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
		b.StartTimer()
		for i := 0; i < 10000; i++ {
			arr1 = append(arr1, arr2...)
		}
		if len(arr1) != 100010 {
			b.Fatal(fmt.Sprintf("was %d", len(arr1)))
		}
	}
}
