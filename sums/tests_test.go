package sums

import (
	"testing"
)

var matrix = [1000][1000]int{}

func init() {
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			matrix[x][y] = x + y
		}
	}
}

func SumRowsFirst() int {
	sum := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			sum += matrix[x][y]
		}
	}

	return sum
}

func SumColsFirst() int {
	sum := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			sum += matrix[y][x]
		}
	}
	return sum
}

func SumRange() int {
	sum := 0
	for _, row := range matrix {
		for _, c := range row {
			sum += c
		}
	}
	return sum
}

func SumRange2() int {
	sum := 0
	for i := range matrix {
		for j := range matrix[i] {
			sum += matrix[i][j]
		}
	}
	return sum
}

func Test_SumRowsFirst(t *testing.T) {
	t.Parallel()

	expected := 999000000
	got := SumRowsFirst()
	if got != expected {
		t.Errorf("got %d want %d", got, expected)
	}
}

func Test_SumRange(t *testing.T) {
	t.Parallel()

	expected := 999000000
	got := SumRange()
	if got != expected {
		t.Errorf("got %d want %d", got, expected)
	}

}

var ignore = 0

func Benchmark_SumRowsFirst(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ignore = SumRowsFirst()
	}
}

func Benchmark_SumColsFirst(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ignore = SumColsFirst()
	}
}

func Benchmark_SumRange(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ignore = SumRange()
	}
}

func Benchmark_SumRange2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ignore = SumRange2()
	}
}
