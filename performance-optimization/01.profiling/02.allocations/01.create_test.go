package allocations

import "testing"

func Benchmark_CreateCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreateCopy()
	}
}

func Benchmark_CreatePointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CreatePointer()
	}
}
