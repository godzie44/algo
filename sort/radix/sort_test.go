package radix

import (
	"algorithms"
	"github.com/stretchr/testify/assert"
	"math"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateTopBoundedRandomSlice(t, math.MaxInt32)
		int32Arr := make([]int32, len(arr))
		for i := range arr {
			int32Arr[i] = int32(arr[i])
		}
		res := Sort(int32Arr)

		intRes := make([]int, len(res))
		for i := range res {
			intRes[i] = int(res[i])
		}

		assert.True(t, sort.IntsAreSorted(intRes))
	}
}
