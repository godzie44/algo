package flow

import (
	"algorithms/graph"
	"algorithms/graph/search"
	"math"
)

func FordFulkerson(g *graph.G, s, t *graph.V) *graph.G {
	g = g.Copy()

	for {
		sp := search.BFS(*g, s)

		var end = sp[t]
		if end.P == nil {
			return g
		}

		var cp = math.MaxInt
		for end.P != nil {
			cp = min(cp, g.Weight(end.P.V, end.V))
			end = end.P
		}

		end = sp[t]
		for end.P != nil {
			u, v := end.P.V, end.V
			uv := graph.Edge{V1: u, V2: v}
			vu := graph.Edge{V1: v, V2: u}

			g.Weights[uv] -= cp
			if g.Weights[uv] == 0 {
				g.RemoveEdge(uv)
			}

			if _, exists := g.Weights[vu]; !exists {
				g.AddEdge(vu)
			}
			g.Weights[vu] += cp

			end = end.P
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
