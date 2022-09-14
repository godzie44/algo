package geometry

import "math"

type Vector struct {
	x, y float64
}

func (v *Vector) mul(other *Vector) float64 {
	return v.x*other.y - other.x*v.y
}

type Point struct {
	x, y float64
}

func (p *Point) sub(other Point) Point {
	return Point{x: p.x - other.x, y: p.y - other.y}
}

func zero(f float64) bool {
	return f < 0.000001 && f > -0.000001
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
	v1 := Vector(pk.sub(pi))
	v2 := Vector(pj.sub(pi))
	return v1.mul(&v2)
}
