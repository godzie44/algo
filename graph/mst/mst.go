package mst

import (
	"algorithms/disjointset"
	"algorithms/graph"
	"algorithms/priorityqueue"
	"math"
	"sort"
)

type SpanningTree []graph.Edge

func KruskalMST(g *graph.G) SpanningTree {
	setMap := make(map[*graph.V]*disjointset.Set)
	for _, v := range g.Vertexes {
		setMap[v] = disjointset.NewSet(v)
	}

	type weightedEdge struct {
		w int
		e graph.Edge
	}
	weights := make([]weightedEdge, 0, len(g.Weights))
	for e, w := range g.Weights {
		weights = append(weights, weightedEdge{
			w: w,
			e: e,
		})
	}

	sort.Slice(weights, func(i, j int) bool {
		return weights[i].w < weights[j].w
	})

	var st SpanningTree
	for _, we := range weights {
		if disjointset.FindSet(setMap[we.e.V1]) != disjointset.FindSet(setMap[we.e.V2]) {
			st = append(st, we.e)
			disjointset.Union(setMap[we.e.V1], setMap[we.e.V2])
		}
	}

	return st
}

func PrimMST(g *graph.G, r *graph.V) SpanningTree {
	type edge struct {
		parent *graph.V
		vert   *graph.V
	}

	queue := priorityqueue.NewPriorityQueue(priorityqueue.MinQueue)
	for _, v := range g.Vertexes {
		key := math.MaxInt
		if v == r {
			key = 0
		}
		queue.Insert(&priorityqueue.QueueElement{
			Key: key,
			Value: &edge{
				parent: nil,
				vert:   v,
			},
		})
	}

	var result SpanningTree
	for {
		u, err := queue.ExtractMaxOrMin()
		if err == priorityqueue.ErrEmptyQueue {
			break
		}

		result = append(result, graph.Edge{V1: u.Value.(*edge).parent, V2: u.Value.(*edge).vert})

		for _, v := range g.Adj[u.Value.(*edge).vert] {
			var vEl *priorityqueue.QueueElement
			var idx int
			queue.ForEach(func(i int, el *priorityqueue.QueueElement) {
				if el.Value.(*edge).vert == v {
					vEl = el
					idx = i
				}
			})

			if vEl != nil && g.Weight(u.Value.(*edge).vert, v) < vEl.Key {
				vEl.Value.(*edge).parent = u.Value.(*edge).vert
				_ = queue.IncreaseKey(idx, g.Weight(u.Value.(*edge).vert, v))
			}
		}
	}

	return result[1:]
}
