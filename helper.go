package algorithms

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func GenerateRandomSlice(t *testing.T) []int {
	data := make([]int, 10_000)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)
		if rand.Intn(4) == 1 {
			data[i] = data[i] * -1
		}
	}

	assert.False(t, sort.IntsAreSorted(data), "terrible rand.Int")
	return data
}

func GenerateTopBoundedRandomSlice(t *testing.T, maxInt int) []int {
	data := make([]int, 10_000)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(maxInt + 1)
	}

	assert.False(t, sort.IntsAreSorted(data), "terrible rand.Int")
	return data
}
