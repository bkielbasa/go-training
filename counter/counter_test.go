package counter_test

import (
	"go-training/counter"
	"testing"
)

var c int

func BenchmarkCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c = counter.StressTest()
	}
}
