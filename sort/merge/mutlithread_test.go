package merge

import (
	"algorithms"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestParallelSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateRandomSlice(t)

		b := make([]int, len(arr))
		ParallelMergeSort(arr, 0, len(arr)-1, b, 0)

		assert.True(t, sort.IntsAreSorted(b))
	}
}
