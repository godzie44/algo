package bug

import "math"

type costs []float64
type weights []int

func Optimize(v costs, w weights, bugW int) (float64, []int) {
	optV := make([]float64, bugW+1)
	optV[0] = 0
	optPath := make([]int, bugW+1)
	optPath[0] = 0

	for i := 1; i <= bugW; i++ {
		q := math.Inf(-1)
		objI := 0
		for j := 0; j < len(v); j++ {
			if w[j] > i {
				continue
			}

			if !check(optPath, w, j, i) {
				continue
			}

			if q < v[j]+optV[i-w[j]] {
				q = v[j] + optV[i-w[j]]
				objI = j
			}
		}

		optV[i] = q
		optPath[i] = objI
	}

	return optV[bugW], optObjects(optPath, w, bugW)
}

func check(ww []int, w weights, target int, bugW int) bool {
	weight := bugW - w[target]

	for weight != 0 {
		if ww[weight] == target {
			return false
		}
		weight = weight - w[ww[weight]]
	}

	return true
}

func optObjects(path []int, w []int, bugW int) []int {
	var result []int

	for bugW != 0 {
		result = append(result, path[bugW])

		bugW = bugW - w[path[bugW]]
	}

	return result
}
