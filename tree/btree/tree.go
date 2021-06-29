package btree

import (
	"fmt"
	"io"
)

type Node struct {
	Keys  []int
	Child []*Node
	Leaf  bool
}

type Tree struct {
	root *Node
	t    int
}

func NewTree(t int) *Tree {
	return &Tree{t: t, root: &Node{Leaf: true}}
}

func (t *Tree) Visualize(w io.Writer) {
	v := treeVisualizer{w}
	v.runOnNode(t.root, "")
}

type treeVisualizer struct {
	w io.Writer
}

func (v *treeVisualizer) runOnNode(node *Node, prefix string) {
	if node == nil {
		return
	}

	for i := range node.Keys {
		if !node.Leaf {
			v.runOnNode(node.Child[i], prefix+"   ├ ")
		}

		if node.Leaf {
			fmt.Fprintf(v.w, "%s%02d \n", prefix, node.Keys[i])
		} else {
			fmt.Fprintf(v.w, "%s%02d ┤\n", prefix, node.Keys[i])
		}

		if !node.Leaf && i == len(node.Keys)-1 {
			v.runOnNode(node.Child[i+1], prefix+"   ├ ")
		}
	}
}

func (t *Tree) Max() int {
	n := t.root
	for !n.Leaf {
		n = n.Child[len(n.Child)-1]
	}
	return n.Keys[len(n.Keys)-1]
}

func (t *Tree) Min() int {
	n := t.root
	for !n.Leaf {
		n = n.Child[0]
	}
	return n.Keys[0]
}

func (t *Tree) Search(key int) (*Node, int) {
	return search(t.root, key)
}

func search(x *Node, key int) (*Node, int) {
	var i int
	for i = 0; i < len(x.Keys) && key > x.Keys[i]; i++ {
	}
	if i < len(x.Keys) && key == x.Keys[i] {
		return x, i
	} else if x.Leaf {
		return nil, 0
	} else {
		return search(x.Child[i], key)
	}
}

func (t *Tree) midElIdx() int {
	return t.t - 1
}

func (t *Tree) splitChild(x *Node, i int) {
	y := x.Child[i]
	z := &Node{Leaf: y.Leaf}

	for j := 1; j < t.t; j++ {
		z.Keys = append(z.Keys, y.Keys[t.midElIdx()+j])
	}

	if !y.Leaf {
		for j := 1; j < t.t+1; j++ {
			z.Child = append(z.Child, y.Child[t.midElIdx()+j])
		}
		y.Child = y.Child[:t.midElIdx()+1]
	}

	x.Child = append(x.Child, nil)
	for j := len(x.Child) - 1; j > i+1; j-- {
		x.Child[j] = x.Child[j-1]
	}
	x.Child[i+1] = z

	x.Keys = append(x.Keys, 0)
	for j := len(x.Keys) - 1; j > i; j-- {
		x.Keys[j] = x.Keys[j-1]
	}
	x.Keys[i] = y.Keys[t.midElIdx()]
	y.Keys = y.Keys[:t.midElIdx()]
}

func (t *Tree) insert(key int) {
	r := t.root
	if len(r.Keys) == 2*t.t-1 {
		s := &Node{Leaf: false, Child: []*Node{r}}
		t.root = s
		t.splitChild(s, 0)
		t.insertNonFull(s, key)
	} else {
		t.insertNonFull(r, key)
	}
}

func (t *Tree) Insert(keys ...int) {
	for _, k := range keys {
		t.insert(k)
	}
}

func (t *Tree) insertNonFull(x *Node, key int) {
	i := len(x.Keys) - 1

	if x.Leaf {
		x.Keys = append(x.Keys, 0)
		for ; i >= 0 && key < x.Keys[i]; i-- {
			x.Keys[i+1] = x.Keys[i]
		}
		x.Keys[i+1] = key
	} else {
		if key < x.Keys[0] {
			i = 0
		} else {
			for ; i >= 0 && key < x.Keys[i]; i-- {
			}
			i++
		}

		if len(x.Child[i].Keys) == 2*t.t-1 {
			t.splitChild(x, i)
			if key > x.Keys[i] {
				i++
			}
		}
		t.insertNonFull(x.Child[i], key)
	}
}

func (t *Tree) Order() (result []int) {
	t.root.walk(func(k int) {
		result = append(result, k)
	})

	return
}

