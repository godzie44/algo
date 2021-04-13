package insertion

import (
	"algorithms"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateRandomSlice(t)
		Sort(arr)
		assert.True(t, sort.IntsAreSorted(arr))
	}
}
