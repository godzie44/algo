package shortpath

import "algorithms/graph"

func BellmanFord(g *graph.G, source *graph.V) (Path, bool) {
	vertexes := initialize(g, source)
	for i := 1; i < len(g.Vertexes); i++ {
		for u, adj := range g.Adj {
			for _, v := range adj {

				relax(g, vertexes[u], vertexes[v])
			}
		}
	}

	for u, adj := range g.Adj {
		for _, v := range adj {
			if vertexes[v].D > vertexes[u].D+g.Weight(u, v) {
				return nil, false
			}
		}
	}

	return mapIntoPath(vertexes), true
}
