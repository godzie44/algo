package approximate

import (
	"testing"
)

func TestGreedySetCover(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fSets := [][]int{
		{1, 2, 5, 6, 9, 10},
		{6, 7, 10, 11},
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{4, 8},
	}

	res := GreedySetCover(x, fSets)

	for _, i := range x {
		found := false
		for _, f := range res {
			for _, val := range f {
				if val == i {
					found = true
				}
			}
		}

		if !found {
			t.Errorf("Cover for el %d not found in sets", i)
		}
	}
}
