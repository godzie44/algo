package linear

import (
	"errors"
	"math"
)

const zero = 1e-9

type Matrix [][]float64
type Vector []float64

func (v Vector) min() (int, float64) {
	var min = math.Inf(1)
	var ind = -1
	for i, val := range v {
		if val < min {
			min = val
			ind = i
		}
	}
	return ind, min
}

type indexVector struct {
	vals   []int
	search map[int]int
}

func newIndexVector() *indexVector {
	return &indexVector{
		vals:   []int{},
		search: map[int]int{},
	}
}

func (iv *indexVector) index(val int) int {
	return iv.search[val]
}

func (iv *indexVector) exists(val int) bool {
	_, ex := iv.search[val]
	return ex
}

func (iv *indexVector) replace(prev, new int) {
	v := iv.search[prev]
	delete(iv.search, prev)
	iv.vals[v] = new
	iv.search[new] = v
}

func (iv *indexVector) add(val int) {
	iv.vals = append(iv.vals, val)
	iv.search[val] = len(iv.vals) - 1
}

func (iv *indexVector) delete(val int) {
	idx := iv.index(val)
	delete(iv.search, val)
	iv.vals = append(iv.vals[:idx], iv.vals[idx+1:]...)
	for i := idx; i < len(iv.vals); i++ {
		iv.search[iv.vals[i]] = i
	}
}

func (iv *indexVector) forAnyIndexGreater(c Vector, target float64) bool {
	for _, j := range iv.vals {
		if c[iv.index(j)] > target {
			return true
		}
	}
	return false
}

func pivot(nVar, bVar *indexVector, a Matrix, b, c Vector, v float64, l, e int) (*indexVector, *indexVector, Matrix, Vector, Vector, float64) {
	row, col := bVar.index(l), nVar.index(e)
	mLen, nLen := len(b), len(c)

	div := a[row][col]
	b[row] /= div
	a[row][col] = 1.0
	for i := range a[row] {
		a[row][i] /= div
	}

	for i := 0; i < mLen; i++ {
		if i == row {
			continue
		}

		b[i] -= a[i][col] * b[row]
		for j := 0; j < nLen; j++ {
			if j == col {
				continue
			}
			a[i][j] -= a[i][col] * a[row][j]
		}
		a[i][col] /= -div
	}

	v += c[col] * b[row]

	for j := 0; j < nLen; j++ {
		if j == col {
			continue
		}
		c[j] -= c[col] * a[row][j]
	}
	c[col] /= -div

	nVar.replace(e, l)
	bVar.replace(l, e)

	return nVar, bVar, a, b, c, v
}

func Simplex(a Matrix, b, c Vector) (Vector, error) {
	nVars, bVars, a, b, c, v, err := initSimplex(a, b, c)
	if err != nil {
		return nil, err
	}

	delta := make(Vector, len(a))
	for nVars.forAnyIndexGreater(c, zero) {
		var e int
		for j := range c {
			if c[j] > 0 {
				e = nVars.vals[j]
				break
			}
		}

		for _, i := range bVars.vals {
			if a[bVars.index(i)][e] > 0 {
				delta[bVars.index(i)] = b[bVars.index(i)] / a[bVars.index(i)][e]
			} else {
				delta[bVars.index(i)] = math.Inf(1)
			}
		}

		l, deltaL := delta.min()
		if math.IsInf(deltaL, 1) {
			return nil, errors.New("no solution")
		} else {
			nVars, bVars, a, b, c, v = pivot(nVars, bVars, a, b, c, v, bVars.vals[l], e)
		}
	}

	x := make(Vector, len(c))
	for i := range x {
		if bVars.exists(i) {
			x[i] = b[bVars.index(i)]
		}
	}

	return x, nil
}

func initSimplex(a Matrix, b, c Vector) (*indexVector, *indexVector, Matrix, Vector, Vector, float64, error) {
	mLen, nLen := len(b), len(c)

	nVars, bVars := newIndexVector(), newIndexVector()
	for i := 0; i < nLen; i++ {
		nVars.add(i)
	}
	for i := len(a[0]); i < nLen+mLen; i++ {
		bVars.add(i)
	}

	k, bk := b.min()
	if bk >= 0 {
		return nVars, bVars, a, b, c, 0, nil
	}

	cStash := c
	c = append(make([]float64, len(c)), -1)
	for i := range a {
		a[i] = append(a[i], -1)
	}
	l := nLen + k
	nVars.add(nLen + mLen)

	var v float64
	nVars, bVars, a, b, c, v = pivot(nVars, bVars, a, b, c, v, l, nLen+mLen)

	delta := make(Vector, len(a))
	counter := 1
	for nVars.forAnyIndexGreater(c, zero) {
		var e int

		if counter%5 == 0 {
			for j := len(c) - 1; j >= 0; j-- {
				if c[j] > zero {
					e = nVars.index(j)
					break
				}
			}
		} else {
			for j := range c {
				if c[j] > zero {
					e = nVars.index(j)
					break
				}
			}
		}
		counter++

		for _, i := range bVars.vals {
			if a[bVars.index(i)][e] > 0 {
				delta[bVars.index(i)] = b[bVars.index(i)] / a[bVars.index(i)][e]
			} else {
				delta[bVars.index(i)] = math.Inf(1)
			}
		}

		l, deltaL := delta.min()
		if math.IsInf(deltaL, 1) {
			return nil, nil, nil, nil, nil, 0, errors.New("no solution")
		} else {
			nVars, bVars, a, b, c, v = pivot(nVars, bVars, a, b, c, v, bVars.vals[l], e)
		}
	}

	x := make(Vector, nLen+mLen+1)
	for i := range x {
		if bVars.exists(i) {
			x[i] = b[bVars.index(i)]
		}
	}

	if math.Abs(x[nLen+mLen]) < zero {
		if bVars.exists(nLen + mLen) {
			i, l := nLen+mLen, nLen+mLen

			var e int
			for _, nVar := range nVars.vals {
				if math.Abs(a[i][nVar]) > zero {
					e = nVar
				}
			}
			nVars, bVars, a, b, c, v = pivot(nVars, bVars, a, b, c, v, l, e)
		}
		col := nVars.index(nLen + mLen)
		for i := range a {
			a[i] = append(a[i][:col], a[i][col+1:]...)
		}
		c = make([]float64, nLen)
		v = 0
		nVars.delete(nLen + mLen)

		for i, val := range cStash {
			if math.Abs(val) < zero {
				continue
			} else if bVars.exists(i) {
				index := bVars.index(i)
				tmp := make([]float64, len(a[index]))
				for j := range a[index] {
					tmp[j] = -val * a[index][j]
				}
				for j := range c {
					c[j] += tmp[j]
				}
				v += val * b[index]
			} else {
				index := nVars.index(i)
				c[index] += val
			}
		}
		return nVars, bVars, a, b, c, v, nil
	}

	return nil, nil, nil, nil, nil, 0, errors.New("no solution")
}
