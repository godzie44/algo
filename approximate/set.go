package approximate

func GreedySetCover[T comparable](x []T, f [][]T) [][]T {
	u := make(map[T]struct{})
	for _, v := range x {
		u[v] = struct{}{}
	}

	fSetIndexes := make(map[int]struct{})
	for i := range f {
		fSetIndexes[i] = struct{}{}
	}

	c := make([][]T, 0)

	for len(u) != 0 {
		minULen := len(u)
		var setIndex int
		for fSetIdx := range fSetIndexes {
			candidate := len(u)
			for _, el := range f[fSetIdx] {
				_, exists := u[el]
				if exists {
					candidate--
				}
			}

			if candidate < minULen {
				minULen = candidate
				setIndex = fSetIdx
			}
		}

		for _, el := range f[setIndex] {
			delete(u, el)
		}
		delete(fSetIndexes, setIndex)

		c = append(c, f[setIndex])
	}

	return c
}
