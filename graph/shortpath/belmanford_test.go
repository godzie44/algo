package shortpath

import (
	"algorithms/graph"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	s = &graph.V{Val: "s"}
	t = &graph.V{Val: "t"}
	x = &graph.V{Val: "x"}
	y = &graph.V{Val: "y"}
	z = &graph.V{Val: "z"}
)

var bfGraph = graph.G{
	Vertexes: []*graph.V{s, t, x, y, z},
	Adj: map[*graph.V][]*graph.V{
		s: {t, y},
		t: {x, z, y},
		x: {t},
		y: {z, x},
		z: {x, s},
	},
	Weights: map[graph.Edge]int{
		{s, t}: 6,
		{s, y}: 7,
		{t, x}: 5,
		{t, z}: -4,
		{t, y}: 8,
		{x, t}: -2,
		{y, z}: 9,
		{y, x}: -3,
		{z, x}: 7,
		{z, s}: 2,
	},
}

func TestBellmanFord(test *testing.T) {
	path, exists := BellmanFord(&bfGraph, s)
	require.True(test, exists)

	var end *PathVertex
	for _, v := range path {
		if v.V == z {
			end = v
		}
	}

	expectedPath := []string{
		"z", "t", "x", "y", "s",
	}

	next := end
	for i := 0; i < 5; i++ {
		assert.Equal(test, expectedPath[i], next.V.Val)
		next = next.P
	}
}
