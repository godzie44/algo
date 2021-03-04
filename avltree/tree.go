package avltree

import (
	"fmt"
	"io"
	"strings"
)

type Node struct {
	val                 int
	h                   int
	left, right, parent *Node
}

type Tree struct {
	root *Node
	Nil  *Node
}

func NewTree() *Tree {
	nilNode := &Node{}
	return &Tree{Nil: nilNode, root: nilNode}
}

func (t *Tree) Visualize(w io.Writer) {
	v := treeVisualizer{w, t.Nil}
	v.runOnNode(t.root, "", false)
}

type treeVisualizer struct {
	w   io.Writer
	nil *Node
}

func (v *treeVisualizer) runOnNode(node *Node, prefix string, hasRightSister bool) {
	if node == v.nil { // if node == t.Nil
		return
	}

	printPrefix := strings.TrimRight(prefix, "│")
	if hasRightSister {
		printPrefix += "├"
	} else {
		prefix = printPrefix + " "
		printPrefix += "└"
	}

	fmt.Fprintf(v.w, "%s── %d - h(%d) \n", printPrefix, node.val, node.h)

	prefix += "   " + "│"
	v.runOnNode(node.left, prefix, node.right != v.nil)
	v.runOnNode(node.right, prefix, false)
}

func (t *Tree) Add(val int) {
	var newNode *Node
	y := t.Nil
	x := t.root
	for x != t.Nil {
		y = x
		if val < x.val {
			x = x.left
		} else {
			x = x.right
		}
	}
	if y == t.Nil {
		newNode = &Node{val: val, parent: t.Nil, left: t.Nil, right: t.Nil}
		t.root = newNode
	} else {
		newNode = &Node{val: val, parent: y, left: t.Nil, right: t.Nil}

		if val < y.val {
			y.left = newNode
		} else {
			y.right = newNode
		}

		t.propagateH(newNode, 1)
	}

	for next := newNode.parent; next != t.Nil; next = next.parent {
		t.balance(next)
	}
}

func (t *Tree) propagateH(n *Node, newHeight int) {
	if n.h < newHeight {
		n.h = newHeight
	}

	if n.parent != t.Nil {
		t.propagateH(n.parent, newHeight+1)
	}
}

func (t *Tree) balance(x *Node) {
	if x.right.h-x.left.h > 1 {
		if x.right.right.h-x.right.left.h < 0 {
			t.rightRotate(x.right)
		}

		t.leftRotate(x)
		t.recalculateH()
	} else if x.left.h-x.right.h > 1 {
		if x.left.right.h-x.left.left.h > 0 {
			t.leftRotate(x.left)
		}

		t.rightRotate(x)
		t.recalculateH()
	}
}

func (t *Tree) recalculateH() {
	if t.root == t.Nil {
		return
	}

	t.walk(t.root, func(n *Node) {
		n.h = 0
	})

	t.walk(t.root, func(n *Node) {
		if n.left == t.Nil && n.right == t.Nil {
			t.propagateH(n, 1)
		}
	})
}

func (t *Tree) walk(n *Node, f func(n *Node)) {
	f(n)

	if n.left != t.Nil {
		t.walk(n.left, f)
	}
	if n.right != t.Nil {
		t.walk(n.right, f)
	}
}

func (t *Tree) leftRotate(x *Node) {
	y := x.right

	x.right = y.left
	if y.left != t.Nil {
		y.left.parent = x
	}

	y.parent = x.parent
	if x.parent == t.Nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *Tree) rightRotate(y *Node) {
	x := y.left

	y.left = x.right
	if x.right != t.Nil {
		x.right.parent = y
	}

	x.parent = y.parent
	if y.parent == t.Nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.right = y
	y.parent = x
}

func (t *Tree) Order() (result []int) {
	x := t.root
	var stack []*Node
	for len(stack) != 0 || x != t.Nil {
		if x != t.Nil {
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

func (t *Tree) Search(val int) *Node {
	x := t.root
	for x != nil && x.val != val {
		if val < x.val {
			x = x.left
		} else {
			x = x.right
		}
	}

	if x == t.Nil {
		return nil
	}

	return x
}

func (t *Tree) Min() *Node {
	x := t.root
	for x.left != t.Nil {
		x = x.left
	}
	return x
}

func (t *Tree) Max() *Node {
	x := t.root
	for x.right != t.Nil {
		x = x.right
	}
	return x
}

func (t *Tree) Successor(n *Node) *Node {
	if n.right != t.Nil {
		rTree := &Tree{root: n.right, Nil: t.Nil}
		return rTree.Min()
	}

	y := n.parent
	for y != t.Nil && n == y.right {
		n = y
		y = y.parent
	}
	return y
}

func (t *Tree) Predecessor(n *Node) *Node {
	if n.left != t.Nil {
		rTree := &Tree{root: n.left, Nil: t.Nil}
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
	balanceFrom := n.parent

	if n.left == t.Nil {
		t.transplant(n, n.right)
		t.propagateH(n.right, n.right.h)
	} else if n.right == t.Nil {
		t.transplant(n, n.left)
		t.propagateH(n.left, n.left.h)
	} else {
		rTree := &Tree{root: n.right, Nil: t.Nil}
		y := rTree.Min()

		if y.parent != n {
			balanceFrom = y.parent

			t.transplant(y, y.right)
			y.right = n.right
			y.right.parent = y
		} else {
			balanceFrom = y
		}

		t.transplant(n, y)
		y.left = n.left
		y.left.parent = y

		t.propagateH(y.left, y.left.h)
	}

	for p := balanceFrom; p != t.Nil; p = p.parent {
		t.balance(p)
	}
}

func (t *Tree) transplant(u, v *Node) {
	if u.parent == t.Nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}
