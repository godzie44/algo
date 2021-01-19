package sel

import (
	"algorithms"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestRandomizedSelect(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateRandomSlice(t)

		rand.Seed(time.Now().UnixNano())
		randPosition := rand.Intn(10_000)

		result := Randomized(arr, randPosition)

		sort.Ints(arr)

		assert.Equal(t, arr[randPosition], result)
	}
}

func TestSelect(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := algorithms.GenerateRandomSlice(t)

		rand.Seed(time.Now().UnixNano())
		randPosition := rand.Intn(len(arr))

		result := Select(arr, randPosition)

		sort.Ints(arr)

		assert.Equal(t, arr[randPosition], result)
	}
}
