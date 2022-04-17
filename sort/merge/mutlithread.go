package merge

import (
	"math"
	"sync"
)

func binarySearch(x int, arr []int) int {
	low := 0
	high := len(arr)

	for low < high {
		mid := int(math.Floor(float64(low+high) / 2))

		if x <= arr[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return high
}

func pMerge(t []int, p1, r1, p2, r2 int, a []int) {
	n1 := r1 - p1 + 1
	n2 := r2 - p2 + 1

	if n1 < n2 {
		p1, p2 = p2, p1
		r1, r2 = r2, r1
		n1, n2 = n2, n1
	}

	if n1 == 0 {
		return
	}

	q1 := int(math.Floor(float64(p1+r1) / 2))
	q2 := binarySearch(t[q1], t[p2:r2+1])
	q3 := (q1 - p1) + q2
	a[q3] = t[q1]

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		pMerge(t, p1, q1-1, p2, p2+q2-1, a)
	}()
	pMerge(t, q1+1, r1, p2+q2, r2, a[q3+1:])

	wg.Wait()
}

func ParallelMergeSort(a []int, p, r int, b []int, s int) {
	n := r - p + 1
	if n == 1 {
		b[s] = a[p]
		return
	}

	t := make([]int, n)
	q := int(math.Floor(float64(p+r) / 2))
	qq := q - p + 1

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		ParallelMergeSort(a, p, q, t, 0)
	}()
	ParallelMergeSort(a, q+1, r, t, qq)

	wg.Wait()

	pMerge(t, 0, qq-1, qq, n-1, b[s:])
}
