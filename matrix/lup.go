package matrix

import (
	"math"
)

func lupDecomposition(a [][]float64) ([][]float64, [][]float64, []int) {
	n := len(a)
	pv := make([]int, n)

	for k := range pv {
		pv[k] = k
	}

	for k := 0; k < n; k++ {
		var p float64
		ks := 0
		for i := k; i < n; i++ {
			if math.Abs(a[i][k]) > p {
				p = math.Abs(a[i][k])
				ks = i
			}
		}

		if p == 0 {
			panic("degenerate matrix")
		}

		pv[k], pv[ks] = pv[ks], pv[k]

		for i := 0; i < n; i++ {
			a[k][i], a[ks][i] = a[ks][i], a[k][i]
		}

		for i := k + 1; i < n; i++ {
			a[i][k] = a[i][k] / a[k][k]
			for j := k + 1; j < n; j++ {
				a[i][j] -= a[i][k] * a[k][j]
			}
		}
	}

	l, u := makeMatrix(n), makeMatrix(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > j {
				l[i][j] = a[i][j]
			} else {
				u[i][j] = a[i][j]
			}

			if i == j {
				l[i][j] = 1
			}
		}
	}

	return l, u, pv
}

func makeMatrix(n int) [][]float64 {
	m := make([][]float64, n)

	for i := 0; i < n; i++ {
		m[i] = make([]float64, n)
	}

	return m
}

func LUPSolve(a [][]float64, b []float64) []float64 {
	n := len(a)
	l, u, p := lupDecomposition(a)

	y := make([]float64, n)
	for i := 0; i < n; i++ {
		var sum float64
		for j := 0; j <= i-1; j++ {
			sum += l[i][j] * y[j]
		}

		y[i] = b[p[i]] - sum
	}

	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		var sum float64
		for j := i + 1; j < n; j++ {
			sum += u[i][j] * x[j]
		}
		x[i] = (y[i] - sum) / u[i][i]
	}

	return x
}
