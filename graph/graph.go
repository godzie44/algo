package graph

type V struct {
	Val interface{}
}

type G struct {
	Vertexes []*V

	Adj map[*V][]*V
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
