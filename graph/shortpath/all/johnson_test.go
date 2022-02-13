package all

import (
	"algorithms/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJohnson(t *testing.T) {
	allPaths, err := Johnson(gr)
	assert.NoError(t, err)

	assert.Equal(t, graph.WeightMatrix{
		{0, 1, -3, 2, -4},
		{3, 0, -4, 1, -1},
		{7, 4, 0, 5, 3},
		{2, -1, -5, 0, -2},
		{8, 5, 1, 6, 0},
	}, allPaths)
}
