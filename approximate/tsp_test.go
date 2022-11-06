package approximate

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
		a: {b, c, d, e},
		b: {a, c, d, e},
		c: {a, b, d, e},
		d: {a, b, c, e},
		e: {a, b, c, d},
	},
	Weights: map[graph.Edge]int{
		{a, b}: 2,
		{a, c}: 3,
		{a, d}: 2,
		{a, e}: 3,
		{b, a}: 2,
		{b, c}: 1,
		{b, d}: 2,
		{b, e}: 3,
		{c, a}: 3,
		{c, b}: 1,
		{c, d}: 4,
		{c, e}: 4,
		{d, a}: 2,
		{d, b}: 2,
		{d, c}: 4,
		{d, e}: 1,
		{e, a}: 3,
		{e, b}: 3,
		{e, c}: 4,
		{e, d}: 1,
	},
}

func TestApproxTspTour(t *testing.T) {
	result, err := ApproxTspTour(&gr)
	assert.NoError(t, err)

	assert.Equal(t, []*graph.V{
		a, d, b, c, e,
	}, result)
}
