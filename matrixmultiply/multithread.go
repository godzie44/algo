package matrixmultiply

import (
	"sync"
)

func ParallelMultiplyRecursive(a, b [][]int) [][]int {
	return parallelMultiplyRecursive(makePart(a, 0, 0, len(a)), makePart(b, 0, 0, len(b))).source
}

func parallelMultiplyRecursive(a, b matrixPart) matrixPart {
	if a.n == 1 {
		return matrixToPart([][]int{{a.source[a.startRow][a.startCol] * b.source[b.startRow][b.startCol]}})
	} else {
		wg := &sync.WaitGroup{}

		newN := a.n / 2
		c, t := matrixToPart(initC(a.n, a.n)), matrixToPart(initC(a.n, a.n))

		a11 := a.cut(a.startRow, a.startCol, newN)
		a12 := a.cut(a.startRow, a.startCol+newN, newN)
		a21 := a.cut(a.startRow+newN, a.startCol, newN)
		a22 := a.cut(a.startRow+newN, a.startCol+newN, newN)
		b11 := b.cut(b.startRow, b.startCol, newN)
		b12 := b.cut(b.startRow, b.startCol+newN, newN)
		b21 := b.cut(b.startRow+newN, b.startCol, newN)
		b22 := b.cut(b.startRow+newN, b.startCol+newN, newN)

		c11, t11 := c.cut(0, 0, newN), t.cut(0, 0, newN)
		c12, t12 := c.cut(0, newN, newN), t.cut(0, newN, newN)
		c21, t21 := c.cut(newN, 0, newN), t.cut(newN, 0, newN)
		c22, t22 := c.cut(newN, newN, newN), t.cut(newN, newN, newN)

		spawn(wg, func() {
			c11.append(parallelMultiplyRecursive(a11, b11))
		})
		spawn(wg, func() {
			c12.append(parallelMultiplyRecursive(a11, b12))
		})
		spawn(wg, func() {
			c21.append(parallelMultiplyRecursive(a21, b11))
		})
		spawn(wg, func() {
			c22.append(parallelMultiplyRecursive(a21, b12))
		})
		spawn(wg, func() {
			t11.append(parallelMultiplyRecursive(a12, b21))
		})
		spawn(wg, func() {
			t12.append(parallelMultiplyRecursive(a12, b22))
		})
		spawn(wg, func() {
			t21.append(parallelMultiplyRecursive(a22, b21))
		})
		t22.append(parallelMultiplyRecursive(a22, b22))

		wg.Wait()

		spawn(wg, func() {
			c11.append(c11.sum(t11))
		})
		spawn(wg, func() {
			c12.append(c12.sum(t12))
		})
		spawn(wg, func() {
			c21.append(c21.sum(t21))
		})
		spawn(wg, func() {
			c22.append(c22.sum(t22))
		})

		wg.Wait()

		return c
	}
}

func spawn(wg *sync.WaitGroup, fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
}
