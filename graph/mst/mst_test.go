package mst

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
	f = &graph.V{Val: "f"}
	g = &graph.V{Val: "g"}
	h = &graph.V{Val: "h"}
	i = &graph.V{Val: "i"}
)

var gr = graph.G{
	Vertexes: []*graph.V{a, b, c, d, e, f, g, h, i},
	Adj: map[*graph.V][]*graph.V{
		a: {b, h},
		b: {a, c, h},
		c: {b, d, i, f},
		d: {c, e, f},
		e: {d, f},
		f: {e, d, c, g},
		g: {f, i, h},
		h: {a, i, g},
		i: {c, h, g},
	},
	Weights: map[graph.Edge]int{
		{h, i}: 7,
		{g, h}: 1,
		{g, i}: 6,
		{f, g}: 2,
		{e, f}: 10,
		{d, f}: 14,
		{d, e}: 9,
		{c, i}: 2,
		{c, f}: 4,
		{c, d}: 7,
		{b, h}: 11,
		{b, c}: 8,
		{a, h}: 8,
		{a, b}: 4,
	},
}

func TestKruskalMST(t *testing.T) {
	mst := KruskalMST(&gr)

	wSum := 0
	for _, e := range mst {
		wSum += gr.Weights[e]
	}

	assert.Equal(t, 37, wSum)
	assert.Len(t, mst, 8)
}

func TestPrimMST(t *testing.T) {
	mst := PrimMST(&gr, a)

	wSum := 0
	for _, e := range mst {
		wSum += gr.Weights[e]
	}

	assert.Equal(t, 37, wSum)
	assert.Len(t, mst, 8)
}
