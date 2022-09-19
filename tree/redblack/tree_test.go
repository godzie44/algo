package redblack

import (
	"algorithms"
	"bytes"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type comparableNum int

func (c comparableNum) Compare(candidate comparableNum) int {
	if c < candidate {
		return -1
	} else if c > candidate {
		return 1
	}
	return 0
}

func makeTree() *Tree[comparableNum] {
	t := NewTree[comparableNum]()
	t.Add(comparableNum(1))
	t.Add(comparableNum(3))
	t.Add(comparableNum(5))
	t.Add(comparableNum(6))
	t.Add(comparableNum(8))
	t.Add(comparableNum(9))
	t.Add(comparableNum(11))
	t.Add(comparableNum(12))
	t.Add(comparableNum(14))
	t.Add(comparableNum(15))
	t.Add(comparableNum(18))
	t.Add(comparableNum(19))
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

	tree := NewTree[comparableNum]()
	for _, v := range arr {
		tree.Add(comparableNum(v))
	}
	assertIsRBTree(t, tree)
}

func TestOrder(t *testing.T) {
	arr := algorithms.GenerateRandomSlice(t)

	tree := NewTree[comparableNum]()
	for _, v := range arr {
		tree.Add(comparableNum(v))
	}

	sort.Ints(arr)
	expected := make([]comparableNum, len(arr))
	for i, v := range arr {
		expected[i] = comparableNum(v)
	}

	assert.Equal(t, expected, tree.Order())
}

func TestDelete(t *testing.T) {
	tree := makeTree()
	tree.Delete(tree.Search(12))
	assertIsRBTree(t, tree)
}

func TestDelete2(t *testing.T) {
	arr := algorithms.GenerateRandomSlice(t)
	tree := NewTree[comparableNum]()
	for _, v := range arr {
		tree.Add(comparableNum(v))
	}

	delIdx := 4444
	delNode := tree.Search(comparableNum(arr[delIdx]))
	tree.Delete(delNode)

	arr = append(arr[:delIdx], arr[delIdx+1:]...)
	sort.Ints(arr)

	assertIsRBTree(t, tree)

	expected := make([]comparableNum, len(arr))
	for i, v := range arr {
		expected[i] = comparableNum(v)
	}

	assert.Equal(t, expected, tree.Order())
}

func TestSearch(t *testing.T) {
	tree := makeTree()
	assert.Nil(t, tree.Search(0))
	assert.Equal(t, comparableNum(5), tree.Search(5).Val)
}

func TestMinMax(t *testing.T) {
	tree := makeTree()
	assert.Equal(t, comparableNum(1), tree.Min().Val)
	assert.Equal(t, comparableNum(19), tree.Max().Val)
}

func TestSuccessor(t *testing.T) {
	tree := makeTree()
	assert.Equal(t, comparableNum(3), tree.Successor(tree.Search(1)).Val)
	assert.Equal(t, comparableNum(8), tree.Successor(tree.Search(6)).Val)
	assert.Equal(t, comparableNum(14), tree.Successor(tree.Search(12)).Val)
	assert.Equal(t, comparableNum(18), tree.Successor(tree.Search(15)).Val)
}

func TestPredecessor(t *testing.T) {
	tree := makeTree()
	assert.Equal(t, comparableNum(1), tree.Predecessor(tree.Search(3)).Val)
	assert.Equal(t, comparableNum(6), tree.Predecessor(tree.Search(8)).Val)
	assert.Equal(t, comparableNum(12), tree.Predecessor(tree.Search(14)).Val)
	assert.Equal(t, comparableNum(15), tree.Predecessor(tree.Search(18)).Val)
}

func assertIsRBTree(t *testing.T, tree *Tree[comparableNum]) {
	assert.Equal(t, Black, tree.root.color, "root node color must be black")
	walk(tree.root, func(n *Node[comparableNum]) {
		if n.color == Red {
			assert.Equal(t, Black, n.left.color, "red node must have black child")
			assert.Equal(t, Black, n.right.color, "red node must have black child")
		}
	})

	heights := make(map[int]struct{})
	walk(tree.root, func(n *Node[comparableNum]) {
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

func walk(n *Node[comparableNum], f func(n *Node[comparableNum])) {
	f(n)

	if n.left != nil {
		walk(n.left, f)
	}
	if n.right != nil {
		walk(n.right, f)
	}
}
