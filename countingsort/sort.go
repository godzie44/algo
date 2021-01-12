package heapsort

func Sort(arr []int, k int) []int {
	c := make([]int, k+1)

	for _, v := range arr {
		c[v]++
	}

	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	result := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		result[c[arr[i]]-1] = arr[i]
		c[arr[i]]--
	}

	return result
}