func searchParent(node *Node, child *Node) (*Node, int) {
	for i, c := range node.Child {
		if c == child {
			return node, i
		}

		if res, ind := searchParent(c, child); res != nil {
			return res, ind
		}
	}
	return nil, 0
}

func (t *Tree) Delete(keys ...int) {
	for _, v := range keys {
		t.delete(v)
	}
}

func (t *Tree) delete(key int) {
	x, _ := t.Search(key)

	if t.root == x && x.Leaf {
		x.clearKey(key)
		return
	} else if x.Leaf {
		if len(x.Keys) > t.t-1 {
			x.clearKey(key)
			return
		} else {
			parent, idx := searchParent(t.root, x)
			if len(parent.Child)-1 > idx && len(parent.Child[idx+1].Keys) >= t.t {
				x.clearKey(key)
				x.Keys = append(x.Keys, parent.Keys[idx])
				parent.Keys[idx] = parent.Child[idx+1].Keys[0]
				parent.Child[idx+1].Keys = parent.Child[idx+1].Keys[1:]
				return
			}
			if idx > 0 && len(parent.Child[idx-1].Keys) >= t.t {
				x.clearKey(key)
				x.Keys = append([]int{parent.Keys[idx-1]}, x.Keys...) // вставляем в начало а не конец
				parent.Keys[idx-1] = parent.Child[idx-1].Keys[len(parent.Child[idx-1].Keys)-1]
				parent.Child[idx-1].Keys = parent.Child[idx-1].Keys[:len(parent.Child[idx-1].Keys)-1]
				return
			}

			x.clearKey(key)
			t.recursiveMerge(x, parent, idx)
		}
	} else {
		i := x.findIdx(key)
		temporaryDelete := (&Tree{root: x.Child[i], t: t.t}).Max()
		t.Delete(temporaryDelete)
		x, _ = t.Search(key)
		x.replaceKey(key, temporaryDelete)
	}
}

func (t *Tree) recursiveMerge(x *Node, parent *Node, idx int) {
	var mergeCandidate, mergedNode *Node
	var child []*Node
	var keys []int
	if idx > 0 {
		mergeCandidate = parent.Child[idx-1]
		child = append(mergeCandidate.Child, x.Child...)
		keys = append(mergeCandidate.Keys, append([]int{parent.Keys[idx-1]}, x.Keys...)...)
		parent.clearKey(parent.Keys[idx-1])
		mergedNode = &Node{Leaf: x.Leaf, Child: child, Keys: keys}
		parent.Child[idx] = mergedNode
		parent.Child = append(parent.Child[:idx-1], parent.Child[idx:]...)
	} else {
		mergeCandidate = parent.Child[idx+1]
		child = append(x.Child, mergeCandidate.Child...)
		keys = append(x.Keys, append([]int{parent.Keys[idx]}, mergeCandidate.Keys...)...)
		parent.clearKey(parent.Keys[idx])
		mergedNode = &Node{Leaf: x.Leaf, Child: child, Keys: keys}
		parent.Child[idx] = mergedNode
		parent.Child = append(parent.Child[:idx+1], parent.Child[idx+2:]...)
	}

	if parent == t.root {
		if len(parent.Keys) == 0 {
			t.root = mergedNode
		}
	} else if len(parent.Keys) < t.t-1 {
		nextParent, nextIDX := searchParent(t.root, parent)

		t.recursiveMerge(parent, nextParent, nextIDX)
	}
}

func (n *Node) clearKey(key int) {
	var i int
	for i = range n.Keys {
		if n.Keys[i] == key {
			break
		}
	}

	n.Keys = append(n.Keys[:i], n.Keys[i+1:]...)
}

func (n *Node) replaceKey(oldKey, newKey int) {
	for i := range n.Keys {
		if n.Keys[i] == oldKey {
			n.Keys[i] = newKey
			return
		}
	}
}

func (n *Node) findIdx(key int) int {
	for i := range n.Keys {
		if n.Keys[i] == key {
			return i
		}
	}
	return 0
}

func (n *Node) walk(fn func(int)) {
	if n.Leaf {
		for _, k := range n.Keys {
			fn(k)
		}
		return
	}

	var i int
	for i = range n.Keys {
		n.Child[i].walk(fn)
		fn(n.Keys[i])
	}
	n.Child[i+1].walk(fn)
}
