package bstree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func makeTree() *Tree {
	t := NewTree()
	t.Add(9)
	t.Add(5)
	t.Add(3)
	t.Add(6)
	t.Add(1)
	t.Add(8)
	t.Add(15)
	t.Add(12)
	t.Add(18)
	t.Add(11)
	t.Add(14)
	t.Add(19)
	return t
}

func TestPrintTree(t *testing.T) {
	makeTree().Print()
}

func TestOrder(t *testing.T) {
	tree := makeTree()
	assert.Equal(t, []int{1, 3, 5, 6, 8, 9, 11, 12, 14, 15, 18, 19}, tree.Order())
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
	tree.Delete(tree.Search(9))
	assert.Equal(t, []int{1, 3, 5, 6, 8, 11, 12, 14, 15, 18, 19}, tree.Order())

	tree.Delete(tree.Search(1))
	assert.Equal(t, []int{3, 5, 6, 8, 11, 12, 14, 15, 18, 19}, tree.Order())

	tree.Delete(tree.Search(11))
	assert.Equal(t, []int{3, 5, 6, 8, 12, 14, 15, 18, 19}, tree.Order())

	tree.Delete(tree.Search(18))
	assert.Equal(t, []int{3, 5, 6, 8, 12, 14, 15, 19}, tree.Order())
}
