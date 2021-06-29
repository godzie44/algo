package btree

import (
	"algorithms"
	"bytes"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func makeTree(tVal int) *Tree {
	t := NewTree(tVal)

	var values []int

	values = []int{16, 1, 21, 24, 20, 13, 7, 15, 14, 25, 23, 12, 11, 3, 10, 17, 2, 26, 18, 22, 9, 5, 6, 4, 8, 19, 28, 29, 49, 33}

	t.Insert(values...)
	return t
}

func TestVisualize(t *testing.T) {
	expected := `   ├    ├ 01 
   ├    ├ 02 
   ├ 03 ┤
   ├    ├ 04 
   ├    ├ 05 
   ├    ├ 06 
   ├ 07 ┤
   ├    ├ 08 
   ├    ├ 09 
   ├    ├ 10 
   ├    ├ 11 
   ├    ├ 12 
13 ┤
   ├    ├ 14 
   ├    ├ 15 
   ├ 16 ┤
   ├    ├ 17 
   ├    ├ 18 
   ├    ├ 19 
   ├ 20 ┤
   ├    ├ 21 
   ├    ├ 22 
   ├    ├ 23 
   ├ 24 ┤
   ├    ├ 25 
   ├    ├ 26 
   ├ 28 ┤
   ├    ├ 29 
   ├    ├ 33 
   ├    ├ 49 
`

	buff := bytes.NewBuffer([]byte{})

	makeTree(3).Visualize(buff)
	assert.Equal(t, expected, buff.String())
}

func TestSearch(t *testing.T) {
	tree := makeTree(3)
	assert.NotNil(t, tree)
	n, ind := tree.Search(5)
	assert.Equal(t, 5, n.Keys[ind])

	n, _ = tree.Search(105)
	assert.Nil(t, n)
}

func TestOrder(t *testing.T) {
	arr := algorithms.GenerateRandomSlice(t)
	tree := NewTree(5)
	tree.Insert(arr...)

	sort.Ints(arr)
	assert.Equal(t, arr, tree.Order())
}

func TestDelete(t *testing.T) {
	tree := makeTree(3)
	assert.NotNil(t, tree)

	tree.Delete(1)
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 28, 29, 33, 49}, tree.Order())

	tree.Delete(13)
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 28, 29, 33, 49}, tree.Order())

	tree.Delete(24)
	assert.Equal(t, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 25, 26, 28, 29, 33, 49}, tree.Order())

	tree.Delete(3, 20, 15, 12, 33, 2, 4, 7, 14, 5, 6, 28, 22, 8, 49, 29, 9, 10, 11, 25, 17, 21, 26)
	assert.Equal(t, []int{16, 18, 19, 23}, tree.Order())
}
