package search

import (
	"algorithms/graph"
	"fmt"
	"io"
	"math"
)

type color int

const (
	_ color = iota
	white
	gray
	black
)

type BFSVert struct {
	V     *graph.V
	color color
	D     int
	P     *BFSVert
}

type BFSTree map[*graph.V]*BFSVert

func BFS(g graph.G, s *graph.V) BFSTree {
	vertexes := make(map[*graph.V]*BFSVert, len(g.Vertexes))
	for _, v := range g.Vertexes {
		vertexes[v] = &BFSVert{
			V:     v,
			color: white,
			D:     math.MaxInt,
			P:     nil,
		}
	}

	vertexes[s] = &BFSVert{
		V:     s,
		color: gray,
		D:     0,
		P:     nil,
	}

	var queue = []*BFSVert{vertexes[s]}
	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]

		for i := range g.Adj[u.V] {
			v := vertexes[g.Adj[u.V][i]]
			if v.color == white {
				v.color = gray
				v.D = u.D + 1
				v.P = u
				queue = append(queue, v)
			}
		}
		u.color = black
	}

	return vertexes
}

func PrintPath(w io.Writer, g graph.G, s, v *BFSVert) {
	if v == s {
		fmt.Fprintln(w, s.V.Val)
	} else if v.P == nil {
		fmt.Println("path not found")
	} else {
		PrintPath(w, g, s, v.P)
		fmt.Fprintln(w, v.V.Val)
	}
}
