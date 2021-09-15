package search

import (
	"algorithms/graph"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBFS(test *testing.T) {
	var (
		r = &graph.V{Val: "r"}
		v = &graph.V{Val: "v"}
		s = &graph.V{Val: "s"}
		w = &graph.V{Val: "w"}
		t = &graph.V{Val: "t"}
		x = &graph.V{Val: "x"}
		u = &graph.V{Val: "u"}
		y = &graph.V{Val: "y"}
	)

	g := graph.G{
		Vertexes: []*graph.V{r, v, s, w, t, x, u, y},
		Adj: map[*graph.V][]*graph.V{
			r: {v, s},
			v: {r},
			s: {r, w},
			w: {s, t, x},
			t: {w, x, u},
			x: {w, t, u, y},
			u: {t, x, y},
			y: {x, u},
		},
	}

	result := BFS(g, s)

	assertBFSVertex(test, result[s], 0, nil, black)
	assertBFSVertex(test, result[r], 1, result[s], black)
	assertBFSVertex(test, result[v], 2, result[r], black)
	assertBFSVertex(test, result[w], 1, result[s], black)
	assertBFSVertex(test, result[t], 2, result[w], black)
	assertBFSVertex(test, result[x], 2, result[w], black)
	assertBFSVertex(test, result[u], 3, result[t], black)
	assertBFSVertex(test, result[y], 3, result[x], black)
}

func assertBFSVertex(t *testing.T, v *BFSVert, d int, p *BFSVert, c color) {
	assert.Equal(t, d, v.D)
	assert.Equal(t, p, v.P)
	assert.Equal(t, c, v.color)
}

func TestBFSPrint(test *testing.T) {
	var (
		r = &graph.V{Val: "r"}
		v = &graph.V{Val: "v"}
		s = &graph.V{Val: "s"}
		w = &graph.V{Val: "w"}
		t = &graph.V{Val: "t"}
		x = &graph.V{Val: "x"}
		u = &graph.V{Val: "u"}
		y = &graph.V{Val: "y"}
	)

	g := graph.G{
		Vertexes: []*graph.V{r, v, s, w, t, x, u, y},
		Adj: map[*graph.V][]*graph.V{
			r: {v, s},
			v: {r},
			s: {r, w},
			w: {s, t, x},
			t: {w, x, u},
			x: {w, t, u, y},
			u: {t, x, y},
			y: {x, u},
		},
	}

	result := BFS(g, s)

	buff := bytes.Buffer{}
	PrintPath(&buff, g, result[s], result[u])
	assert.Equal(test, "s\nw\nt\nu\n", buff.String())
}
