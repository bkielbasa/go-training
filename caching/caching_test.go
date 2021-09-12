// go test -run none -bench . -benchtime 3s

package caching

import "testing"

var fa int

func BenchmarkLinkListTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = LinkedListTraverse()
	}

	fa = a
}

func BenchmarkColumnTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = ColumnTraverse()
	}

	fa = a
}

func BenchmarkRowTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = RowTraverse()
	}

	fa = a
}
