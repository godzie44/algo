package qsort

import (
	"algorithms"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSortWithLomutoPartition(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateRandomSlice(t)
		Lomuto(arr)
		assert.True(t, sort.IntsAreSorted(arr))
	}
}

func TestSortWithHoarePartition(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateRandomSlice(t)
		Hoare(arr)
		assert.True(t, sort.IntsAreSorted(arr))
	}
}
