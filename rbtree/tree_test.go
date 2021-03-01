package rbtree

import (
	"algorithms"
	"bytes"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func makeTree() *Tree {
	t := NewTree()
	t.Add(1)
	t.Add(3)
	t.Add(5)
	t.Add(6)
	t.Add(8)
	t.Add(9)
	t.Add(11)
	t.Add(12)
	t.Add(14)
	t.Add(15)
	t.Add(18)
	t.Add(19)
	return t
}

func TestPrintTree(t *testing.T) {
	expect := `└── b6 
    ├── b3 
    │   ├── b1 
    │   │   ├── b0 
    │   │   └── b0 
    │   └── b5 
    │       ├── b0 
    │       └── b0 
    └── b12 
        ├── r9 
        │   ├── b8 
        │   │   ├── b0 
        │   │   └── b0 
        │   └── b11 
        │       ├── b0 
        │       └── b0 
        └── r15 
            ├── b14 
            │   ├── b0 
            │   └── b0 
            └── b18 
                ├── b0 
                └── r19 
                    ├── b0 
                    └── b0 
`

	buff := bytes.Buffer{}
	makeTree().Visualize(&buff)

	assert.Equal(t, expect, buff.String())
}

func TestAdd(t *testing.T) {
	arr := algorithms.GenerateRandomSlice(t)
	tree := NewTree()
	for _, v := range arr {
		tree.Add(v)
	}
	assertIsRBTree(t, tree)
}

func TestOrder(t *testing.T) {
	arr := algorithms.GenerateRandomSlice(t)
	tree := NewTree()
	for _, v := range arr {
		tree.Add(v)
	}

	sort.Ints(arr)
	assert.Equal(t, arr, tree.Order())
}

func TestDelete(t *testing.T) {
	tree := makeTree()
	tree.Delete(tree.Search(12))
	assertIsRBTree(t, tree)
}

func TestDelete2(t *testing.T) {
	arr := algorithms.GenerateRandomSlice(t)
	tree := NewTree()
	for _, v := range arr {
		tree.Add(v)
	}

	delIdx := 4444
	delNode := tree.Search(arr[delIdx])
	tree.Delete(delNode)

	arr = append(arr[:delIdx], arr[delIdx+1:]...)
	sort.Ints(arr)

	assertIsRBTree(t, tree)
	assert.Equal(t, arr, tree.Order())
}

func TestSearch(t *testing.T) {
	tree := makeTree()
	assert.Nil(t, tree.Search(0))
	assert.Equal(t, 5, tree.Search(5).val)
}

func TestMinMax(t *testing.T) {
	tree := makeTree()
	assert.Equal(t, 1, tree.Min().val)
	assert.Equal(t, 19, tree.Max().val)
}

func TestSuccessor(t *testing.T) {
	tree := makeTree()
	assert.Equal(t, 3, tree.Successor(tree.Search(1)).val)
	assert.Equal(t, 8, tree.Successor(tree.Search(6)).val)
	assert.Equal(t, 14, tree.Successor(tree.Search(12)).val)
	assert.Equal(t, 18, tree.Successor(tree.Search(15)).val)
}

func TestPredecessor(t *testing.T) {
	tree := makeTree()
	assert.Equal(t, 1, tree.Predecessor(tree.Search(3)).val)
	assert.Equal(t, 6, tree.Predecessor(tree.Search(8)).val)
	assert.Equal(t, 12, tree.Predecessor(tree.Search(14)).val)
	assert.Equal(t, 15, tree.Predecessor(tree.Search(18)).val)
}

func assertIsRBTree(t *testing.T, tree *Tree) {
	assert.Equal(t, Black, tree.root.color, "root node color must be black")
	walk(tree.root, func(n *Node) {
		if n.color == Red {
			assert.Equal(t, Black, n.left.color, "red node must have black child")
			assert.Equal(t, Black, n.right.color, "red node must have black child")
		}
	})

	heights := make(map[int]struct{})
	walk(tree.root, func(n *Node) {
		if n.left == tree.Nil && n.right == tree.Nil {
			var h int
			for n.parent != tree.Nil {
				if n.color == Black {
					h++
				}
				n = n.parent
			}
			heights[h] = struct{}{}
		}
	})

	assert.Len(t, heights, 1, "all leaves heights must be equal")
}

func walk(n *Node, f func(n *Node)) {
	f(n)

	if n.left != nil {
		walk(n.left, f)
	}
	if n.right != nil {
		walk(n.right, f)
	}
}
