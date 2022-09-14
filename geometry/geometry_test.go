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
