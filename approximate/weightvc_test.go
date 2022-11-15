package approximate

import (
	"algorithms/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApproxMinWeightVC(t *testing.T) {
	var (
		a = &graph.V{Val: "a"}
		b = &graph.V{Val: "b"}
		c = &graph.V{Val: "c"}
	)

	testGraph := graph.G{
		Vertexes: []*graph.V{a, b, c},
		Adj: map[*graph.V][]*graph.V{
			a: {b},
			b: {a, c},
			c: {b},
		},
	}

	res, err := ApproxMinWeightVC(&testGraph, []float64{
		1, 3, 1,
	})
	assert.NoError(t, err)
	assert.Equal(t, []*graph.V{a, c}, res)
}
