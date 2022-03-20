package flow

import (
	"algorithms/graph"
	"container/list"
	"fmt"
	"math"
	"sort"
)

type context struct {
	g *graph.G
	e map[*graph.V]int
	h map[*graph.V]int
	f map[graph.Edge]int
	c map[graph.Edge]int
}

type relabelMaxFlowAlgo struct {
	ctx context
}

func newAlgo(g *graph.G, s *graph.V) *relabelMaxFlowAlgo {
	a := &relabelMaxFlowAlgo{
		ctx: context{
			g: g,
			e: map[*graph.V]int{},
			h: map[*graph.V]int{},
			f: map[graph.Edge]int{},
			c: map[graph.Edge]int{},
		},
	}

	initPreflow(&a.ctx, g, s)
	return a
}

func (a *relabelMaxFlowAlgo) do(s, t *graph.V) map[graph.Edge]int {
	for {
		var u *graph.V
		for _, v := range a.ctx.g.Vertexes {
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

		var vert *graph.V
		for _, v := range a.ctx.g.Vertexes {
			if a.ctx.c[graph.Edge{u, v}]-a.ctx.f[graph.Edge{u, v}] > 0 && a.ctx.h[u] == a.ctx.h[v]+1 {
				vert = v
				break
			}
		}

		if vert != nil {
			push(&a.ctx, u, vert)
		} else {
			relabel(&a.ctx, u)
		}
	}

	return a.ctx.f
}

func children(ctx *context, u *graph.V) (c []*graph.V) {
	for _, v := range ctx.g.Vertexes {
		if ctx.c[graph.Edge{u, v}]-ctx.f[graph.Edge{u, v}] == 0 {
			continue
		}

		c = append(c, v)
	}
	return c
}

func push(ctx *context, u, v *graph.V) {
	d := min(ctx.e[u], ctx.c[graph.Edge{u, v}]-ctx.f[graph.Edge{u, v}])
	ctx.f[graph.Edge{u, v}] += d
	ctx.f[graph.Edge{v, u}] -= d
	ctx.e[u] -= d
	ctx.e[v] += d
}

func relabel(ctx *context, u *graph.V) {
	var res = math.MaxInt

	for _, v := range children(ctx, u) {
		res = min(res, ctx.h[v])
	}

	ctx.h[u] = res + 1
}

func PushRelabelMaxFlow(g *graph.G, s, t *graph.V) map[graph.Edge]int {
	a := newAlgo(g, s)
	return a.do(s, t)
}

func dump(ctx *context) {
	fmt.Println()

	for _, v := range ctx.g.Vertexes {
		fmt.Println("n: ", v.Val, "h:", ctx.h[v], "e:", ctx.e[v])
	}

	fmt.Printf("Edge     |  c ->   | flow -> | c <- | flow <- | \n")

	var edges []graph.Edge
	for e := range ctx.g.Weights {
		edges = append(edges, e)
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].V1.Val.(string) > edges[j].V1.Val.(string)
	})

	for e := range ctx.g.Weights {
		e2 := graph.Edge{e.V2, e.V1}
		fmt.Printf("%s -> %s  |  %02d     |  %02d     |  %02d  |  %d    | \n", e.V1.Val, e.V2.Val, ctx.c[e], ctx.f[e], ctx.c[e2], ctx.f[e2])
	}
}

type relabelToFrontAlgo struct {
	g    *graph.G
	allN map[*graph.V]*neighbours
	ctx  context
}

func RelabelToFront(g *graph.G, s, t *graph.V) map[graph.Edge]int {
	alg := newRelableToFrontAlgo(g, s, t)
	return alg.do(s, t)
}

func newRelableToFrontAlgo(g *graph.G, s, t *graph.V) *relabelToFrontAlgo {
	alg := &relabelToFrontAlgo{
		g: g,
		ctx: context{
			g: g,
			e: map[*graph.V]int{},
			h: map[*graph.V]int{},
			f: map[graph.Edge]int{},
			c: map[graph.Edge]int{},
		},
		allN: map[*graph.V]*neighbours{},
	}
	initPreflow(&alg.ctx, g, s)

	for _, v := range g.Vertexes {
		alg.allN[v] = &neighbours{}

		for _, j := range g.Vertexes {
			_, ex1 := alg.ctx.f[graph.Edge{v, j}]
			_, ex2 := alg.ctx.f[graph.Edge{j, v}]
			if ex1 || ex2 {
				alg.allN[v].n = append(alg.allN[v].n, j)
			}
		}
	}

	return alg
}

type neighbours struct {
	current int
	n       []*graph.V
}

func (a *relabelToFrontAlgo) do(s, t *graph.V) map[graph.Edge]int {
	l := list.New()

	for _, v := range a.ctx.g.Vertexes {
		if v != s && v != t {
			l.PushBack(v)
			a.allN[v].current = 0
		}
	}

	u := l.Front()

	for u != nil {
		uVert := u.Value.(*graph.V)

		oldH := a.ctx.h[uVert]
		a.discharge(uVert)
		if a.ctx.h[uVert] > oldH {
			l.MoveToFront(u)
		}
		u = u.Next()
	}
	return a.ctx.f
}

func (a *relabelToFrontAlgo) discharge(u *graph.V) {
	for a.ctx.e[u] > 0 {
		var v *graph.V
		if a.allN[u].current >= len(a.allN[u].n) {
			v = nil
		} else {
			v = a.allN[u].n[a.allN[u].current]
		}

		if v == nil {
			relabel(&a.ctx, u)
			a.allN[u].current = 0
		} else if a.ctx.c[graph.Edge{u, v}]-a.ctx.f[graph.Edge{u, v}] > 0 && a.ctx.h[u] == (a.ctx.h[v]+1) {
			push(&a.ctx, u, v)
		} else {
			a.allN[u].current += 1
		}
	}
}

func initPreflow(ctx *context, g *graph.G, s *graph.V) {
	for e, w := range g.Weights {
		ctx.c[e] = w
		ctx.f[e] = 0
		ctx.f[graph.Edge{e.V2, e.V1}] = 0
	}

	for _, vert := range g.Vertexes {
		ctx.h[vert] = 0
		ctx.e[vert] = 0
	}

	for _, u := range g.Adj[s] {
		ctx.f[graph.Edge{s, u}] = ctx.c[graph.Edge{s, u}]
		ctx.f[graph.Edge{u, s}] = -ctx.c[graph.Edge{s, u}]
		ctx.e[u] = ctx.c[graph.Edge{s, u}]
		ctx.e[s] -= ctx.c[graph.Edge{s, u}]
	}
	ctx.h[s] = len(g.Vertexes)
}
