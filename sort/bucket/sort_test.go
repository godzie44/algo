package bucketsort

import (
	"algorithms"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateTopBoundedRandomSlice(t, 100)

		floatArr := make([]float64, len(arr))
		for i, v := range arr {
			floatArr[i] = float64(v) / 100
		}

		Sort(floatArr)
		assert.True(t, sort.Float64sAreSorted(floatArr))
	}
}
