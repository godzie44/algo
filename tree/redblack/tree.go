package redblack

import (
	"fmt"
	"io"
	"strings"
)

type Color int

const (
	_ Color = iota
	Red
	Black
)

type Node struct {
	val                 int
	color               Color
	left, right, parent *Node
}

type Tree struct {
	root *Node
	Nil  *Node
}

func NewTree() *Tree {
	nilNode := &Node{color: Black}
	return &Tree{Nil: nilNode, root: nilNode}
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

	if node.color == Red {
		fmt.Fprintf(v.w, "%s── r%d \n", printPrefix, node.val)
	} else {
		fmt.Fprintf(v.w, "%s── b%d \n", printPrefix, node.val)
	}

	prefix += "   " + "│"
	v.runOnNode(node.left, prefix, node.right != nil)
	v.runOnNode(node.right, prefix, false)
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

func (t *Tree) Add(val int) {
	newNode := &Node{val: val}
	y := t.Nil
	x := t.root
	for x != t.Nil {
		y = x
		if newNode.val < x.val {
			x = x.left
		} else {
			x = x.right
		}
	}

	newNode.parent = y
	if y == t.Nil {
		t.root = newNode
	} else if newNode.val < y.val {
		y.left = newNode
	} else {
		y.right = newNode
	}
	newNode.left = t.Nil
	newNode.right = t.Nil
	newNode.color = Red
	t.insertFixup(newNode)
}

func (t *Tree) insertFixup(z *Node) {
	for z.parent.color == Red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			if y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = Black
				z.parent.parent.color = Red
				t.rightRotate(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == Red {
				z.parent.color = Black
				y.color = Black
				z.parent.parent.color = Red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = Black
				z.parent.parent.color = Red
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = Black
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

func (t *Tree) Delete(z *Node) {
	y := z
	yOriginalColor := y.color

	var x *Node
	if z.left == t.Nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == t.Nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		rTree := &Tree{
			root: z.right,
			Nil:  t.Nil,
		}
		y = rTree.Min()
		yOriginalColor = y.color
		x = y.right
		if y.parent == z {
			x.parent = y
		} else {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color
	}

	if yOriginalColor == Black {
		t.deleteFixUp(x)
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

func (t *Tree) deleteFixUp(x *Node) {
	for x != t.root && x.color == Black {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == Red {
				w.color = Black
				x.parent.color = Red
				t.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == Black && w.right.color == Black {
				w.color = Red
				x = x.parent
			} else {
				if w.right.color == Black {
					w.left.color = Black
					w.color = Red
					t.rightRotate(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = Black
				w.right.color = Black
				t.leftRotate(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == Red {
				w.color = Black
				x.parent.color = Red
				t.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.right.color == Black && w.left.color == Black {
				w.color = Red
				x = x.parent
			} else {
				if w.left.color == Black {
					w.right.color = Black
					w.color = Red
					t.leftRotate(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = Black
				w.left.color = Black
				t.rightRotate(x.parent)
				x = t.root
			}
		}
	}
	x.color = Black
}
