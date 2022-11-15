package approximate

import (
	"algorithms/graph"
	"algorithms/linear"
)

func ApproxMinWeightVC(g *graph.G, weights []float64) ([]*graph.V, error) {
	var a = linear.Matrix{}
	var b = linear.Vector{}
	var c = linear.Vector{}

	vertexMap := map[*graph.V]int{}
	for vIdx, v := range g.Vertexes {
		vertexMap[v] = vIdx
	}
	seenEdges := map[graph.Edge]struct{}{}

	for vIdx := range g.Vertexes {
		c = append(c, -1*weights[vIdx])
	}

	for u, vertexes := range g.Adj {
		for _, v := range vertexes {
			if _, exists := seenEdges[graph.Edge{u, v}]; exists {
				continue
			}
			if _, exists := seenEdges[graph.Edge{v, u}]; exists {
				continue
			}
			seenEdges[graph.Edge{u, v}] = struct{}{}
			a = append(a, makeCond2Var(len(g.Vertexes), vertexMap[u], vertexMap[v]))
			b = append(b, -1)
		}
	}

	for vIdx := range g.Vertexes {
		a = append(a, makeCond1Var(len(g.Vertexes), vIdx))
		b = append(b, 1)
	}

	linearRes, err := linear.Simplex(a, b, c)
	if err != nil {
		return nil, err
	}

	vertexCover := []*graph.V{}
	for idx, res := range linearRes {
		if res >= 0.5 {
			vertexCover = append(vertexCover, g.Vertexes[idx])
		}
	}
	return vertexCover, nil
}

func makeCond1Var(len int, idx int) linear.Vector {
	vec := make(linear.Vector, len)
	vec[idx] = 1
	return vec
}

func makeCond2Var(len int, idx1, idx2 int) linear.Vector {
	vec := make(linear.Vector, len)
	vec[idx1] = -1
	vec[idx2] = -1
	return vec
}
