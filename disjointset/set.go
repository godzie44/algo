package disjointset

type Set struct {
	Val    rune
	parent *Set
	rank   int
}

func NewSet(x rune) *Set {
	set := &Set{
		Val:  x,
		rank: 0,
	}
	set.parent = set

	return set
}

func Union(x, y *Set) {
	link(FindSet(x), FindSet(y))
}

func link(x, y *Set) {
	if x.rank > y.rank {
		y.parent = x
	} else {
		x.parent = y
		if x.rank == y.rank {
			y.rank++
		}
	}
}

func FindSet(x *Set) *Set {
	if x != x.parent {
		x.parent = FindSet(x.parent)
	}
	return x.parent
}
