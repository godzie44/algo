package search

import (
	"algorithms/graph"
	"sort"
)

type DFSVert struct {
	V        *graph.V
	color    color
	D, F     int
	P        *DFSVert
	Children []*DFSVert
}

type DFSForest map[*graph.V]*DFSVert

func (f DFSForest) FlatTree(root *DFSVert) (res []*DFSVert) {
	res = []*DFSVert{root}
	stack := []*DFSVert{root}

	for len(stack) != 0 {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, v.Children...)
		stack = append(stack, v.Children...)
	}

	return res
}

type dfsCtx struct {
	time         int
	searchTree   DFSForest
	graph        graph.G
	vertexDoneFn func(v *graph.V)
}

func DFS(g graph.G) DFSForest {
	return dfs(g, nil)
}

func dfs(g graph.G, vertexDoneFn func(v *graph.V)) DFSForest {
	vertexes := make(DFSForest, len(g.Vertexes))
	for _, v := range g.Vertexes {
		vertexes[v] = &DFSVert{
			V:     v,
			color: white,
			P:     nil,
		}
	}

	ctx := &dfsCtx{
		time:         0,
		searchTree:   vertexes,
		graph:        g,
		vertexDoneFn: vertexDoneFn,
	}

	for _, v := range g.Vertexes {
		dfsVert := vertexes[v]
		if dfsVert.color == white {
			visit(ctx, dfsVert)
		}
	}

	return vertexes
}

func visit(ctx *dfsCtx, u *DFSVert) {
	ctx.time++
	u.D = ctx.time
	u.color = gray

	for _, child := range ctx.graph.Adj[u.V] {
		childV := ctx.searchTree[child]
		if childV.color == white {
			childV.P = u
			u.Children = append(u.Children, childV)
			visit(ctx, childV)
		}
	}

	u.color = black
	ctx.time++
	u.F = ctx.time

	if ctx.vertexDoneFn != nil {
		ctx.vertexDoneFn(u.V)
	}
}

func TopologicalSort(g graph.G) []*graph.V {
	result := make([]*graph.V, len(g.Vertexes))
	pos := len(result) - 1

	dfs(g, func(v *graph.V) {
		result[pos] = v
		pos--
	})

	return result
}

type ConnectedComponents [][]*graph.V

func StronglyConnectedComponents(g graph.G) ConnectedComponents {
	dfsForest := DFS(g)

	vertexes := make([]*DFSVert, 0, len(dfsForest))
	for _, v := range dfsForest {
		vertexes = append(vertexes, v)
	}
	sort.Slice(vertexes, func(i, j int) bool {
		return vertexes[i].F > vertexes[j].F
	})

	gT := graph.Transpose(g)
	for i, dfsV := range vertexes {
		gT.Vertexes[i] = dfsV.V
	}

	tDfsForest := DFS(gT)

	result := make(ConnectedComponents, 0)
	for _, v := range tDfsForest {
		if v.P == nil {
			tree := tDfsForest.FlatTree(v)
			var vertexTree []*graph.V
			for _, v := range tree {
				vertexTree = append(vertexTree, v.V)
			}
			result = append(result, vertexTree)
		}
	}

	return result
}
