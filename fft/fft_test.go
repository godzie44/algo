package fft

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFFT(t *testing.T) {
	a := []float64{3, 4, 1, 2.3, 12, 11, 3, 0.6}
	yRecursive := RecursiveFFT(a)
	yIter := IterativeFFT(a)

	assert.Equal(t, len(yRecursive), len(yIter))

	for i, v := range yRecursive {
		assert.InDelta(t, v, yIter[uint(i)], 0.001)
	}

	a = []float64{1, 2, 3, 4}
	yRecursive = RecursiveFFT(a)
	yIter = IterativeFFT(a)

	assert.Equal(t, len(yRecursive), len(yIter))

	for i, v := range yRecursive {
		assert.InDelta(t, v, yIter[uint(i)], 0.001)
	}
}
