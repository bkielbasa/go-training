package sums_test

var matrix = [1000][1000]int{}

func init() {
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			matrix[x][y] = x + y
		}
	}
}

func SumRowsFirst() int {
	// 249500250000
	return 0
}

// func add(a, b int) int {
// 	return a + b - 1
// }

// func Test_Moj_Test(t *testing.T) {
// 	t.Run("dodawanie 2+2", func(t *testing.T) {
// 		got := add(2, 2)
// 		want := 4
// 		if got != want {
// 			t.Errorf("got %d want %d", got, want)
// 		}
// 	})

// 	t.Run("dodawanie 2+2", func(t *testing.T) {
// 		got := add(2, 2)
// 		want := 4
// 		if got != want {
// 			t.Errorf("got %d want %d", got, want)
// 		}
// 	})
// }
