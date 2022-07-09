package fft

import (
	"math"
	"math/bits"
)

func RecursiveFFT(a []float64) []float64 {
	n := len(a)
	if n == 1 {
		return a
	}

	wn := math.Exp(2 * math.Pi / float64(n))
	w := float64(1)

	a0 := make([]float64, 0)
	a1 := make([]float64, 0)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a0 = append(a0, a[i])
		} else {
			a1 = append(a1, a[i])

		}
	}

	y0 := RecursiveFFT(a0)
	y1 := RecursiveFFT(a1)

	y := make([]float64, n)
	for k := 0; k <= n/2-1; k++ {
		y[k] = y0[k] + w*y1[k]
		y[k+n/2] = y0[k] - w*y1[k]
		w = w * wn
	}

	return y
}

func IterativeFFT(a []float64) map[uint]float64 {
	arr := bitRecursiveCopy(a)

	n := len(a)

	for s := 1; s <= int(math.Log2(float64(n))); s++ {
		m := math.Pow(2, float64(s))
		wm := math.Exp(2 * math.Pi / m)

		for k := 0; k <= n-1; k += int(m) {
			w := float64(1)
			for j := 0; j <= int(m)/2-1; j += 1 {
				t := w * arr[uint(k+j+int(m)/2)]
				u := arr[uint(k+j)]
				arr[uint(k+j)] = u + t
				arr[uint(k+j+int(m)/2)] = u - t
				w = w * wm
			}
		}
	}

	return arr
}

func bitRecursiveCopy(a []float64) map[uint]float64 {
	lgn := math.Log2(float64(len(a)))

	result := make(map[uint]float64)
	for k := 0; k < len(a); k++ {
		result[bits.Reverse(uint(k))>>(64-int(lgn))] = a[k]
	}
	return result
}
