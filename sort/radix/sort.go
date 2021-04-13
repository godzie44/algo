package radix

const (
	byteMask = 0xFF
)

func Sort(arr []int32) []int32 {
	for i := 0; i < 4; i++ {
		arr = cSort(arr, i)
	}

	return arr
}

func cSort(arr []int32, byteNumber int) []int32 {
	c := make([][]int32, 256)

	shift := byteNumber * 8
	for _, v := range arr {
		byteVal := byte(v & (byteMask << shift) >> shift)
		c[byteVal] = append(c[byteVal], v)
	}

	result := make([]int32, 0, len(arr))
	for _, values := range c {
		for _, v := range values {
			result = append(result, v)
		}
	}

	return result
}
