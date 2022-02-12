package all

import (
	"algorithms/graph"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloydWarshall(t *testing.T) {
	allPaths, pred := FloydWarshall(gr)

	assert.Equal(t, graph.WeightMatrix{
		{0, 1, -3, 2, -4},
		{3, 0, -4, 1, -1},
		{7, 4, 0, 5, 3},
		{2, -1, -5, 0, -2},
		{8, 5, 1, 6, 0},
	}, allPaths)

	assert.Equal(t, graph.MatrixG{
		{-1, 2, 3, 4, 0},
		{3, -1, 3, 1, 0},
		{3, 2, -1, 1, 0},
		{3, 2, 3, -1, 0},
		{3, 2, 3, 4, -1},
	}, pred)

	buff := bytes.NewBuffer([]byte{})
	PrintAllPairsShortestPath(buff, pred, 0, 1)
	assert.Equal(t, " 0 4 3 2 1", buff.String())
}
