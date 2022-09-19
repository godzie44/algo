package geometry

import (
	"github.com/stretchr/testify/assert"
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
