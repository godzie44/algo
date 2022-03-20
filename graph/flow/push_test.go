package flow

import (
	"algorithms/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushRelabelMaxFlow(t *testing.T) {
	testAlg(t, PushRelabelMaxFlow)
}

func TestRelabelToFront(t *testing.T) {
	testAlg(t, RelabelToFront)
}

func testAlg(t *testing.T, algo func(g *graph.G, s *graph.V, t *graph.V) map[graph.Edge]int) {
	{
		var (
			source = &graph.V{Val: "source"}
			a      = &graph.V{Val: "a"}
			b      = &graph.V{Val: "b"}
			c      = &graph.V{Val: "c"}
			d      = &graph.V{Val: "d"}
			sink   = &graph.V{Val: "sink"}
		)

		var gr = graph.G{
			Vertexes: []*graph.V{source, b, a, c, d, sink},
			Adj: map[*graph.V][]*graph.V{
				source: {a, b},
				a:      {c, sink},
				b:      {a},
				c:      {d},
				d:      {sink},
				sink:   {},
			},
			Weights: map[graph.Edge]int{
				{source, a}: 13,
				{source, b}: 10,
				{a, c}:      6,
				{a, sink}:   7,
				{b, a}:      3,
				{c, d}:      10,
				{d, sink}:   5,
			},
		}

		result := algo(&gr, source, sink)

		assert.Equal(t, result[graph.Edge{sink, d}], -5)
		assert.Equal(t, result[graph.Edge{sink, a}], -7)
		assert.Equal(t, result[graph.Edge{d, c}], -5)
		assert.Equal(t, result[graph.Edge{c, a}], -5)
		assert.Equal(t, result[graph.Edge{a, b}], -3)
		assert.Equal(t, result[graph.Edge{a, source}], -9)
		assert.Equal(t, result[graph.Edge{b, source}], -3)
	}

	{
		var (
			source = &graph.V{Val: "source"}
			a      = &graph.V{Val: "a"}
			b      = &graph.V{Val: "b"}
			c      = &graph.V{Val: "c"}
			d      = &graph.V{Val: "d"}
			sink   = &graph.V{Val: "sink"}
		)

		var gr = graph.G{
			Vertexes: []*graph.V{source, sink, a, b, c, d},
			Adj: map[*graph.V][]*graph.V{
				source: {a, b},
				a:      {c},
				b:      {a, d},
				c:      {b, sink},
				d:      {c, sink},
				sink:   {},
			},
			Weights: map[graph.Edge]int{
				{source, a}: 16,
				{source, b}: 13,
				{a, c}:      12,
				{b, a}:      4,
				{b, d}:      14,
				{c, b}:      9,
				{c, sink}:   20,
				{d, c}:      7,
				{d, sink}:   4,
			},
		}

		result := algo(&gr, source, sink)

		assert.Equal(t, result[graph.Edge{sink, c}], -19)
		assert.Equal(t, result[graph.Edge{sink, d}], -4)
		assert.Equal(t, result[graph.Edge{c, d}], -7)
		assert.Equal(t, result[graph.Edge{c, a}], -12)
		assert.Equal(t, result[graph.Edge{d, b}], -11)
		assert.Equal(t, result[graph.Edge{a, source}], -12)
		assert.Equal(t, result[graph.Edge{b, source}], -11)
	}

	{
		var (
			source = &graph.V{Val: "source"}
			a      = &graph.V{Val: "a"}
			b      = &graph.V{Val: "y"}
			c      = &graph.V{Val: "c"}
			sink   = &graph.V{Val: "sink"}
		)

		var gr = graph.G{
			Vertexes: []*graph.V{source, a, b, c, sink},
			Adj: map[*graph.V][]*graph.V{
				source: {a, b},
				a:      {b, sink},
				b:      {c},
				c:      {sink, a},
				sink:   {},
			},
			Weights: map[graph.Edge]int{
				{source, a}: 12,
				{source, b}: 14,
				{a, b}:      5,
				{a, sink}:   16,
				{b, c}:      8,
				{c, sink}:   10,
				{c, a}:      7,
			},
		}

		result := algo(&gr, source, sink)
		assert.Equal(t, result[graph.Edge{a, source}], -12)
		assert.Equal(t, result[graph.Edge{b, source}], -8)
		assert.Equal(t, result[graph.Edge{c, b}], -8)
		assert.Equal(t, result[graph.Edge{sink, c}], -8)
		assert.Equal(t, result[graph.Edge{sink, a}], -12)

	}
}
