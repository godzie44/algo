package all

import (
	"algorithms/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	a = &graph.V{Val: "a"}
	b = &graph.V{Val: "b"}
	c = &graph.V{Val: "c"}
	d = &graph.V{Val: "d"}
	e = &graph.V{Val: "e"}
)

var gr = graph.G{
	Vertexes: []*graph.V{a, b, c, d, e},
	Adj: map[*graph.V][]*graph.V{
		a: {b, e, c},
		b: {d, e},
		c: {b},
		d: {c, a},
		e: {d},
	},
	Weights: map[graph.Edge]int{
		{a, b}: 3,
		{a, c}: 8,
		{a, e}: -4,
		{b, d}: 1,
		{b, e}: 7,
		{c, b}: 4,
		{d, c}: -5,
		{d, a}: 2,
		{e, d}: 6,
	},
}

func TestSlowAllPairsShortestPaths(t *testing.T) {
	allPaths := SlowAllPairsShortestPaths(gr)

	assert.Equal(t, graph.WeightMatrix{
		{0, 1, -3, 2, -4},
		{3, 0, -4, 1, -1},
		{7, 4, 0, 5, 3},
		{2, -1, -5, 0, -2},
		{8, 5, 1, 6, -0},
	}, allPaths)
}

func TestFasterAllPairsShortestPaths(t *testing.T) {
	allPaths := FasterAllPairsShortestPaths(gr)

	assert.Equal(t, graph.WeightMatrix{
		{0, 1, -3, 2, -4},
		{3, 0, -4, 1, -1},
		{7, 4, 0, 5, 3},
		{2, -1, -5, 0, -2},
		{8, 5, 1, 6, -0},
	}, allPaths)
}
