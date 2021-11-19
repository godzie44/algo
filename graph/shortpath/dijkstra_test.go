package shortpath

import (
	"algorithms/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dijGraph = graph.G{
	Vertexes: []*graph.V{s, t, x, y, z},
	Adj: map[*graph.V][]*graph.V{
		s: {t, y},
		t: {x, y},
		x: {z},
		y: {t, z, x},
		z: {x, s},
	},
	Weights: map[graph.Edge]int{
		{s, t}: 10,
		{s, y}: 5,
		{t, x}: 1,
		{t, y}: 2,
		{x, z}: 4,
		{y, z}: 2,
		{y, x}: 9,
		{y, t}: 3,
		{z, x}: 6,
		{z, s}: 7,
	},
}

func TestDijkstra(test *testing.T) {
	path := Dijkstra(&dijGraph, s)

	var end *PathVertex

	for _, v := range path {
		if v.V == z {
			end = v
		}
	}
	zPath := []string{
		"z", "y", "s",
	}
	next := end
	for i := 0; i < len(zPath); i++ {
		assert.Equal(test, zPath[i], next.V.Val)
		next = next.P
	}

	for _, v := range path {
		if v.V == x {
			end = v
		}
	}
	xPath := []string{
		"x", "t", "y", "s",
	}
	next = end
	for i := 0; i < len(xPath); i++ {
		assert.Equal(test, xPath[i], next.V.Val)
		next = next.P
	}
}
