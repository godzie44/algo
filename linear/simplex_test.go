package linear

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimplex(t *testing.T) {
	res, err := Simplex([][]float64{
		{2, -1},
		{1, -5},
	}, []float64{2, -4}, []float64{2, -1})
	assert.NoError(t, err)
	assert.InDelta(t, 1.5555, res[0], 0.01)
	assert.InDelta(t, 1.1111, res[1], 0.01)

	res, err = Simplex([][]float64{
		{1, 1, 3},
		{2, 2, 5},
		{4, 1, 2},
	}, []float64{30, 24, 36}, []float64{3, 1, 2})
	assert.NoError(t, err)
	assert.InDelta(t, 8, res[0], 0.01)
	assert.InDelta(t, 4, res[1], 0.01)
	assert.InDelta(t, 0, res[2], 0.01)

	res, err = Simplex([][]float64{
		{-1, 1, 3},
		{2, 2, -5},
		{4, 1, 2},
	}, []float64{0, -24, 0}, []float64{3, -1, -2})
	assert.Error(t, err)

	res, err = Simplex([][]float64{
		{-1, -1, 0},
		{0, -1, -1},
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}, []float64{-1, -1, 1, 1, 1}, []float64{-1, -3, -1})
	assert.NoError(t, err)
	assert.InDelta(t, 1, res[0], 0.01)
	assert.InDelta(t, 0, res[1], 0.01)
	assert.InDelta(t, 1, res[2], 0.01)
}
