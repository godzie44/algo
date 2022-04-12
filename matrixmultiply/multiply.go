package matrixmultiply

func multiply(m1, m2 [][]int) [][]int {
	if len(m1) != len(m2[0]) {
		panic("l must be equal n")
	}

	n := len(m1)

	c := initC(n, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				c[i][j] += m1[i][k] * m2[k][j]
			}
		}
	}

	return c
}

func initC(rows, cols int) [][]int {
	result := make([][]int, rows)

	for i := range result {
		result[i] = make([]int, cols)
	}

	return result
}

type matrixPart struct {
	source             [][]int
	startRow, startCol int
	n                  int
}

func makePart(matrix [][]int, startRow int, startCol int, n int) matrixPart {
	return matrixPart{source: matrix, startRow: startRow, startCol: startCol, n: n}
}

func matrixToPart(matrix [][]int) matrixPart {
	return matrixPart{source: matrix, startRow: 0, startCol: 0, n: len(matrix)}
}

func (m matrixPart) sum(matrix matrixPart) matrixPart {
	res := make([][]int, m.n)
	for i := 0; i < m.n; i++ {
		res[i] = make([]int, m.n)
		for j := 0; j < m.n; j++ {
			res[i][j] = m.source[m.startRow+i][m.startCol+j] + matrix.source[matrix.startRow+i][matrix.startCol+j]
		}
	}

	return matrixToPart(res)
}

func (m matrixPart) diff(matrix matrixPart) matrixPart {
	res := make([][]int, m.n)
	for i := 0; i < m.n; i++ {
		res[i] = make([]int, m.n)
		for j := 0; j < m.n; j++ {
			res[i][j] = m.source[m.startRow+i][m.startCol+j] - matrix.source[matrix.startRow+i][matrix.startCol+j]
		}
	}

	return matrixToPart(res)
}

func (m matrixPart) append(matrix matrixPart) {
	for i := 0; i < m.n; i++ {
		for j := 0; j < m.n; j++ {
			m.source[m.startRow+i][m.startCol+j] = matrix.source[matrix.startRow+i][matrix.startCol+j]
		}
	}
}

func (m matrixPart) cut(row, col int, n int) matrixPart {
	return matrixPart{
		source:   m.source,
		startRow: row,
		startCol: col,
		n:        n,
	}
}

func multiplyDivideAndRule(a, b [][]int) [][]int {
	return multiplyRecursive(makePart(a, 0, 0, len(a)), makePart(b, 0, 0, len(b))).source
}

func multiplyRecursive(a, b matrixPart) matrixPart {
	if a.n == 1 {
		return matrixToPart([][]int{{a.source[a.startRow][a.startCol] * b.source[b.startRow][b.startCol]}})
	} else {
		newN := a.n / 2

		c := matrixToPart(initC(a.n, a.n))

		a11 := a.cut(a.startRow, a.startCol, newN)
		a12 := a.cut(a.startRow, a.startCol+newN, newN)
		a21 := a.cut(a.startRow+newN, a.startCol, newN)
		a22 := a.cut(a.startRow+newN, a.startCol+newN, newN)

		b11 := b.cut(b.startRow, b.startCol, newN)
		b12 := b.cut(b.startRow, b.startCol+newN, newN)
		b21 := b.cut(b.startRow+newN, b.startCol, newN)
		b22 := b.cut(b.startRow+newN, b.startCol+newN, newN)

		c11 := c.cut(0, 0, newN)
		c12 := c.cut(0, newN, newN)
		c21 := c.cut(newN, 0, newN)
		c22 := c.cut(newN, newN, newN)

		c11.append(multiplyRecursive(a11, b11).sum(multiplyRecursive(a12, b21)))
		c12.append(multiplyRecursive(a11, b12).sum(multiplyRecursive(a12, b22)))
		c21.append(multiplyRecursive(a21, b11).sum(multiplyRecursive(a22, b21)))
		c22.append(multiplyRecursive(a21, b12).sum(multiplyRecursive(a22, b22)))

		return c
	}
}

func multiplyShtrassen(a, b [][]int) [][]int {
	return multiplyRecursiveShtrassen(makePart(a, 0, 0, len(a)), makePart(b, 0, 0, len(b))).source
}

func multiplyRecursiveShtrassen(a, b matrixPart) matrixPart {
	if a.n == 1 {
		return matrixToPart([][]int{{a.source[a.startRow][a.startCol] * b.source[b.startRow][b.startCol]}})
	} else {
		newN := a.n / 2

		c := matrixToPart(initC(a.n, a.n))

		a11 := a.cut(a.startRow, a.startCol, newN)
		a12 := a.cut(a.startRow, a.startCol+newN, newN)
		a21 := a.cut(a.startRow+newN, a.startCol, newN)
		a22 := a.cut(a.startRow+newN, a.startCol+newN, newN)

		b11 := b.cut(b.startRow, b.startCol, newN)
		b12 := b.cut(b.startRow, b.startCol+newN, newN)
		b21 := b.cut(b.startRow+newN, b.startCol, newN)
		b22 := b.cut(b.startRow+newN, b.startCol+newN, newN)

		s1 := b12.diff(b22)
		s2 := a11.sum(a12)
		s3 := a21.sum(a22)
		s4 := b21.diff(b11)
		s5 := a11.sum(a22)
		s6 := b11.sum(b22)
		s7 := a12.diff(a22)
		s8 := b21.sum(b22)
		s9 := a11.diff(a21)
		s10 := b11.sum(b12)

		p1 := multiplyRecursiveShtrassen(a11, s1)
		p2 := multiplyRecursiveShtrassen(s2, b22)
		p3 := multiplyRecursiveShtrassen(s3, b11)
		p4 := multiplyRecursiveShtrassen(a22, s4)
		p5 := multiplyRecursiveShtrassen(s5, s6)
		p6 := multiplyRecursiveShtrassen(s7, s8)
		p7 := multiplyRecursiveShtrassen(s9, s10)

		c11 := c.cut(0, 0, newN)
		c12 := c.cut(0, newN, newN)
		c21 := c.cut(newN, 0, newN)
		c22 := c.cut(newN, newN, newN)

		c11.append(p5.sum(p4).diff(p2).sum(p6))
		c12.append(p1.sum(p2))
		c21.append(p3.sum(p4))
		c22.append(p5.sum(p1).diff(p3).diff(p7))

		return c
	}
}
