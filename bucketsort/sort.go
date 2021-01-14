package bucketsort

import "math"

func Sort(arr []float64) {
	n := len(arr)
	b := make([][]float64, n+1)

	for _, v := range arr {
		ind := math.Floor(float64(n) * v)
		b[int(ind)] = append(b[int(ind)], v)
	}

	var ind int
	for _, bucket := range b {
		insertionSort(bucket)

		for _, v := range bucket {
			arr[ind] = v
			ind++
		}
	}
}

func insertionSort(arr []float64) {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i+1] = arr[i]
			i--
		}
		arr[i+1] = key
	}
}
