package priorityqueue

import (
	"algorithms"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestQueueMax(t *testing.T) {
	randSlice, sortedSlice := prepareRandAndSortedArray(t)

	var queue PriorityQueue
	for _, v := range randSlice {
		queue.Insert(&QueueElement{Key: v, Value: nil})
	}

	assert.Equal(t, sortedSlice[0], queue.Maximum().Key)
}

func TestQueueExtractMax(t *testing.T) {
	randSlice, sortedSlice := prepareRandAndSortedArray(t)

	var queue PriorityQueue
	for _, v := range randSlice {
		queue.Insert(&QueueElement{Key: v, Value: nil})
	}

	queueResult := make([]int, 0, len(randSlice))
	for {
		v, err := queue.ExtractMaximum()
		if err != nil {
			break
		}

		queueResult = append(queueResult, v.Key)
	}

	assert.Equal(t, sortedSlice, queueResult)
	assert.Len(t, queue.a, 0)
}

func prepareRandAndSortedArray(t *testing.T) ([]int, []int) {
	randSlice := algorithms.GenerateRandomSlice(t)
	sortedSlice := make([]int, len(randSlice))
	copy(sortedSlice, randSlice)
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i] > sortedSlice[j]
	})

	return randSlice, sortedSlice
}
