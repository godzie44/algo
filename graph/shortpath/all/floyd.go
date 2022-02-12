package all

import (
	"algorithms/graph"
	"fmt"
	"io"
	"math"
)

func emptyMatrix(n int) [][]int {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}
	return g
}

func p0(w graph.WeightMatrix) graph.MatrixG {
	p := make(graph.MatrixG, len(w))
	for i := 0; i < len(w); i++ {
		p[i] = make([]int, len(w))

		for j := 0; j < len(w); j++ {
			if i == j || w[i][j] == math.MaxInt {
				p[i][j] = -1
			} else {
				p[i][j] = i
			}
		}
	}

	return p
}

func FloydWarshall(g graph.G) (graph.WeightMatrix, graph.MatrixG) {
	sumNoOverflow := func(a, b int) int {
		if a == math.MaxInt || b == math.MaxInt {
			return math.MaxInt
		}

		return a + b
	}

	w := g.WeightMatrix()
	n := len(w)
	dList := make([]graph.WeightMatrix, n+1)
	pList := make([]graph.MatrixG, n+1)
	dList[0] = w
	pList[0] = p0(w)

	for k := 1; k <= n; k++ {
		dList[k] = emptyMatrix(n)
		pList[k] = emptyMatrix(n)

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				var candidatePathLen = sumNoOverflow(dList[k-1][i][k-1], dList[k-1][k-1][j])
				if dList[k-1][i][j] <= candidatePathLen {
					dList[k][i][j] = dList[k-1][i][j]
					pList[k][i][j] = pList[k-1][i][j]
				} else {
					dList[k][i][j] = candidatePathLen
					pList[k][i][j] = pList[k-1][k-1][j]
				}
			}
		}
	}

	return dList[n], pList[n]
}

func PrintAllPairsShortestPath(w io.Writer, p graph.MatrixG, i, j int) {
	if i == j {
		fmt.Fprintf(w, " %d", i)
	} else if p[i][j] == math.MaxInt {
		fmt.Fprintf(w, "path from %d to %d not found", i, j)
	} else {
		PrintAllPairsShortestPath(w, p, i, p[i][j])
		fmt.Fprintf(w, " %d", j)
	}
}
