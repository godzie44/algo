package matrix

import (
	"fmt"
	"math"
)

type Matrix struct {
	a, b int
}

func ChainOrder(matrices []Matrix) ([][]int, string) {
	p := make([]int, len(matrices)+1)
	for i, m := range matrices {
		p[i] = m.a
		p[i+1] = m.b
	}

	n := len(p)

	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}

	s := make([][]int, n-1)
	for i := range s {
		s[i] = make([]int, n)
	}

	for l := 2; l < n; l++ {
		for i := 1; i <= n-l; i++ {
			j := i + l - 1

			m[i][j] = math.MaxInt64

			for k := i; k <= j-1; k++ {
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}

	return m, extractExpr(s, 1, len(matrices))
}

func extractExpr(s [][]int, i, j int) string {
	if i == j {
		return fmt.Sprintf("A%d", i)
	} else {
		return fmt.Sprintf("(") +
			extractExpr(s, i, s[i][j]) +
			extractExpr(s, s[i][j]+1, j) +
			fmt.Sprintf(")")
	}
}
