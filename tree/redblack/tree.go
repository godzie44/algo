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

type Comparable[T any] interface {
	Compare(candidate T) int
}

type Node[T Comparable[T]] struct {
	Val                 T
	color               Color
	left, right, parent *Node[T]
}

type Tree[T Comparable[T]] struct {
	root *Node[T]
	Nil  *Node[T]
}

func NewTree[T Comparable[T]]() *Tree[T] {
	nilNode := &Node[T]{color: Black}
	return &Tree[T]{Nil: nilNode, root: nilNode}
}

func (t *Tree[T]) Visualize(w io.Writer) {
	v := treeVisualizer[T]{w}
	v.runOnNode(t.root, "", false)
}

type treeVisualizer[T Comparable[T]] struct {
	w io.Writer
}

func (v *treeVisualizer[T]) runOnNode(node *Node[T], prefix string, hasRightSister bool) {
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
		fmt.Fprintf(v.w, "%s── r%v \n", printPrefix, node.Val)
	} else {
		fmt.Fprintf(v.w, "%s── b%v \n", printPrefix, node.Val)
	}

	prefix += "   " + "│"
	v.runOnNode(node.left, prefix, node.right != nil)
	v.runOnNode(node.right, prefix, false)
}

func (t *Tree[T]) Search(val T) *Node[T] {
	x := t.root
	for x != nil && x.Val.Compare(val) != 0 {
		if val.Compare(x.Val) == -1 {
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

func (t *Tree[T]) Min() *Node[T] {
	x := t.root
	for x.left != t.Nil {
		x = x.left
	}
	return x
}

func (t *Tree[T]) Max() *Node[T] {
	x := t.root
	for x.right != t.Nil {
		x = x.right
	}
	return x
}

func (t *Tree[T]) Successor(n *Node[T]) *Node[T] {
	if n.right != t.Nil {
		rTree := &Tree[T]{root: n.right, Nil: t.Nil}
		return rTree.Min()
	}

	y := n.parent
	for y != t.Nil && n == y.right {
		n = y
		y = y.parent
	}
	return y
}

func (t *Tree[T]) Predecessor(n *Node[T]) *Node[T] {
	if n.left != t.Nil {
		rTree := &Tree[T]{root: n.left, Nil: t.Nil}
		return rTree.Max()
	}

	y := n.parent
	for y != nil && n == y.left {
		n = y
		y = y.parent
	}
	return y
}

func (t *Tree[T]) Add(val T) *Node[T] {
	newNode := &Node[T]{Val: val}
	y := t.Nil
	x := t.root
	for x != t.Nil {
		y = x
		if newNode.Val.Compare(x.Val) == -1 {
			x = x.left
		} else {
			x = x.right
		}
	}

	newNode.parent = y
	if y == t.Nil {
		t.root = newNode
	} else if newNode.Val.Compare(y.Val) == -1 {
		y.left = newNode
	} else {
		y.right = newNode
	}
	newNode.left = t.Nil
	newNode.right = t.Nil
	newNode.color = Red
	t.insertFixup(newNode)

	return newNode
}

func (t *Tree[T]) insertFixup(z *Node[T]) {
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

func (t *Tree[T]) leftRotate(x *Node[T]) {
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

func (t *Tree[T]) rightRotate(y *Node[T]) {
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

func (t *Tree[T]) Order() (result []T) {
	x := t.root
	var stack []*Node[T]
	for len(stack) != 0 || x != t.Nil {
		if x != t.Nil {
			stack = append(stack, x)
			x = x.left
		} else {
			x = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, x.Val)
			x = x.right
		}
	}
	return
}

func (t *Tree[T]) Delete(z *Node[T]) {
	y := z
	yOriginalColor := y.color

	var x *Node[T]
	if z.left == t.Nil {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == t.Nil {
		x = z.left
		t.transplant(z, z.left)
	} else {
		rTree := &Tree[T]{
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

func (t *Tree[T]) transplant(u, v *Node[T]) {
	if u.parent == t.Nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func (t *Tree[T]) deleteFixUp(x *Node[T]) {
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
