package bst

import "math"

func Optimal(p, q []float64) ([][]float64, [][]int) {
	n := len(p) - 1
	e, w, root := initArrs(n)

	for i := 1; i <= n+1; i++ {
		e[i][i-1] = q[i-1]
		w[i][i-1] = q[i-1]
	}

	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			e[i][j] = math.MaxFloat64
			w[i][j] = w[i][j-1] + p[j] + q[j]

			for r := i; r <= j; r++ {
				t := e[i][r-1] + e[r+1][j] + w[i][j]
				if t < e[i][j] {
					e[i][j] = t
					root[i][j] = r
				}
			}
		}
	}

	return e, root
}

func initArrs(n int) ([][]float64, [][]float64, [][]int) {
	e := make([][]float64, n+2)
	for i := range e {
		e[i] = make([]float64, n+1)
	}

	w := make([][]float64, n+2)
	for i := range w {
		w[i] = make([]float64, n+1)
	}

	root := make([][]int, n+1)
	for i := range root {
		root[i] = make([]int, n+1)
	}

	return e, w, root
}
