package allocations

func buildOrginalData() []int {
	s := make([]int, 1024)
	for i := range s {
		s[i] = i
	}
	return s
}

func check(v int) bool {
	return v%2 == 0
}

func FilterOneAllocation(data []int) []int {
	var r = make([]int, 0, len(data))
	for _, v := range data {
		if check(v) {
			r = append(r, v)
		}
	}
	return r
}

func FilterNoAllocations(data []int) []int {
	var k = 0
	for i, v := range data {
		if check(v) {
			data[i] = data[k]
			data[k] = v
			k++
		}
	}
	return data[:k]
}
