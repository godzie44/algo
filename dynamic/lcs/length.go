package lcs

type arrow int

const (
	topLeft = 3
	top     = 2
	left    = 1
)

func Length(x, y string) (int, string) {
	m := len(x) + 1
	n := len(y) + 1

	c := make([][]int, m)
	for i := range c {
		c[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if x[i-1] == y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
			} else {
				c[i][j] = c[i][j-1]
			}
		}
	}

	return c[m-1][n-1], reestablishLCS(c, x, m-1, n-1)
}

func reestablishLCS(c [][]int, x string, i, j int) string {
	if i == 0 || j == 0 {
		return ""
	}

	if c[i][j] == c[i-1][j] {
		return reestablishLCS(c, x, i-1, j)
	} else if c[i][j] == c[i][j-1] {
		return reestablishLCS(c, x, i, j-1)
	} else {
		return reestablishLCS(c, x, i-1, j-1) + string(x[i-1])
	}
}
