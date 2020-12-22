package heapsort

import (
	"algorithms"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSortWithRecursiveHeapify(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateRandomSlice(t)
		Sort(arr, MaxHeapifyRecursive)
		assert.True(t, sort.IntsAreSorted(arr))
	}
}

func TestSortWithLinearHeapify(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateRandomSlice(t)
		Sort(arr, MaxHeapifyLinear)
		assert.True(t, sort.IntsAreSorted(arr))
	}
}
