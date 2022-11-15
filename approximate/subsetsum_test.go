package approximate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApproxSubsetSum(t *testing.T) {
	sum := ApproxSubsetSum([]float64{104, 102, 201, 101}, 308, 0.4)
	assert.InDelta(t, 302, sum, 0.1)
}
