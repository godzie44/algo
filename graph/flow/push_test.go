package flow

import (
	"algorithms/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	source = &graph.V{Val: "source"}
	a      = &graph.V{Val: "a"}
	b      = &graph.V{Val: "b"}
	cc     = &graph.V{Val: "c"}
	d      = &graph.V{Val: "d"}
	sink   = &graph.V{Val: "sink"}
)

var gr2 = graph.G{
	Vertexes: []*graph.V{source, b, a, cc, d, sink},
	Adj: map[*graph.V][]*graph.V{
		source: {a, b},
		a:      {cc, sink},
		b:      {a},
		cc:     {d},
		d:      {sink},
		sink:   {},
	},
	Weights: map[graph.Edge]int{
		{source, a}: 13,
		{source, b}: 10,
		{a, cc}:     6,
		{a, sink}:   7,
		{b, a}:      3,
		{cc, d}:     10,
		{d, sink}:   5,
	},
}

func TestPushRelabelMaxFlow(test *testing.T) {
	result := PushRelabelMaxFlow(&gr, s, t)

	assert.Equal(test, result[graph.Edge{t, v3}], -19)
	assert.Equal(test, result[graph.Edge{t, v4}], -4)
	assert.Equal(test, result[graph.Edge{v3, v4}], -7)
	assert.Equal(test, result[graph.Edge{v3, v1}], -12)
	assert.Equal(test, result[graph.Edge{v4, v2}], -11)
	assert.Equal(test, result[graph.Edge{v1, s}], -12)
	assert.Equal(test, result[graph.Edge{v2, s}], -11)
}

func TestPushRelabelMaxFlow2(test *testing.T) {
	result := PushRelabelMaxFlow(&gr2, source, sink)

	assert.Equal(test, result[graph.Edge{sink, d}], -5)
	assert.Equal(test, result[graph.Edge{sink, a}], -7)
	assert.Equal(test, result[graph.Edge{d, cc}], -5)
	assert.Equal(test, result[graph.Edge{cc, a}], -5)
	assert.Equal(test, result[graph.Edge{a, b}], -3)
	assert.Equal(test, result[graph.Edge{a, source}], -9)
	assert.Equal(test, result[graph.Edge{b, source}], -3)
}
