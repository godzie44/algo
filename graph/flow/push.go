package flow

import (
	"algorithms/graph"
	"fmt"
	"math"
	"sort"
)

type (
	context struct {
		e map[*graph.V]int
		h map[*graph.V]int
		f map[graph.Edge]int
		c map[graph.Edge]int
	}
	algo struct {
		g   *graph.G
		ctx context
	}
)

func newAlgo(g *graph.G, s *graph.V) *algo {
	a := &algo{
		g: g,
		ctx: context{
			e: map[*graph.V]int{},
			h: map[*graph.V]int{},
			f: map[graph.Edge]int{},
			c: map[graph.Edge]int{},
		},
	}

	for e, w := range a.g.Weights {
		a.ctx.c[e] = w
		a.ctx.f[e] = 0
		a.ctx.f[graph.Edge{e.V2, e.V1}] = 0
	}

	for _, vert := range g.Vertexes {
		a.ctx.h[vert] = 0
		a.ctx.e[vert] = 0
	}

	for _, u := range a.g.Adj[s] {
		a.ctx.f[graph.Edge{s, u}] = a.ctx.c[graph.Edge{s, u}]
		a.ctx.f[graph.Edge{u, s}] = -a.ctx.c[graph.Edge{s, u}]
		a.ctx.e[u] = a.ctx.c[graph.Edge{s, u}]
		a.ctx.e[s] -= a.ctx.c[graph.Edge{s, u}]
	}
	a.ctx.h[s] = len(g.Vertexes)
	return a
}

func (a *algo) do(s, t *graph.V) map[graph.Edge]int {
	for {
		var u *graph.V
		for _, v := range a.g.Vertexes {
			if v == s || v == t {
				continue
			}

			if a.ctx.e[v] != 0 {
				u = v
				break
			}
		}
		if u == nil {
			break
		}

		less, _ := a.children(u)
		if len(less) == 0 {
			a.relabel(u)
		} else {
			a.push(u, less[len(less)-1])
		}
	}

	return a.ctx.f
}

func (a *algo) children(u *graph.V) (less []*graph.V, notLess []*graph.V) {
	for _, v := range a.g.Vertexes {
		if _, exists := a.ctx.f[graph.Edge{u, v}]; exists {
			if a.ctx.h[v] < a.ctx.h[u] {
				less = append(less, v)
			} else {
				notLess = append(notLess, v)
			}
		}
	}
	return less, notLess
}

func (a *algo) push(u, v *graph.V) {
	d := min(a.ctx.e[u], a.ctx.c[graph.Edge{u, v}]-a.ctx.f[graph.Edge{u, v}])
	a.ctx.f[graph.Edge{u, v}] += d
	a.ctx.f[graph.Edge{v, u}] -= d
	a.ctx.e[u] -= d
	a.ctx.e[v] += d

	if a.ctx.f[graph.Edge{u, v}] == a.ctx.c[graph.Edge{u, v}] {
		delete(a.ctx.f, graph.Edge{u, v})
	}
}

func (a *algo) relabel(u *graph.V) {
	var res = math.MaxInt

	var less, notLess = a.children(u)

	for _, v := range append(less, notLess...) {
		res = min(res, a.ctx.h[v])
	}

	a.ctx.h[u] = res + 1
}

func PushRelabelMaxFlow(g *graph.G, s, t *graph.V) map[graph.Edge]int {
	a := newAlgo(g, s)
	return a.do(s, t)
}

func (a *algo) dump() {
	for _, v := range a.g.Vertexes {
		fmt.Println("n: ", v.Val, "h:", a.ctx.h[v], "e:", a.ctx.e[v])
	}

	fmt.Printf("Edge     |  c ->   | flow -> | c <- | flow <- | \n")

	var edges []graph.Edge
	for e := range a.g.Weights {
		edges = append(edges, e)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].V1.Val.(string) > edges[j].V1.Val.(string)
	})

	for e := range a.g.Weights {
		e2 := graph.Edge{e.V2, e.V1}
		fmt.Printf("%s -> %s  |  %02d     |  %02d     |  %02d  |  %d    | \n", e.V1.Val, e.V2.Val, a.ctx.c[e], a.ctx.f[e], a.ctx.c[e2], a.ctx.f[e2])
	}
}
