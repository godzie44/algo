package heapsort

import (
	"algorithms"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateTopBoundedRandomSlice(t, 100)
		res := Sort(arr, 100)
		assert.True(t, sort.IntsAreSorted(res))
	}
}
