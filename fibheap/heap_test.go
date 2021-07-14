package fibheap

import (
	"algorithms"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestHeapInsert(t *testing.T) {
	h := NewHeap()

	h.Insert(1)
	h.Insert(-1)
	h.Insert(3)
	h.Insert(6)
	h.Insert(12)

	assert.Equal(t, 5, h.n)
	assert.Equal(t, -1, h.Min())
}

func TestHeapRemove(t *testing.T) {
	h := NewHeap()

	h.Insert(1)
	h.Insert(-1)
	decCandidate := h.Insert(3)
	h.Insert(12)
	h.Insert(6)

	n := h.ExtractMin()
	assert.Equal(t, 4, h.n)
	assert.Equal(t, -1, n.Key)

	h.Delete(decCandidate)
	assert.Equal(t, 3, h.n)

	n = h.ExtractMin()
	assert.Equal(t, 1, n.Key)

	n = h.ExtractMin()
	assert.Equal(t, 6, n.Key)

	n = h.ExtractMin()
	assert.Equal(t, 12, n.Key)
}

func TestHeapDecreaseKey(t *testing.T) {
	h := NewHeap()

	arr := algorithms.GenerateRandomSlice(t)

	var decCandidates []*Node
	for i, v := range arr {
		if i%100 == 0 {
			decCandidates = append(decCandidates, h.Insert(v))
		} else {
			h.Insert(v)
		}
	}

	for _, node := range decCandidates {
		assert.NoError(t, h.DecreaseKey(node, node.Key-50))
	}

	result := make([]int, 0, len(arr))
	min := h.ExtractMin()
	for min != nil {
		result = append(result, min.Key)
		min = h.ExtractMin()
	}

	assert.True(t, sort.IntsAreSorted(result))
}

func TestHeapExtractMin(t *testing.T) {
	h := NewHeap()

	arr := algorithms.GenerateRandomSlice(t)

	for _, v := range arr {
		h.Insert(v)
	}

	result := make([]int, 0, len(arr))
	min := h.ExtractMin()
	for min != nil {
		result = append(result, min.Key)
		min = h.ExtractMin()
	}

	assert.True(t, sort.IntsAreSorted(result))
}
