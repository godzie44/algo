package approximate

func ApproxSubsetSum(s []float64, t float64, eps float64) float64 {
	n := len(s)
	lists := make([][]float64, n+1)
	lists[0] = []float64{0}

	for i := 1; i <= n; i++ {
		lists[i] = mergeList(lists[i-1], incSet(lists[i-1], s[i-1]))
		lists[i] = trim(lists[i], eps/(2*float64(n)))
		lists[i] = removeGt(lists[i], t)
	}
	return lists[n][len(lists[n])-1]
}

func removeGt(set []float64, t float64) []float64 {
	newSet := make([]float64, 0)
	for _, v := range set {
		if v <= t {
			newSet = append(newSet, v)
		}
	}
	return newSet
}

func trim(list []float64, eps float64) []float64 {
	newList := []float64{list[0]}
	last := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > last*(1+eps) {
			newList = append(newList, list[i])
			last = list[i]
		}
	}
	return newList
}

func mergeList(s1 []float64, s2 []float64) []float64 {
	newSet := make([]float64, 0, len(s1)+len(s2))

	i, j := 0, 0
	for {
		if i == len(s1) {
			newSet = append(newSet, s2[j:]...)
			break
		}
		if j == len(s2) {
			newSet = append(newSet, s1[i:]...)
			break
		}

		if s1[i] < s2[j] {
			newSet = append(newSet, s1[i])
			i++
		} else {
			newSet = append(newSet, s2[j])
			j++
		}
	}

	return newSet
}

func incSet(set []float64, x float64) []float64 {
	newSet := make([]float64, len(set))
	for i := range set {
		newSet[i] = set[i] + x
	}
	return newSet
}
