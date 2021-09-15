package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSimplePath(test *testing.T) {
	var (
		m = &V{Val: "m"}
		n = &V{Val: "n"}
		o = &V{Val: "o"}
		p = &V{Val: "p"}
		q = &V{Val: "q"}
		r = &V{Val: "r"}
		s = &V{Val: "s"}
		t = &V{Val: "t"}
		u = &V{Val: "u"}
		v = &V{Val: "v"}
		w = &V{Val: "w"}
		x = &V{Val: "x"}
		y = &V{Val: "y"}
		z = &V{Val: "z"}
	)

	graph := G{
		Vertexes: []*V{m, n, o, p, q, r, s, t, u, v, w, x, y, z},
		Adj: map[*V][]*V{
			m: {q, r, x},
			n: {o, u, q},
			o: {r, s, v},
			p: {o, s, z},
			q: {t},
			r: {u, y},
			s: {r},
			t: {},
			u: {t},
			v: {w, x},
			w: {z},
			x: {},
			y: {v},
			z: {},
		},
	}

	assert.Equal(test, 4, CountSimplePaths(graph, p, v))
}
