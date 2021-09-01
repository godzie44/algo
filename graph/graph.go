package graph

type V struct {
	Val interface{}
}

type G struct {
	Vertexes []*V

	Adj map[*V][]*V
}
