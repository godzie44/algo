package approximate

import (
	"algorithms/graph"
)

type EdgeList struct {
	data       map[graph.Edge]struct{}
	incidental map[*graph.V]struct{}
}

func newEdgeList() *EdgeList {
	return &EdgeList{data: map[graph.Edge]struct{}{}, incidental: map[*graph.V]struct{}{}}
}

func (e *EdgeList) add(v, u *graph.V) {
	e1 := graph.Edge{v, u}
	e2 := graph.Edge{u, v}

	_, e1Exists := e.data[e1]
	_, e2Exists := e.data[e2]

	if e1Exists || e2Exists {
		return
	}

	e.data[e1] = struct{}{}
}

func (e *EdgeList) removeIncidental(v *graph.V) {
	e.incidental[v] = struct{}{}
}
func (e *EdgeList) has(v *graph.V, u *graph.V) bool {
	_, exists := e.data[graph.Edge{v, u}]
	if !exists {
		return false
	}

	if _, exists := e.incidental[v]; exists {
		return false
	}
	if _, exists := e.incidental[u]; exists {
		return false
	}

	return true
}

func ApproxVertexCover(g *graph.G) []*graph.V {
	c := make([]*graph.V, 0)

	edges := newEdgeList()
	for _, u := range g.Vertexes {
		for _, v := range g.Adj[u] {
			edges.add(u, v)
		}
	}

	for _, u := range g.Vertexes {
		for _, v := range g.Adj[u] {
			if !edges.has(u, v) {
				continue
			}

			c = append(c, u)
			c = append(c, v)

			edges.removeIncidental(v)
			edges.removeIncidental(u)
		}
	}

	return c
}
