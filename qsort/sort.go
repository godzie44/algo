package qsort

import (
	"math/rand"
)

type Partition func(arr []int) int

func RandomizePartition(arr []int, p Partition) int {
	i := rand.Intn(len(arr))
	arr[len(arr)-1], arr[i] = arr[i], arr[len(arr)-1]

	return p(arr)
}

func Lomuto(arr []int) {
	if len(arr) > 1 {
		q := RandomizePartition(arr, LomutoPartition)
		Lomuto(arr[:q])
		Lomuto(arr[q+1:])
	}
}

func LomutoPartition(arr []int) int {
	x := arr[len(arr)-1]
	i := -1
	for j := 0; j < len(arr)-1; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]
	return i + 1
}

func Hoare(arr []int) {
	if len(arr) > 1 {
		q := RandomizePartition(arr, HoarePartition)
		Hoare(arr[:q+1])
		Hoare(arr[q+1:])
	}
}

func HoarePartition(arr []int) int {
	pivot := arr[0]
	i := -1
	j := len(arr)

	for {
		j--
		for ; j >= 0 && arr[j] > pivot; j-- {
		}

		i++
		for ; i < len(arr)-1 && arr[i] < pivot; i++ {
		}

		if i >= j {
			return j
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
}
