package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLupDecomposition(t *testing.T) {
	a := [][]float64{
		{2, 0, 2, 0.6},
		{3, 3, 4, -2},
		{5, 5, 4, 2},
		{-1, -2, 3.4, -1},
	}

	l, u, p := lupDecomposition(a)

	expectedL := [][]float64{
		{1, 0, 0, 0},
		{0.4, 1, 0, 0},
		{-0.2, 0.5, 1, 0},
		{0.6, 0, 0.4, 1},
	}
	expectedU := [][]float64{
		{5, 5, 4, 2},
		{0, -2, 0.4, -0.2},
		{0, 0, 4, -0.5},
		{0, 0, 0, -3},
	}

	assert.Equal(t, []int{2, 0, 3, 1}, p)

	for i := 0; i < len(expectedL); i++ {
		for j := 0; j < len(expectedL); j++ {
			assert.InDelta(t, expectedL[i][j], l[i][j], 0.01)
			assert.InDelta(t, expectedU[i][j], u[i][j], 0.01)
		}
	}
}

func TestLUPSolve(t *testing.T) {
	a := [][]float64{
		{1, 2, 0},
		{3, 4, 4},
		{5, 6, 3},
	}
	b := []float64{3, 7, 8}

	x := LUPSolve(a, b)

	expectedX := []float64{-1.4, 2.2, 0.6}

	for i := 0; i < len(expectedX); i++ {
		assert.InDelta(t, expectedX[i], x[i], 0.01)
	}
}
