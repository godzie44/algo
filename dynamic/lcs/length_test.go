package lcs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testCases = []struct {
	x, y        string
	expectedLen int
	expectedLcs string
}{
	{"ABCBDAB", "BDCABA", 4, "BCBA"},
	{"ABCBDABCCC", "BDCABACAC4", 6, "BCBACC"},
}

func TestLCSLength(t *testing.T) {
	for _, tc := range testCases {
		lcsLen, lcs := Length(tc.x, tc.y)
		assert.Equal(t, tc.expectedLen, lcsLen)
		assert.Equal(t, tc.expectedLcs, lcs)
	}
}
