package allocations

import "testing"

func Benchmark_FilterOneAllocation(b *testing.B) {
	data := buildOrginalData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterOneAllocation(data)
	}
}

func Benchmark_FilterNoAllocations(b *testing.B) {
	data := buildOrginalData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = FilterNoAllocations(data)
	}
}
