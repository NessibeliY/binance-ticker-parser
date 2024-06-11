package pkg

func DivideSlice(arr []string, n int) [][]string {
	res := make([][]string, n)
	for i, num := range arr {
		resIndex := i % n
		res[resIndex] = append(res[resIndex], num)
	}
	return res
}
