package binary_test

import (
	"algorithms/binary"
	"testing"
)

func BenchmarkBinarySearch(b *testing.B) {
	//arr := []int{133,44,77,88}
	for i := 0; i < b.N; i++ {
		if res := binary.BinarySearch([]int{133,444,777,888}, 444); res != 1 {
			b.Fatal("unexpected result: ...", res)
		}
	}
}
