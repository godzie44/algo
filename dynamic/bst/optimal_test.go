package bst

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

var expectedTreeView = `k2 is root 
k1 is left child of k2 
d0 left child of k1 
d1 right child of k1 
k5 is right child of k2 
k4 is left child of k5 
k3 is left child of k4 
d2 left child of k3 
d3 right child of k3 
d4 right child of k4 
d5 right child of k5 
`

func TestOptimal(t *testing.T) {
	e, r := Optimal(
		[]float64{0, 0.15, 0.10, 0.05, 0.1, 0.2},
		[]float64{0.05, 0.1, 0.05, 0.05, 0.05, 0.1},
	)

	assert.InDelta(t, 2.75, e[1][5], 0.0001)

	buff := bytes.Buffer{}
	viewTree(&buff, r, 1, 5, "root")

	assert.Equal(t, expectedTreeView, buff.String())
}

func viewTree(w io.Writer, r [][]int, i, j int, parentInfo string) {
	k := r[i][j]
	fmt.Fprintf(w, "k%d is %s \n", k, parentInfo)

	lSubTreeI, lSubTreeJ := i, k-1
	hasLeftSubTree := lSubTreeI <= lSubTreeJ
	if hasLeftSubTree {
		viewTree(w, r, lSubTreeI, lSubTreeJ, fmt.Sprintf("left child of k%d", k))
	} else {
		fmt.Fprintf(w, "d%d left child of k%d \n", k-1, k)
	}

	rSubTreeI, rSubTreeJ := k+1, j
	hasRightSubTree := rSubTreeI <= rSubTreeJ
	if hasRightSubTree {
		viewTree(w, r, rSubTreeI, rSubTreeJ, fmt.Sprintf("right child of k%d", k))
	} else {
		fmt.Fprintf(w, "d%d right child of k%d \n", k, k)
	}
}
