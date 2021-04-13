package sel

import (
	"algorithms/sort/insertion"
	"algorithms/sort/qsort"
)

func Randomized(arr []int, ind int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	q := qsort.RandomizePartition(arr, qsort.LomutoPartition)

	if ind == q {
		return arr[q]
	} else if ind < q {
		return Randomized(arr[:q], ind)
	} else {
		return Randomized(arr[q:], ind-q)
	}
}

func Select(arr []int, ind int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	var med int

	medians := findMedians(arr)
	for {
		if len(medians) <= 5 {
			med = midEl(medians)
			break
		}

		medians = findMedians(medians)
	}

	k := partition(arr, med)

	if ind == k {
		return med
	} else if ind < k {
		return Select(arr[:k], ind)
	} else {
		return Select(arr[k+1:], ind-k-1)
	}
}

func findMedians(arr []int) []int {
	groups := make([][]int, len(arr)/5+1)

	for i, v := range arr {
		groups[i/5] = append(groups[i/5], v)
	}

	var medians []int
	for i := range groups {
		if len(groups[i]) == 0 {
			continue
		}

		insertion.Sort(groups[i])
		medians = append(medians, midEl(groups[i]))
	}

	return medians
}

func partition(arr []int, pivot int) int {
	var pivotI int
	for i := range arr {
		if arr[i] == pivot {
			pivotI = i
			break
		}
	}
	arr[pivotI], arr[len(arr)-1] = arr[len(arr)-1], arr[pivotI]

	i := -1
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]
	return i + 1
}

func midEl(arr []int) int {
	if len(arr)%2 == 0 {
		return arr[(len(arr)-1)/2]
	} else {
		return arr[(len(arr))/2]
	}
}
