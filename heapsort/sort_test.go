package heapsort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortWithRecursiveHeapify(t *testing.T) {
	testCases := []struct {
		Input    []int
		Expected []int
	}{
		{Input: []int{1, 2, 3}, Expected: []int{1, 2, 3}},
		{Input: []int{2, 111, 555, 9}, Expected: []int{2, 9, 111, 555}},
		{Input: []int{1, -2, -3, 4}, Expected: []int{-3, -2, 1, 4}},
		{Input: []int{0, -6, 11, 2, 1, 9, 3}, Expected: []int{-6, 0, 1, 2, 3, 9, 11}},
	}

	for _, tcase := range testCases {
		Sort(tcase.Input, MaxHeapifyRecursive)

		assert.Equal(t, tcase.Expected, tcase.Input)
	}
}

func TestSortWithLinearHeapify(t *testing.T) {
	testCases := []struct {
		Input    []int
		Expected []int
	}{
		{Input: []int{1, 2, 3}, Expected: []int{1, 2, 3}},
		{Input: []int{2, 111, 555, 9}, Expected: []int{2, 9, 111, 555}},
		{Input: []int{1, -2, -3, 4}, Expected: []int{-3, -2, 1, 4}},
		{Input: []int{0, -6, 11, 2, 1, 9, 3}, Expected: []int{-6, 0, 1, 2, 3, 9, 11}},
	}

	for _, tcase := range testCases {
		Sort(tcase.Input, MaxHeapifyLinear)

		assert.Equal(t, tcase.Expected, tcase.Input)
	}
}
