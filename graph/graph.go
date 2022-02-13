package graph

import "math"

type V struct {
	Val interface{}
}

type Edge struct {
	V1, V2 *V
}

type G struct {
	Vertexes []*V

	Adj map[*V][]*V

	Weights map[Edge]int
}

type WeightMatrix [][]int

func (g *G) Copy() *G {
	copyG := &G{
		Adj:     map[*V][]*V{},
		Weights: map[Edge]int{},
	}

	copyG.Vertexes = append(copyG.Vertexes, g.Vertexes...)

	for k, v := range g.Adj {
		copyG.Adj[k] = v
	}
	for k, v := range g.Weights {
		copyG.Weights[k] = v
	}

	return copyG
}

func (g *G) WeightMatrix() (w WeightMatrix) {
	n := len(g.Vertexes)
	w = make(WeightMatrix, n)
	for i := 0; i < n; i++ {
		w[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if we, exists := g.Weights[Edge{g.Vertexes[i], g.Vertexes[j]}]; exists {
				w[i][j] = we
			} else {
				w[i][j] = math.MaxInt
			}
		}
	}
	return w
}

func (g *G) Weight(v1, v2 *V) int {
	if w, exists := g.Weights[Edge{v1, v2}]; exists {
		return w
	}

	return g.Weights[Edge{v2, v1}]
}

func Transpose(g G) G {
	transposed := G{Vertexes: g.Vertexes, Adj: map[*V][]*V{}}

	for v := range g.Adj {
		for _, u := range g.Adj[v] {
			transposed.Adj[u] = append(transposed.Adj[u], v)
		}
	}

	return transposed
}

func count(d map[*V]int, w map[*V]bool, g G, v *V) int {
	if w[v] {
		return d[v]
	} else {
		sum := 0
		w[v] = true

		for _, c := range g.Adj[v] {
			sum += count(d, w, g, c)
		}
		d[v] = sum
		return sum
	}
}

func CountSimplePaths(g G, s, t *V) int {
	var d = map[*V]int{}
	var w = map[*V]bool{}

	d[s] = 1
	w[s] = true

	answer := count(d, w, Transpose(g), t)
	return answer
}

type MatrixG [][]int
