package shortpath

import (
	"algorithms/graph"
	"algorithms/priorityqueue"
)

func Dijkstra(g *graph.G, source *graph.V) Path {
	vertexes := initialize(g, source)

	s := make(Path, 0)

	queue := priorityqueue.NewPriorityQueue(priorityqueue.MinQueue)
	for _, v := range vertexes {
		queue.Insert(&priorityqueue.QueueElement{
			Key:   v.D,
			Value: v,
		})
	}

	for {
		u, _ := queue.ExtractMaxOrMin()
		s = append(s, u.Value.(*PathVertex))

		if queue.Empty() {
			return s
		}

		for _, v := range g.Adj[u.Value.(*PathVertex).V] {
			relax(g, u.Value.(*PathVertex), vertexes[v])

			var idx int
			queue.ForEach(func(i int, el *priorityqueue.QueueElement) {
				if el.Value.(*PathVertex).V == v {
					idx = i
				}
			})

			_ = queue.IncreaseKey(idx, vertexes[v].D)
		}
	}
}
