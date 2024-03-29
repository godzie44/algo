package geometry

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSegmentsIntersect(t *testing.T) {
	type TestCase struct {
		p1, p2, p3, p4 Point
		intersect      bool
	}

	testCases := []TestCase{
		{
			p1:        Point{0, 0},
			p2:        Point{5, 0},
			p3:        Point{0, 1},
			p4:        Point{5, 1},
			intersect: false,
		},
		{
			p1:        Point{0, 0},
			p2:        Point{5, 0},
			p3:        Point{2, 1},
			p4:        Point{2, -1},
			intersect: true,
		},
		{
			p1:        Point{0, 0},
			p2:        Point{5, 0},
			p3:        Point{3, 0},
			p4:        Point{5, 3},
			intersect: true,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.intersect, SegmentsIntersect(tc.p1, tc.p2, tc.p3, tc.p4))
	}
}

func TestAnySegmentIntersect(t *testing.T) {
	type TestCase struct {
		segments  []Segment
		intersect bool
	}
	testCases := []TestCase{
		{
			segments: []Segment{
				NewSegment(Point{0, 0}, Point{5, 0}),
				NewSegment(Point{2, 4}, Point{5, 4}),
			},
			intersect: false,
		},
		{
			segments: []Segment{
				NewSegment(Point{0, 0}, Point{5, 0}),
				NewSegment(Point{2, 4}, Point{5, 4}),
				NewSegment(Point{4, 1}, Point{8, 6}),
			},
			intersect: false,
		},
		{
			segments: []Segment{
				NewSegment(Point{0, 0}, Point{5, 0}),
				NewSegment(Point{2, 4}, Point{5, 4}),
				NewSegment(Point{4, 1}, Point{8, 6}),
				NewSegment(Point{5.5, 1}, Point{5.5, 8}),
			},
			intersect: true,
		},
	}

	for _, tc := range testCases {
		intersect := AnySegmentIntersect(tc.segments)
		assert.Equal(t, tc.intersect, intersect)
	}
}

func TestGrahamScan(t *testing.T) {
	points := []Point{
		{1, 0},
		{8, 1},
		{7, 2},
		{9, 3},
		{6.5, 3},
		{6, 4},
		{5, 5},
		{4, 4.5},
		{3, 3.5},
		{2, 4.5},
		{1.5, 8},
		{1, 4.5},
		{0, 3},
	}

	assert.Equal(t, []Point{{1, 0}, {8, 1}, {9, 3}, {1.5, 8}, {0, 3}}, GrahamScan(points))
}

func TestJarvisScan(t *testing.T) {
	points := []Point{
		{1, 0},
		{8, 1},
		{7, 2},
		{9, 3},
		{6.5, 3},
		{6, 4},
		{5, 5},
		{4, 4.5},
		{3, 3.5},
		{2, 4.5},
		{1.5, 8},
		{1, 4.5},
		{0, 3},
	}

	assert.Equal(t, []Point{{1, 0}, {8, 1}, {9, 3}, {1.5, 8}, {0, 3}}, JarvisScan(points))
}

func TestNearPoints(t *testing.T) {
	points := []*Point{
		{1, 0},
		{3, 0},
		{6, 0},
		{2, 1},
		{5, 1},
		{7, 1},
		{0, 1.5},
		{3, 2},
		{4, 2.2},
		{5, 2.5},
		{0, 3},
		{5, 3},
		{2, 3.5},
		{6, 3.5},
		{0.5, 4},
		{4, 4},
	}

	var expectedP1, expectedP2 *Point
	var minDist = math.MaxFloat64
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			dist := points[i].distance(points[j])
			if dist < minDist {
				minDist = dist
				expectedP1, expectedP2 = points[i], points[j]
			}
		}
	}

	p1, p2 := NearPoints(points)
	assert.Equal(t, expectedP1, p1)
	assert.Equal(t, expectedP2, p2)
}
