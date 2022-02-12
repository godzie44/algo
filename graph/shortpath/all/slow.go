package all

import (
	"algorithms/graph"
	"math"
)

func SlowAllPairsShortestPaths(g graph.G) graph.WeightMatrix {
	w := g.WeightMatrix()
	n := len(g.Vertexes)
	lMatrices := make([]graph.WeightMatrix, n-1)
	lMatrices[0] = w

	for i := 1; i < n-1; i++ {
		lMatrices[i] = extendShortestPath(lMatrices[i-1], w)
	}

	return lMatrices[n-2]
}

func FasterAllPairsShortestPaths(g graph.G) graph.WeightMatrix {
	w := g.WeightMatrix()
	n := len(w)
	lMatrices := make([]graph.WeightMatrix, (n-1)*2)
	lMatrices[0] = w
	m := 0

	for m < n-1 {
		var nextM int
		if m == 0 {
			nextM = 2
		} else {
			nextM = 2 * m
		}

		lMatrices[nextM] = extendShortestPath(lMatrices[m], lMatrices[m])
		m = nextM
	}

	return lMatrices[m]
}

func extendShortestPath(l, w graph.WeightMatrix) graph.WeightMatrix {
	n := len(l)
	res := make(graph.WeightMatrix, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	sumNoOverflow := func(a, b int) int {
		if a == math.MaxInt || b == math.MaxInt {
			return math.MaxInt
		}

		return a + b
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = math.MaxInt
			for k := 0; k < n; k++ {
				res[i][j] = min(res[i][j], sumNoOverflow(l[i][k], w[k][j]))
			}
		}
	}

	return res
}
