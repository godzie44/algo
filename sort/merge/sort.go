package merge

import (
	"math"
)

func merge(a []int, q int) {
	left := make([]int, q)
	copy(left, a[:q])
	right := make([]int, len(a)-q)
	copy(right, a[q:])

	left = append(left, math.MaxInt64)
	right = append(right, math.MaxInt64)

	var i, j int
	for k := 0; k < len(a); k++ {
		if left[i] <= right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
		}
	}
}

func Sort(a []int) {
	if len(a) > 1 {
		q := (len(a)) / 2

		Sort(a[0:q])
		Sort(a[q:])

		merge(a, q)
	}
}
