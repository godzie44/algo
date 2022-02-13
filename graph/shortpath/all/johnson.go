package all

import (
	"algorithms/graph"
	"algorithms/graph/shortpath"
	"errors"
)

func extendG(g *graph.G) (*graph.G, *graph.V) {
	extG := g.Copy()

	newVertex := &graph.V{
		Val: "s",
	}

	extG.Vertexes = append(extG.Vertexes, newVertex)
	extG.Adj[newVertex] = g.Vertexes

	for _, v := range g.Vertexes {
		newEdge := graph.Edge{
			V1: newVertex,
			V2: v,
		}
		extG.Weights[newEdge] = 0
	}

	return extG, newVertex
}

func Johnson(g graph.G) (graph.WeightMatrix, error) {
	extG, s := extendG(&g)

	path, exists := shortpath.BellmanFord(extG, s)
	if !exists {
		return nil, errors.New("cycle with negative weight found")
	}

	h := map[*graph.V]int{}
	for _, vert := range extG.Vertexes {
		h[vert] = path.LenFor(vert)
	}

	for edge := range extG.Weights {
		extG.Weights[edge] += h[edge.V1] - h[edge.V2]
	}

	tempG := g.Copy()
	tempG.Weights = extG.Weights
	d := emptyMatrix(len(tempG.Vertexes))

	for uInd, u := range tempG.Vertexes {
		path := shortpath.Dijkstra(tempG, u)

		for vInd, v := range tempG.Vertexes {
			d[uInd][vInd] = path.LenFor(v) + h[v] - h[u]
		}
	}

	return d, nil
}
