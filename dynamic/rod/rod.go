package rod

import "math"

type costList []float64

func MemorizedCutRod(p costList, n int) (float64, []int) {
	r := make([]float64, n+1)
	s := make([]int, n+1)
	for i := 0; i <= n; i++ {
		r[i] = math.Inf(-1)
	}

	q := memorizedCutRodAux(p, n, r, s)

	return q, createSplitPath(s, n)
}

func memorizedCutRodAux(p costList, n int, r []float64, s []int) float64 {
	if r[n] >= 0 {
		return r[n]
	}

	var q = math.Inf(-1)
	if n == 0 {
		q = 0
	} else {
		for i := 1; i < n+1; i++ {
			sm := p[i] + memorizedCutRodAux(p, n-i, r, s)
			if q < sm {
				q = sm
				s[n] = i
			}
		}
	}

	r[n] = q
	return q
}

func BottomUpCutRod(p costList, n int) (float64, []int) {
	r := make([]float64, n+1)
	s := make([]int, n+1)
	r[0] = 0

	for j := 1; j <= n; j++ {
		q := math.Inf(-1)
		for i := 1; i <= j; i++ {
			if q < p[i]+r[j-i] {
				q = p[i] + r[j-i]
				s[j] = i
			}
		}
		r[j] = q
	}

	return r[n], createSplitPath(s, n)
}

func createSplitPath(s []int, n int) []int {
	splitting := make([]int, 0)
	for n > 0 {
		splitting = append(splitting, s[n])
		n = n - s[n]
	}
	return splitting
}
