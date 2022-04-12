package matrixmultiply

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParallelMultiplyRecursive(t *testing.T) {
	result := ParallelMultiplyRecursive(aSmall, bSmall)
	assert.Equal(t, expectedCSmall, result)

	result = multiplyDivideAndRule(a, b)
	assert.Equal(t, expectedC, result)
}
