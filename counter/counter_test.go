package counter_test

import (
	"go-training/counter"
	"testing"
)

var c int

func BenchmarkCounter(b *testing.B) {
	b.ReportAllocs()

	b.Run("aa", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c = counter.StressTest()
		}
	})

	b.Run("bb", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c = counter.StressTest()
		}
	})
}
