package shortpath

import (
	"algorithms/graph"
	"math"
)

type PathVertex struct {
	V *graph.V
	D int
	P *PathVertex
}

func mapIntoPath(m map[*graph.V]*PathVertex) Path {
	path := make([]*PathVertex, 0, len(m))
	for _, v := range m {
		path = append(path, v)
	}
	return path
}

type Path []*PathVertex

func (p Path) Source() *PathVertex {
	for _, v := range p {
		if v.D == 0 {
			return v
		}
	}
	return nil
}

func initialize(g *graph.G, source *graph.V) map[*graph.V]*PathVertex {
	result := make(map[*graph.V]*PathVertex, len(g.Vertexes))

	for _, v := range g.Vertexes {
		result[v] = &PathVertex{
			V: v,
			D: math.MaxInt,
			P: nil,
		}
	}
	result[source].D = 0

	return result
}

func sumIfInf(a, b int) int {
	if a == math.MaxInt || b == math.MaxInt {
		return math.MaxInt
	}
	return a + b
}

func relax(g *graph.G, u, v *PathVertex) {
	if v.D > sumIfInf(u.D, g.Weight(u.V, v.V)) {
		v.D = u.D + g.Weight(u.V, v.V)
		v.P = u
	}
}
