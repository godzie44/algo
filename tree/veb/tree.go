package veb

import "math"

type node struct {
	u        int
	min, max *int

	cluster []*node

	summary *node
}

func (n *node) isMember(x int) bool {
	if (n.min != nil && x == *n.min) || (n.max != nil && x == *n.max) {
		return true
	} else if n.u == 2 {
		return false
	} else {
		return n.cluster[n.high(x)].isMember(n.low(x))
	}
}

func botSqrt(u int) float64 {
	return math.Pow(2, math.Floor(math.Log2(float64(u))/2))
}

func topSqrt(u int) float64 {
	return math.Pow(2, math.Ceil(math.Log2(float64(u))/2))
}

func (n *node) high(x int) int {
	return int(math.Floor(float64(x) / botSqrt(n.u)))
}

func (n *node) low(x int) int {
	return x % int(botSqrt(n.u))
}

func (n *node) index(x, y int) int {
	return int(float64(x)*botSqrt(n.u)) + y
}

func (n *node) successor(x int) (int, bool) {
	if n.u == 2 {
		if x == 0 && *n.max == 1 {
			return 1, true
		}
		return 0, false
	}

	if n.min != nil && x < *n.min {
		return *n.min, true
	}

	maxLow := n.cluster[n.high(x)].max
	if maxLow != nil && n.low(x) < *maxLow {
		offset, _ := n.cluster[n.high(x)].successor(n.low(x))
		return n.index(n.high(x), offset), true
	}

	succCluster, exists := n.summary.successor(n.high(x))
	if !exists {
		return 0, false
	}

	offset := n.cluster[succCluster].min
	return n.index(succCluster, *offset), true
}

func (n *node) predecessor(x int) (int, bool) {
	if n.u == 2 {
		if x == 1 && *n.min == 0 {
			return 0, true
		}
		return 0, false
	}

	if n.max != nil && x > *n.max {
		return *n.max, true
	}

	minLow := n.cluster[n.high(x)].min
	if minLow != nil && n.low(x) > *minLow {
		offset, _ := n.cluster[n.high(x)].predecessor(n.low(x))
		return n.index(n.high(x), offset), true
	}

	predCluster, exists := n.summary.predecessor(n.high(x))
	if !exists {
		if n.min != nil && x > *n.min {
			return *n.min, true
		}
		return 0, false
	}

	offset := n.cluster[predCluster].max
	return n.index(predCluster, *offset), true
}

func (n *node) insert(x int) {
	if n.min == nil {
		n.insertEmpty(x)
	} else {
		if x < *n.min {
			oldX := x
			x = *n.min
			n.min = &oldX
		}

		if n.u > 2 {
			if n.cluster[n.high(x)].min == nil {
				n.summary.insert(n.high(x))
				n.cluster[n.high(x)].insertEmpty(n.low(x))
			} else {
				n.cluster[n.high(x)].insert(n.low(x))
			}
		}

		if x > *n.max {
			n.max = &x
		}
	}
}

func (n *node) insertEmpty(x int) {
	x1 := x
	n.min = &x1
	x2 := x
	n.max = &x2
}

func (n *node) delete(x int) {
	if *n.min == *n.max {
		n.min = nil
		n.max = nil
	} else if n.u == 2 {
		if x == 0 {
			one := 1
			n.min = &one
		} else {
			zero := 0
			n.min = &zero
		}
		*n.max = *n.min
	} else {
		if x == *n.min {
			firstCluster := n.summary.min
			x = n.index(*firstCluster, *n.cluster[*firstCluster].min)
			n.min = &x
		}

		n.cluster[n.high(x)].delete(n.low(x))

		if n.cluster[n.high(x)].min == nil {
			n.summary.delete(n.high(x))
			if x == *n.max {
				summaryMax := n.summary.max
				if summaryMax == nil {
					n.max = n.min
				} else {
					idx := n.index(*summaryMax, *n.cluster[*summaryMax].max)
					n.max = &idx
				}
			}
		} else if x == *n.max {
			idx := n.index(n.high(x), *n.cluster[n.high(x)].max)
			n.max = &idx
		}
	}
}

func newNode(u int) *node {
	n := &node{u: u}

	if u == 2 {
		return n
	}

	n.cluster = make([]*node, int(topSqrt(u)))
	for i := range n.cluster {
		n.cluster[i] = newNode(int(botSqrt(u)))
	}

	n.summary = newNode(int(topSqrt(u)))

	return n
}

type Tree struct {
	root *node
}

func NewTree(u int) *Tree {
	return &Tree{root: newNode(u)}
}

func (t *Tree) Insert(x int) {
	t.root.insert(x)
}

func (t *Tree) Max() *int {
	return t.root.max
}

func (t *Tree) Min() *int {
	return t.root.min
}

func (t *Tree) Successor(x int) (int, bool) {
	return t.root.successor(x)
}

func (t *Tree) Predecessor(x int) (int, bool) {
	return t.root.predecessor(x)
}

func (t *Tree) IsMember(x int) bool {
	return t.root.isMember(x)
}

func (t *Tree) Delete(x int) {
	t.root.delete(x)
}
