package veb

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTreeCreate(t *testing.T) {
	tree := NewTree(16)

	assert.NotNil(t, tree)
}

func TestTreeInsert(t *testing.T) {
	tree := NewTree(16)

	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(14)
	tree.Insert(15)

	assert.NotNil(t, tree)
}

func TestTreeIsMember(t *testing.T) {
	tree := NewTree(16)

	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(14)
	tree.Insert(15)

	assert.True(t, tree.IsMember(2))
	assert.True(t, tree.IsMember(3))
	assert.True(t, tree.IsMember(7))
	assert.True(t, tree.IsMember(14))
	assert.False(t, tree.IsMember(13))
	assert.False(t, tree.IsMember(11))
	assert.False(t, tree.IsMember(1))
}

func TestTreeSuccessor(t *testing.T) {
	tree := NewTree(16)

	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(14)
	tree.Insert(15)

	succ, _ := tree.Successor(2)
	assert.Equal(t, 3, succ)
	succ, _ = tree.Successor(7)
	assert.Equal(t, 14, succ)

	_, exists := tree.Successor(15)
	assert.False(t, exists)
}

func TestTreePredecessor(t *testing.T) {
	tree := NewTree(16)

	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(14)
	tree.Insert(15)

	pred, _ := tree.Predecessor(7)
	assert.Equal(t, 5, pred)
	pred, _ = tree.Predecessor(4)
	assert.Equal(t, 3, pred)

	_, exists := tree.Predecessor(2)
	assert.False(t, exists)
}

func TestTreeDelete(t *testing.T) {
	tree := NewTree(16)

	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(14)
	tree.Insert(15)

	assert.True(t, tree.IsMember(4))
	tree.Delete(4)
	assert.False(t, tree.IsMember(4))

	assert.True(t, tree.IsMember(15))
	tree.Delete(15)
	assert.False(t, tree.IsMember(15))

	assert.True(t, tree.IsMember(7))
	tree.Delete(7)
	assert.False(t, tree.IsMember(7))

	assert.True(t, tree.IsMember(2))
	tree.Delete(2)
	assert.False(t, tree.IsMember(2))
}
