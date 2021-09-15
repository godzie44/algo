package search

import (
	"algorithms/graph"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestDFS(t *testing.T) {
	var (
		u = &graph.V{Val: "u"}
		v = &graph.V{Val: "v"}
		w = &graph.V{Val: "w"}
		x = &graph.V{Val: "x"}
		y = &graph.V{Val: "y"}
		z = &graph.V{Val: "z"}
	)

	g := graph.G{
		Vertexes: []*graph.V{u, v, w, x, y, z},
		Adj: map[*graph.V][]*graph.V{
			u: {v, x},
			v: {y},
			w: {z, y},
			x: {v},
			y: {x},
			z: {z},
		},
	}

	forest := DFS(g)

	assertDFSVertex(t, forest[u], 1, 8, nil, black)
	assertDFSVertex(t, forest[v], 2, 7, forest[u], black)
	assertDFSVertex(t, forest[w], 9, 12, nil, black)
	assertDFSVertex(t, forest[x], 4, 5, forest[y], black)
	assertDFSVertex(t, forest[y], 3, 6, forest[v], black)
	assertDFSVertex(t, forest[z], 10, 11, forest[w], black)
}

func assertDFSVertex(t *testing.T, v *DFSVert, d, f int, p *DFSVert, c color) {
	assert.Equal(t, d, v.D)
	assert.Equal(t, f, v.F)
	assert.Equal(t, p, v.P)
	assert.Equal(t, c, v.color)
}

func TestTopologicalSort(t *testing.T) {
	var (
		socks      = &graph.V{Val: "socks"}
		underpants = &graph.V{Val: "underpants"}
		pants      = &graph.V{Val: "pants"}
		shoes      = &graph.V{Val: "shoes"}
		belt       = &graph.V{Val: "belt"}
		shirt      = &graph.V{Val: "shirt"}
		tie        = &graph.V{Val: "tie"}
		blazer     = &graph.V{Val: "blazer"}
		watch      = &graph.V{Val: "clock"}
	)

	g := graph.G{
		Vertexes: []*graph.V{shirt, watch, underpants, pants, shoes, tie, belt, blazer, socks},
		Adj: map[*graph.V][]*graph.V{
			watch:      {},
			underpants: {pants, shoes},
			socks:      {shoes},
			pants:      {shoes, belt},
			shoes:      {},
			shirt:      {tie, belt},
			belt:       {blazer},
			tie:        {blazer},
			blazer:     {},
		},
	}

	sorted := TopologicalSort(g)

	assert.Equal(t, []*graph.V{socks, underpants, pants, shoes, watch, shirt, belt, tie, blazer}, sorted)
}

func TestStronglyConnectedComponents(t *testing.T) {
	var (
		a = &graph.V{Val: "a"}
		b = &graph.V{Val: "b"}
		c = &graph.V{Val: "c"}
		d = &graph.V{Val: "d"}
		e = &graph.V{Val: "e"}
		f = &graph.V{Val: "f"}
		g = &graph.V{Val: "g"}
		h = &graph.V{Val: "h"}
	)

	testGraph := graph.G{
		Vertexes: []*graph.V{a, b, c, d, e, f, g, h},
		Adj: map[*graph.V][]*graph.V{
			a: {b},
			b: {c, e, f},
			c: {d, g},
			d: {c, h},
			e: {a, f},
			f: {g},
			g: {f, h},
			h: {h},
		},
	}

	components := StronglyConnectedComponents(testGraph)

	sort.Slice(components, func(i, j int) bool {
		return components[i][0].Val.(string) > components[j][0].Val.(string)
	})

	assert.Equal(t, ConnectedComponents{
		{h},
		{g, f},
		{c, d},
		{a, e, b},
	}, components)
}
