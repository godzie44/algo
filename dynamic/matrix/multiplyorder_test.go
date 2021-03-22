package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrixChainOrder(t *testing.T) {
	m, s := ChainOrder([]Matrix{{30, 35}, {35, 15}, {15, 5}, {5, 10}, {10, 20}, {20, 25}})

	assert.Equal(t, "((A1(A2A3))((A4A5)A6))", s)
	assert.Equal(t, 15125, m[1][6])
}
