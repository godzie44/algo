package geometry

import (
	"algorithms/tree/redblack"
	"math"
	"sort"
)

const EPSILON = 1e-9

type Point struct {
	x, y float64
}

func (p *Point) vectorMul(other *Point) float64 {
	return p.x*other.y - other.x*p.y
}

func (p *Point) sub(other Point) Point {
	return Point{x: p.x - other.x, y: p.y - other.y}
}

func zero(f float64) bool {
	return math.Abs(f) < EPSILON
}

func SegmentsIntersect(p1, p2, p3, p4 Point) bool {
	d1 := direction(p3, p4, p1)
	d2 := direction(p3, p4, p2)
	d3 := direction(p1, p2, p3)
	d4 := direction(p1, p2, p4)
	if ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) && ((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0)) {
		return true
	} else if zero(d1) && onSegment(p3, p4, p1) {
		return true
	} else if zero(d2) && onSegment(p3, p4, p2) {
		return true
	} else if zero(d3) && onSegment(p1, p2, p3) {
		return true
	} else if zero(d4) && onSegment(p1, p2, p4) {
		return true
	}
	return false
}

func onSegment(pi, pj, pk Point) bool {
	if math.Min(pi.x, pj.x) <= pk.x && pk.x <= math.Max(pi.x, pj.x) &&
		math.Min(pi.y, pj.y) <= pk.y && pk.y <= math.Max(pi.y, pj.y) {
		return true
	}
	return false
}

func direction(pi, pj, pk Point) float64 {
	v1 := pk.sub(pi)
	v2 := pj.sub(pi)
	return v1.vectorMul(&v2)
}

type SegmentPoint struct {
	Point
	segment *Segment
}

type Segment struct {
	p, q SegmentPoint
}

func NewSegment(p, q Point) Segment {
	s := Segment{
		p: SegmentPoint{p, nil},
		q: SegmentPoint{q, nil},
	}
	s.p.segment = &s
	s.q.segment = &s
	return s
}

func (s *Segment) GetY(x float64) float64 {
	if zero(s.p.x - s.q.x) {
		return s.p.y
	}

	return s.p.y + (s.q.y-s.p.y)*(x-s.p.x)/(s.q.x-s.p.x)
}

func (s *Segment) Compare(candidate *Segment) int {
	x := math.Max(math.Min(s.p.x, s.q.x), math.Min(candidate.p.x, candidate.q.x))
	s1 := s.GetY(x)
	s2 := candidate.GetY(x)

	if zero(s1 - s2) {
		return 0
	}

	if s1 < s2 {
		return -1
	}
	return 1
}

func AnySegmentIntersect(segments []Segment) bool {
	points := make([]SegmentPoint, len(segments)*2)

	for i, seg := range segments {
		points[i*2] = seg.p
		points[i*2+1] = seg.q
	}

	sort.Slice(points, func(i, j int) bool {
		if zero(points[i].x - points[j].x) {
			return points[i].y < points[j].y
		}
		return points[i].x < points[j].x
	})

	set := redblack.NewTree[*Segment]()
	nodes := make(map[*Segment]*redblack.Node[*Segment])

	for _, p := range points {
		if p.segment.p == p {
			node := set.Add(p.segment)
			nodes[p.segment] = node
			if existsAndIntersect(set, p.segment, set.Predecessor(node)) || existsAndIntersect(set, p.segment, set.Successor(node)) {
				return true
			}
		}

		if p.segment.q == p {
			above := set.Successor(nodes[p.segment])
			below := set.Predecessor(nodes[p.segment])

			if exists(set, above) && exists(set, below) && SegmentsIntersect(above.Val.p.Point, above.Val.q.Point, below.Val.p.Point, below.Val.q.Point) {
				return true
			}
			set.Delete(nodes[p.segment])
		}
	}

	return false
}

func exists(set *redblack.Tree[*Segment], node *redblack.Node[*Segment]) bool {
	return node != nil && node != set.Nil
}

func existsAndIntersect(set *redblack.Tree[*Segment], s *Segment, node *redblack.Node[*Segment]) bool {
	if !exists(set, node) {
		return false
	}
	s2 := node.Val

	return SegmentsIntersect(s.p.Point, s.q.Point, s2.p.Point, s2.q.Point)
}
