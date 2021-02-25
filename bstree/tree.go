package bstree

import (
	"fmt"
	"io"
	"strings"
)

type Node struct {
	val                 int
	left, right, parent *Node
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Visualize(w io.Writer) {
	v := treeVisualizer{w}
	v.runOnNode(t.root, "", false)
}

type treeVisualizer struct {
	w io.Writer
}

func (v *treeVisualizer) runOnNode(node *Node, prefix string, hasRightSister bool) {
	if node == nil {
		return
	}

	printPrefix := strings.TrimRight(prefix, "│")
	if hasRightSister {
		printPrefix += "├"
	} else {
		prefix = printPrefix + " "
		printPrefix += "└"
	}

	fmt.Fprintf(v.w, "%s── %d \n", printPrefix, node.val)

	prefix += "   " + "│"
	v.runOnNode(node.left, prefix, node.right != nil)
	v.runOnNode(node.right, prefix, false)
}

func (t *Tree) Add(val int) {
	var y *Node
	x := t.root
	for x != nil {
		y = x
		if val < x.val {
			x = x.left
		} else {
			x = x.right
		}
	}
	if y == nil {
		t.root = &Node{val: val}
	} else {
		if val < y.val {
			y.left = &Node{val: val, parent: y}
		} else {
			y.right = &Node{val: val, parent: y}
		}
	}
}

func (t *Tree) Search(val int) *Node {
	x := t.root
	for x != nil && x.val != val {
		if val < x.val {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func (t *Tree) Min() *Node {
	x := t.root
	for x.left != nil {
		x = x.left
	}
	return x
}

func (t *Tree) Max() *Node {
	x := t.root
	for x.right != nil {
		x = x.right
	}
	return x
}

func (t *Tree) Successor(n *Node) *Node {
	if n.right != nil {
		rTree := &Tree{root: n.right}
		return rTree.Min()
	}

	y := n.parent
	for y != nil && n == y.right {
		n = y
		y = y.parent
	}
	return y
}

func (t *Tree) Predecessor(n *Node) *Node {
	if n.left != nil {
		rTree := &Tree{root: n.left}
		return rTree.Max()
	}

	y := n.parent
	for y != nil && n == y.left {
		n = y
		y = y.parent
	}
	return y
}

func (t *Tree) Delete(n *Node) {
	if n.left == nil {
		t.transplant(n, n.right)
	} else if n.right == nil {
		t.transplant(n, n.left)
	} else {
		rTree := &Tree{root: n.right}
		y := rTree.Min()
		if y.parent != n {
			t.transplant(y, y.right)
			y.right = n.right
			y.right.parent = y
		}
		t.transplant(n, y)
		y.left = n.left
		y.left.parent = y
	}
}

func (t *Tree) transplant(u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (t *Tree) Order() (result []int) {
	x := t.root
	var stack []*Node
	for len(stack) != 0 || x != nil {
		if x != nil {
			stack = append(stack, x)
			x = x.left
		} else {
			x = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, x.val)
			x = x.right
		}
	}
	return
}
