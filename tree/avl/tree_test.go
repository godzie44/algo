package avl

import (
	"algorithms"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"
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
	expected := `└── 12 - h(4) 
    ├── 6 - h(3) 
    │   ├── 3 - h(2) 
    │   │   ├── 1 - h(1) 
    │   │   └── 5 - h(1) 
    │   └── 9 - h(2) 
    │       ├── 8 - h(1) 
    │       └── 11 - h(1) 
    └── 15 - h(3) 
        ├── 14 - h(1) 
        └── 18 - h(2) 
            └── 19 - h(1) 
`

	buff := bytes.Buffer{}
	tree := makeTree()
	assertIsAVLTree(t, tree)

	tree.Visualize(&buff)
	assert.Equal(t, expected, buff.String())
}

func TestAVLTreeAdd(t *testing.T) {
	arr := algorithms.GenerateRandomSlice(t)
	tree := NewTree()
	for _, v := range arr {
		tree.Add(v)
	}
	assertIsAVLTree(t, tree)
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

func TestDelete(t *testing.T) {
	tree := makeTree()
	tree.Visualize(os.Stdout)

	tree.Delete(tree.Search(12))
	assertIsAVLTree(t, tree)
}

func TestDelete2(t *testing.T) {
	rand.Seed(time.Now().Unix())
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

	assertIsAVLTree(t, tree)
	assert.Equal(t, arr, tree.Order())
}

func assertIsAVLTree(t *testing.T, tree *Tree) {
	tree.recalculateH()

	tree.walk(tree.root, func(n *Node) {
		if n == tree.Nil {
			return
		}

		assert.True(t, math.Abs(float64(n.left.h-n.right.h)) < 2)
		if math.Abs(float64(n.left.h-n.right.h)) > 1 {
			fmt.Println("wtf", n.val)
		}

	})
}
