package pkg

func DivideSlice(arr []string, n int) [][]string {
	res := make([][]string, n)
	for i, num := range arr {
		resIndex := i % n
		res[resIndex] = append(res[resIndex], num)
	}
	return res
}

func CompareSlicesOfSlices[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
