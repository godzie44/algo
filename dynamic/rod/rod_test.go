package rod

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	p         costList = []float64{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	testCases          = []struct {
		n         int
		q         float64
		splitting []int
	}{
		{0, 0, []int{}},
		{1, 1, []int{1}},
		{2, 5, []int{2}},
		{3, 8, []int{3}},
		{4, 10, []int{2, 2}},
		{5, 13, []int{2, 3}},
		{6, 17, []int{6}},
		{7, 18, []int{1, 6}},
		{8, 22, []int{2, 6}},
		{9, 25, []int{3, 6}},
		{10, 30, []int{10}},
	}
)

func TestMemorizedCutRod(t *testing.T) {
	for _, tc := range testCases {
		q, splitting := MemorizedCutRod(p, tc.n)
		assert.Equal(t, tc.q, q)
		assert.Equal(t, tc.splitting, splitting)
	}
}

func TestBottomUpCutRod(t *testing.T) {
	for _, tc := range testCases {
		q, splitting := BottomUpCutRod(p, tc.n)
		assert.Equal(t, tc.q, q)
		assert.Equal(t, tc.splitting, splitting)
	}
}
