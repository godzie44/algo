package fibheap

import (
	"fmt"
	"math"
)

type Node struct {
	degree int
	parent *Node
	child  []*Node
	mark   bool

	Key int

	next *Node
	prev *Node
}

type (
	Heap struct {
		n   int
		min *Node

		rootListHead *Node
	}
)

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Insert(key int) *Node {
	x := &Node{
		degree: 0,
		parent: nil,
		child:  nil,
		mark:   false,
		Key:    key,
	}

	if h.min == nil {
		h.min = x
		h.addToRoots(x)
	} else {
		h.addToRoots(x)

		if x.Key < h.min.Key {
			h.min = x
		}
	}

	h.n++

	return x
}

func (h *Heap) addToRoots(x *Node) {
	x.next = h.rootListHead

	if h.rootListHead != nil {
		h.rootListHead.prev = x
	}
	h.rootListHead = x
}

func (h *Heap) delFromRoot(x *Node) {
	if x.prev != nil {
		x.prev.next = x.next
	}

	if x.next != nil {
		x.next.prev = x.prev
	}

	if h.rootListHead == x {
		h.rootListHead = x.next
	}
}

func (h *Heap) Min() int {
	return h.min.Key
}

func (h *Heap) UnionWith(heap *Heap) *Heap {
	newHeap := &Heap{
		n:   h.n + heap.n,
		min: h.min,
	}

	newHeap.rootListHead = h.rootListHead
	h.rootListHead.next = heap.rootListHead
	heap.rootListHead.prev = h.rootListHead

	if h.min == nil && (heap.min != nil && heap.min.Key < h.min.Key) {
		newHeap.min = heap.min
	}

	return newHeap
}

func (h *Heap) ExtractMin() *Node {
	z := h.min
	if z != nil {
		for _, c := range z.child {
			h.addToRoots(c)
			c.parent = nil
		}
		h.delFromRoot(z)
		if h.rootListHead == nil {
			h.min = nil
		} else {
			h.min = h.rootListHead
			h.consolidate()
		}
		h.n--
	}

	return z
}

func (h *Heap) roots() (result []*Node) {
	w := h.rootListHead
	for w != nil {
		result = append(result, w)
		w = w.next
	}
	return
}

func (h *Heap) consolidate() {
	a := make([]*Node, calcD(h.n))

	for _, w := range h.roots() {
		x := w
		d := x.degree
		for a[d] != nil {
			y := a[d]
			if x.Key > y.Key {
				x, y = y, x
			}
			h.link(y, x)
			a[d] = nil
			d++
		}
		a[d] = x
	}

	h.min = nil

	for i := 0; i < len(a); i++ {
		if a[i] != nil {
			if h.min == nil {
				h.rootListHead = nil
				h.addToRoots(a[i])
				h.min = a[i]
			} else {
				h.addToRoots(a[i])
				if a[i].Key < h.min.Key {
					h.min = a[i]
				}
			}
		}
	}
}

func (h *Heap) link(y, x *Node) {
	h.delFromRoot(y)
	x.child = append(x.child, y)
	y.parent = x
	x.degree++
	y.mark = false
}

func (h *Heap) DecreaseKey(x *Node, key int) error {
	if key > x.Key {
		return fmt.Errorf("new Key greater then current")
	}

	x.Key = key
	y := x.parent

	if y != nil && x.Key < y.Key {
		h.cut(x, y)
		h.cascadingCut(y)
	}

	if x.Key < h.min.Key {
		h.min = x
	}

	return nil
}

func (h *Heap) cut(x, y *Node) {
	for i := range y.child {
		if y.child[i] == x {
			y.child = append(y.child[:i], y.child[i+1:]...)
			y.degree--
			break
		}
	}

	h.addToRoots(x)
	x.parent = nil
	x.mark = false
}

func (h *Heap) cascadingCut(y *Node) {
	z := y.parent
	if z != nil {
		if y.mark == false {
			y.mark = true
		} else {
			h.cut(y, z)
			h.cascadingCut(z)
		}
	}
}

func (h *Heap) Delete(x *Node) {
	_ = h.DecreaseKey(x, math.MinInt64)
	h.ExtractMin()
}

func calcD(n int) int {
	return int(math.Floor(math.Log(float64(n)) / math.Log(1.61803)))
}
