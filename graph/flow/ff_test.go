package flow

import (
	"algorithms/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	s  = &graph.V{Val: "s"}
	v1 = &graph.V{Val: "v1"}
	v2 = &graph.V{Val: "v2"}
	v3 = &graph.V{Val: "v3"}
	v4 = &graph.V{Val: "v4"}
	t  = &graph.V{Val: "t"}
)

var gr = graph.G{
	Vertexes: []*graph.V{s, t, v1, v2, v3, v4},
	Adj: map[*graph.V][]*graph.V{
		s:  {v1, v2},
		v1: {v3},
		v2: {v1, v4},
		v3: {v2, t},
		v4: {v3, t},
		t:  {},
	},
	Weights: map[graph.Edge]int{
		{s, v1}:  16,
		{s, v2}:  13,
		{v1, v3}: 12,
		{v2, v1}: 4,
		{v2, v4}: 14,
		{v3, v2}: 9,
		{v3, t}:  20,
		{v4, v3}: 7,
		{v4, t}:  4,
	},
}

func TestFordFulkerson(test *testing.T) {
	result := FordFulkerson(&gr, s, t)

	assert.Equal(test, result.Weights[graph.Edge{t, v3}], 19)
	assert.Equal(test, result.Weights[graph.Edge{t, v4}], 4)
	assert.Equal(test, result.Weights[graph.Edge{v3, v1}], 12)
	assert.Equal(test, result.Weights[graph.Edge{v3, v2}], 9)
	assert.Equal(test, result.Weights[graph.Edge{v4, v2}], 11)
	assert.Equal(test, result.Weights[graph.Edge{v1, s}], 12)
	assert.Equal(test, result.Weights[graph.Edge{v2, s}], 11)
}
