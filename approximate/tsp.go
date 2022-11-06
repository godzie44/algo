package approximate

import (
	"algorithms/graph"
	"algorithms/graph/mst"
	"algorithms/graph/search"
	"errors"
)

func ApproxTspTour(g *graph.G) ([]*graph.V, error) {
	r := g.Vertexes[0]
	mstTree := mst.PrimMST(g, r)

	newG := graph.G{
		Vertexes: []*graph.V{},
		Adj:      map[*graph.V][]*graph.V{},
		Weights:  map[graph.Edge]int{},
	}

	for _, e := range mstTree {
		if newG.Adj[e.V1] == nil {
			newG.Vertexes = append(newG.Vertexes, e.V1)
			newG.Adj[e.V1] = []*graph.V{}
		}
		if newG.Adj[e.V2] == nil {
			newG.Vertexes = append(newG.Vertexes, e.V2)
			newG.Adj[e.V2] = []*graph.V{}
		}
		newG.Adj[e.V1] = append(newG.Adj[e.V1], e.V2)
		newG.Weights[e] = g.Weight(e.V1, e.V2)
	}

	dfsTree := search.DFS(newG)

	for _, v := range dfsTree {
		if v.V == mstTree[0].V1 {
			ft := dfsTree.FlatTree(v)
			vertexes := make([]*graph.V, len(ft))
			for i := range ft {
				vertexes[i] = ft[i].V
			}
			return vertexes, nil
		}
	}

	return nil, errors.New("fail to resolve TSP tour")
}
