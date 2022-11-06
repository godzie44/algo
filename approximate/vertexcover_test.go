package approximate

import (
	"algorithms/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApproxVertexCover(t *testing.T) {
	var (
		a = &graph.V{Val: "a"}
		b = &graph.V{Val: "b"}
		c = &graph.V{Val: "c"}
		d = &graph.V{Val: "d"}
		e = &graph.V{Val: "e"}
		f = &graph.V{Val: "f"}
		g = &graph.V{Val: "g"}
	)

	testGraph := graph.G{
		Vertexes: []*graph.V{c, a, b, d, e, f, g},
		Adj: map[*graph.V][]*graph.V{
			a: {b},
			b: {a, c},
			c: {b, d, e},
			d: {c, g, e, f},
			e: {c, d, f},
			f: {d, e},
			g: {d},
		},
	}

	assert.Equal(t, []*graph.V{c, b, d, g, e, f}, ApproxVertexCover(&testGraph))
}
